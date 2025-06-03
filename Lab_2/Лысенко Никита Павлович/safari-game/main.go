package main

import (
    "bufio"
    "encoding/json"
    "html/template"
    "log"
    "math/rand"
    "net/http"
    "os"
    "strings"
    "sync"
    "time"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Game struct {
    PlayerX  float64 `json:"player_x"`
    PlayerY  float64 `json:"player_y"`
    Meat     []Item  `json:"meat"`
    Animals  []Item  `json:"animals"`
    Score    int     `json:"score"`
    GameOver bool    `json:"game_over"`
    mu       sync.Mutex
}

type Item struct {
    X      float64 `json:"x"`
    Y      float64 `json:"y"`
    Width  float64 `json:"width"`
    Height float64 `json:"height"`
}

var (
    game        Game
    currentUser  string
)

func serveMenu(w http.ResponseWriter, r *http.Request) {
    if currentUser  == "" {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }
    // Перенаправление на страницу с игрой
    http.Redirect(w, r, "/static/index.html", http.StatusSeeOther)
}

func serveLogin(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")

        if checkCredentials(username, password) {
            currentUser  = username
            http.Redirect(w, r, "/menu", http.StatusSeeOther)
            return
        }
        http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
        return
    }
    http.ServeFile(w, r, "./static/login.html")
}

func serveRegister(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")

        if username == "" || password == "" {
            http.Error(w, "Введите логин и пароль", http.StatusBadRequest)
            return
        }

        if userExists(username) {
            http.Error(w, "Пользователь уже существует", http.StatusConflict)
            return
        }

        err := saveUser (username, password)
        if err != nil {
            http.Error(w, "Ошибка при сохранении пользователя", http.StatusInternalServerError)
            return
        }

        currentUser  = username
        http.Redirect(w, r, "/menu", http.StatusSeeOther)
    } else {
        http.ServeFile(w, r, "./static/register.html")
    }
}

func serveProfile(w http.ResponseWriter, r *http.Request) {
    if currentUser  == "" {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }
    data := struct {
        Username string
    }{Username: currentUser }
    t, err := template.ParseFiles("./static/profile.html")
    if err != nil {
        http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
        return
    }
    t.Execute(w, data)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    currentUser  = ""
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        game.mu.Lock()
        defer game.mu.Unlock()
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(game)
    case http.MethodPost:
        var input struct {
            PlayerX float64 `json:"player_x"`
            PlayerY float64 `json:"player_y"`
        }
        err := json.NewDecoder(r.Body).Decode(&input)
        if err != nil {
            http.Error(w, "Неверные данные", http.StatusBadRequest)
            return
        }

        game.mu.Lock()
        game.PlayerX = input.PlayerX
        game.PlayerY = input.PlayerY
        game.checkCollisions()
        game.mu.Unlock()
    default:
        http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
    }
}

func (g *Game) checkCollisions() {
    // Проверяем столкновение с мясом
    for i := len(g.Meat) - 1; i >= 0; i-- {
        meat := g.Meat[i]
        if meat.Y+meat.Height >= g.PlayerY &&
            meat.X < g.PlayerX+40 &&
            meat.X+meat.Width > g.PlayerX &&
            meat.Y < g.PlayerY+40 {
            g.Score += 10
            g.Meat = append(g.Meat[:i], g.Meat[i+1:]...)
        }
    }

    // Проверяем столкновение с животными
    for _, animal := range g.Animals {
        if animal.Y+animal.Height >= g.PlayerY &&
            animal.X < g.PlayerX+40 &&
            animal.X+animal.Width > g.PlayerX &&
            animal.Y < g.PlayerY+40 {
            g.GameOver = true
            break
        }
    }
}

func gameLoop() {
    rand.Seed(time.Now().UnixNano())
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()

    for {
        <-ticker.C
        game.mu.Lock()
        if !game.GameOver {
            spawnMeat()
            spawnAnimal()
        }
        game.mu.Unlock()
    }
}

func spawnMeat() {
    if len(game.Meat) < 5 {
        newMeat := Item{
            X:      float64(rand.Intn(580)), // 600 - 20 ширина мяса
            Y:      float64(rand.Intn(380)), // 400 - 20 высота мяса
            Width:  20,
            Height: 20,
        }

        if !isOverlapping(newMeat, game.Meat) && !isOverlapping(newMeat, game.Animals) {
            game.Meat = append(game.Meat, newMeat)
        }
    }
}

func spawnAnimal() {
    if len(game.Animals) < 15 {
        newAnimal := Item{
            X:      float64(rand.Intn(560)), // 600 - 40 ширина животного
            Y:      float64(rand.Intn(360)), // 400 - 40 высота животного
            Width:  40,
            Height: 40,
        }

        if !isOverlapping(newAnimal, game.Animals) && !isOverlapping(newAnimal, game.Meat) {
            game.Animals = append(game.Animals, newAnimal)
        }
    }
}

func isOverlapping(newItem Item, existingItems []Item) bool {
    for _, item := range existingItems {
        if newItem.X < item.X+item.Width &&
            newItem.X+newItem.Width > item.X &&
            newItem.Y < item.Y+item.Height &&
            newItem.Y+newItem.Height > item.Y {
            return true
        }
    }
    return false
}

func userExists(username string) bool {
    file, err := os.Open("users.txt")
    if err != nil {
        // Если файла нет, значит пользователя нет
        if os.IsNotExist(err) {
            return false
        }
        log.Println("Ошибка при открытии users.txt:", err)
        return false
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, ":", 2)
        if len(parts) == 2 && parts[0] == username {
            return true
        }
    }
    return false
}

func saveUser (username, password string) error {
    file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println("Ошибка при открытии файла для записи:", err)
        return err
    }
    defer file.Close()

    _, err = file.WriteString(username + ":" + password + "\n")
    if err != nil {
        log.Println("Ошибка при записи в файл:", err)
        return err
    }
    return nil
}

func checkCredentials(username, password string) bool {
    file, err := os.Open("users.txt")
    if err != nil {
        log.Println("Ошибка при открытии users.txt:", err)
        return false
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, ":", 2)
        if len(parts) == 2 && parts[0] == username && parts[1] == password {
            return true
        }
    }
    return false
}

func main() {
    // Инициализация игры
    game = Game{
        PlayerX:  float64(280), // По центру по X (canvas ширина 600 - playerSize 40)/2
        PlayerY:  float64(350), // Почти внизу (canvas высота 400 - playerSize 40 - 10)
        Meat:     []Item{},
        Animals:  []Item{},
        Score:    0,
        GameOver: false,
    }

    http.HandleFunc("/", serveMenu) // Начальная страница - меню
    http.HandleFunc("/game", gameHandler)
    http.HandleFunc("/menu", serveMenu)
    http.HandleFunc("/login", serveLogin)
    http.HandleFunc("/register", serveRegister)
    http.HandleFunc("/profile", serveProfile)
    http.HandleFunc("/logout", logoutHandler)

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    go gameLoop()

    log.Println("Server started at :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
