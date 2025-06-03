import { GAME_CONFIG, COLORS } from './constants.js';

// Класс Paddle представляет платформу, которой управляет игрок
export class Paddle {
  constructor(
    x, y,
    width = GAME_CONFIG.INITIAL_PADDLE_WIDTH,
    height = GAME_CONFIG.PADDLE_HEIGHT,
    speed = GAME_CONFIG.PADDLE_SPEED
  ) {
    this.x = x;                   // Позиция по X
    this.y = y;                   // Позиция по Y
    this.width = width;          // Ширина платформы
    this.height = height;        // Высота платформы
    this.speed = speed;          // Скорость перемещения
    this.direction = "stop";     // Направление движения (left, right, stop)
    this.initialState = { x, y, width }; // Сохранение исходного состояния
    this.initImage();            // Загрузка изображения платформы
  }

  // Инициализация изображения платформы
  initImage() {
    this.image = new Image();
    this.image.src = 'assets/paddle.png';
    this.image.onload = () => this.imageLoaded = true;
    this.imageLoaded = false;
  }

  // Сброс состояния платформы (используется при рестарте игры)
  reset(width = this.initialState.width) {
    this.x = this.initialState.x;
    this.width = width;         // Можно менять ширину при сбросе (например, при бонусах)
    this.direction = "stop";    // Останавливаем движение
  }

  // Установка направления движения по клавишам
  setDirection(dir) {
    this.direction = dir;
  }

  // Обновление позиции платформы в зависимости от направления и ограничения по краям
  update(canvasWidth) {
    if (this.direction === "left") {
      this.x -= this.speed;
    } else if (this.direction === "right") {
      this.x += this.speed;
    }

    // Не выходим за пределы экрана
    this.x = Math.max(0, Math.min(this.x, canvasWidth - this.width));
  }

  // Отрисовка платформы: либо изображение, либо цветной прямоугольник
  draw(ctx) {
    if (this.imageLoaded) {
      ctx.drawImage(this.image, this.x, this.y, this.width, this.height);
    } else {
      ctx.fillStyle = COLORS.PADDLE;
      ctx.fillRect(this.x, this.y, this.width, this.height);
    }
  }
}
