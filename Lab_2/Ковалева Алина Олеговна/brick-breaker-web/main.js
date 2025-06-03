// Импортируем игровые сущности и конфигурации
import { Ball } from './game/ball.js';
import { Paddle } from './game/paddle.js';
import { Block } from './game/block.js';
import { setupInput } from './game/input.js';
import { GAME_CONFIG, BRICKS_CONFIG, TEXTS } from './game/constants.js';

class Game {
    constructor() {
        // Получаем доступ к canvas и его контексту
        this.canvas = document.getElementById('gameCanvas');
        this.ctx = this.canvas.getContext('2d');
      
        // Получаем элементы стартового меню
        this.startMenu = document.getElementById('startMenu');
        this.startButton = document.getElementById('startButton');
      
        // Получаем элементы меню окончания игры
        this.gameOverMenu = document.getElementById('gameOverMenu');
        this.gameOverText = document.getElementById('gameOverText');
        this.restartButton = document.getElementById('restartButton');
      
        // Настраиваем размеры canvas и события
        this.setupCanvas();
        this.setupEventListeners();
    }
      
  // Устанавливаем размеры игрового поля
  setupCanvas() {
    this.canvas.width = GAME_CONFIG.CANVAS_WIDTH;
    this.canvas.height = GAME_CONFIG.CANVAS_HEIGHT;
  }
  
  // Инициализируем игровые объекты: шарик, платформу и кирпичи
  initGameObjects() {
    const centerX = this.canvas.width / 2;
    const centerY = this.canvas.height / 2;
    
    // Создаём шарик по центру
    this.ball = new Ball(centerX, centerY);

    // Создаём платформу внизу экрана
    this.paddle = new Paddle(
      centerX - GAME_CONFIG.INITIAL_PADDLE_WIDTH / 2,
      this.canvas.height - GAME_CONFIG.PADDLE_HEIGHT
    );
    
    // Создаём массив кирпичей
    this.createBricks();

    // Устанавливаем управление платформой
    setupInput(this.paddle);
  }
  
  // Генерируем сетку кирпичей
  createBricks() {
    this.blocks = [];
    const brickWidth = (this.canvas.width - BRICKS_CONFIG.PADDING * (BRICKS_CONFIG.COLS - 1)) / BRICKS_CONFIG.COLS;
    
    for (let row = 0; row < BRICKS_CONFIG.ROWS; row++) {
      for (let col = 0; col < BRICKS_CONFIG.COLS; col++) {
        const x = col * (brickWidth + BRICKS_CONFIG.PADDING);
        const y = row * (BRICKS_CONFIG.HEIGHT + BRICKS_CONFIG.PADDING);
        const textureIndex = (row + col) % BRICKS_CONFIG.TEXTURES_COUNT + 1;
        
        this.blocks.push(new Block(x, y, brickWidth, BRICKS_CONFIG.HEIGHT, false, textureIndex));
      }
    }
  }
  
  // Настраиваем обработчики нажатий на кнопки меню
  setupEventListeners() {
    this.startButton.addEventListener('click', () => {
      this.startMenu.classList.add('hidden'); // Скрываем стартовое меню
      this.initGameObjects(); // Инициализируем игру
      this.startGame(); // Запускаем цикл игры
    });
  
    this.restartButton.addEventListener('click', () => this.resetGame()); // Кнопка рестарта
  }
  
  // Запуск основной логики игры
  startGame() {
    this.gameStartTime = Date.now();
    this.lastDifficultyUpdate = 0;
    this.gameActive = true;
    this.gameLoop(); // Запускаем игровой цикл
  }
  
  // Увеличение сложности игры (ускорение мяча и уменьшение платформы)
  increaseDifficulty() {
    // Уменьшаем ширину платформы
    const newWidth = Math.max(
      GAME_CONFIG.MIN_PADDLE_WIDTH, 
      this.paddle.width * (1 / GAME_CONFIG.DIFFICULTY_INCREASE_RATE)
    );
    
    const widthDiff = this.paddle.width - newWidth;
    this.paddle.width = newWidth;
    this.paddle.x += widthDiff / 2;

    // Увеличиваем скорость мяча
    this.ball.speedX = Math.min(
      GAME_CONFIG.MAX_BALL_SPEED, 
      this.ball.speedX * GAME_CONFIG.DIFFICULTY_INCREASE_RATE
    );
    
    this.ball.speedY = Math.min(
      GAME_CONFIG.MAX_BALL_SPEED, 
      this.ball.speedY * GAME_CONFIG.DIFFICULTY_INCREASE_RATE
    );

    // Показываем уведомление об увеличении сложности
    this.showDifficultyNotification();
  }
  
  // Показать уведомление об увеличении сложности
  showDifficultyNotification() {
    const notification = document.createElement('div');
    notification.className = 'difficulty-notification';
    notification.textContent = TEXTS.DIFFICULTY_INCREASE;
    document.body.appendChild(notification);
    
    setTimeout(() => notification.remove(), 2000); // Удаляем через 2 секунды
  }
  
  // Сброс игры при рестарте
  resetGame() {
    cancelAnimationFrame(this.animationId); // Останавливаем предыдущий цикл
    this.initGameObjects(); // Инициализируем заново
    this.gameOverMenu.classList.add('hidden'); // Скрываем меню окончания
    this.startGame(); // Запускаем игру снова
  }
  
  // Проверка условия победы — все кирпичи уничтожены
  checkWin() {
    return this.blocks.every(block => block.destroyed);
  }
  
  // Показать финальный экран (победа или поражение)
  showGameOver(isWin) {
    this.gameOverText.textContent = isWin ? TEXTS.WIN : TEXTS.LOSE;
    this.gameOverMenu.classList.remove('hidden');
    this.gameActive = false;
    cancelAnimationFrame(this.animationId); // Остановить анимацию
  }
  
  // Обновление состояния игры на каждом кадре
  update() {
    if (!this.gameActive) return;
    
    const currentTime = Date.now();

    // Проверка на необходимость увеличения сложности
    if (currentTime - this.gameStartTime - this.lastDifficultyUpdate >= GAME_CONFIG.DIFFICULTY_INTERVAL) {
      this.increaseDifficulty();
      this.lastDifficultyUpdate = currentTime - this.gameStartTime;
    }

    // Обновляем позицию мяча и платформы
    const ballLost = this.ball.update(this.canvas.width, this.canvas.height, this.paddle);
    this.paddle.update(this.canvas.width);

    // Проверяем столкновение мяча с каждым кирпичом
    this.blocks.forEach(block => {
      if (!block.destroyed) {
        block.checkCollision(this.ball);
      }
    });

    // Проверка условий окончания игры
    if (ballLost) {
      this.showGameOver(false); // Проигрыш
    } else if (this.checkWin()) {
      this.showGameOver(true); // Победа
    }
  }
  
  // Отрисовка объектов на canvas
  render() {
    this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height); // Очистка экрана
    
    // Рисуем все кирпичи, мяч и платформу
    this.blocks.forEach(block => block.draw(this.ctx));
    this.ball.draw(this.ctx);
    this.paddle.draw(this.ctx);
  }
  
  // Главный игровой цикл
  gameLoop() {
    this.update(); // Обновление состояния
    this.render(); // Отрисовка
    
    if (this.gameActive) {
      this.animationId = requestAnimationFrame(() => this.gameLoop()); // Следующий кадр
    }
  }
}

// Запуск игры при полной загрузке страницы
window.addEventListener('load', () => {
  const game = new Game();
});
