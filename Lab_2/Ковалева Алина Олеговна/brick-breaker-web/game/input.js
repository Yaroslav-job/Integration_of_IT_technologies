// Функция для настройки управления с клавиатуры
// Принимает объект paddle (платформа), которому будет передаваться направление движения
export function setupInput(paddle) {
    // Обработчик события нажатия клавиши
    const handleKeyDown = (e) => {
      if (e.key === "ArrowLeft") paddle.setDirection("left");   // Влево
      if (e.key === "ArrowRight") paddle.setDirection("right"); // Вправо
    };
  
    // Обработчик события отпускания клавиши
    const handleKeyUp = (e) => {
      if (["ArrowLeft", "ArrowRight"].includes(e.key)) {
        paddle.setDirection("stop"); // Остановка при отпускании стрелок
      }
    };
  
    // Удаление предыдущих обработчиков, если они уже были назначены
    // Это помогает избежать повторного добавления и дублирования событий
    document.removeEventListener("keydown", handleKeyDown);
    document.removeEventListener("keyup", handleKeyUp);
  
    // Назначение новых обработчиков событий
    document.addEventListener("keydown", handleKeyDown);
    document.addEventListener("keyup", handleKeyUp);
  }
  