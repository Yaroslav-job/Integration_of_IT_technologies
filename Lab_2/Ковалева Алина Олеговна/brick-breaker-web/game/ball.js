import { GAME_CONFIG, COLORS } from './constants.js';

export class Ball {
  // Конструктор: создаёт мяч с заданной позицией, радиусом и скоростью
  constructor(x, y, radius = GAME_CONFIG.BALL_RADIUS, speed = GAME_CONFIG.INITIAL_BALL_SPEED) {
    this.x = x;                       // Позиция по X
    this.y = y;                       // Позиция по Y
    this.radius = radius;             // Радиус мяча
    this.speedX = speed;              // Горизонтальная скорость
    this.speedY = -speed;             // Вертикальная скорость (вверх)
    this.initialState = {            // Сохраняем начальное состояние для сброса
      x,
      y,
      speedX: speed,
      speedY: -speed
    };
    this.initImage();                // Инициализируем изображение мяча
  }

  // Загружает изображение мяча и сохраняет его размеры после загрузки
  initImage() {
    this.image = new Image();
    this.image.src = 'assets/ball.png';
    this.image.onload = () => {
      this.imageLoaded = true;
      this.scaledWidth = this.radius * 2;
      this.scaledHeight = this.radius * 2;
    };
    this.imageLoaded = false;
  }

  // Сброс мяча в начальное состояние
  reset() {
    Object.assign(this, this.initialState);
  }

  // Проверяет столкновение мяча с платформой
  checkPaddleCollision(paddle) {
    // Проверяем, находится ли мяч в пределах платформы по X и Y
    const withinX = this.x > paddle.x && this.x < paddle.x + paddle.width;
    const withinY = this.y + this.radius >= paddle.y;

    if (withinX && withinY) {
      // Мяч отскакивает вверх
      this.speedY *= -1;

      // Вычисляем, где именно по ширине платформы произошёл удар
      const paddleCenter = paddle.x + paddle.width / 2;
      const hitPosition = (this.x - paddleCenter) / (paddle.width / 2); // от -1 до 1
      const deltaX = hitPosition * GAME_CONFIG.BALL_BOUNCE_FACTOR * GAME_CONFIG.MAX_BALL_SPEED;

      // Корректируем горизонтальную скорость в зависимости от места удара
      this.speedX += deltaX;

      // Ограничиваем скорость мяча по максимуму
      this.speedX = Math.max(-GAME_CONFIG.MAX_BALL_SPEED, Math.min(GAME_CONFIG.MAX_BALL_SPEED, this.speedX));

      // Смещаем мяч наверх, чтобы он не прилип к платформе
      this.y = paddle.y - this.radius;

      return true;
    }

    return false;
  }

  // Обновляет положение мяча на каждом кадре игры
  update(canvasWidth, canvasHeight, paddle) {
    // Перемещаем мяч по координатам
    this.x += this.speedX;
    this.y += this.speedY;

    // Ограничиваем скорость по X и Y
    this.speedX = Math.min(GAME_CONFIG.MAX_BALL_SPEED, Math.max(-GAME_CONFIG.MAX_BALL_SPEED, this.speedX));
    this.speedY = Math.min(GAME_CONFIG.MAX_BALL_SPEED, Math.max(-GAME_CONFIG.MAX_BALL_SPEED, this.speedY));

    // Отскок от боковых стен
    if (this.x - this.radius < 0 || this.x + this.radius > canvasWidth) {
      this.speedX *= -1;
    }

    // Отскок от верхней стены
    if (this.y - this.radius < 0) {
      this.speedY *= -1;
    }

    // Проверка столкновения с платформой
    this.checkPaddleCollision(paddle);

    // Возвращает true, если мяч упал за нижнюю границу (проигрыш)
    return this.y + this.radius > canvasHeight;
  }

  // Отрисовывает мяч
  draw(ctx) {
    if (this.imageLoaded) {
      // Рисуем изображение мяча, если оно загружено
      ctx.drawImage(
        this.image,
        this.x - this.radius,
        this.y - this.radius,
        this.scaledWidth,
        this.scaledHeight
      );
    } else {
      // Временная отрисовка круга, если изображение ещё не загружено
      ctx.beginPath();
      ctx.fillStyle = COLORS.BALL;
      ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2);
      ctx.fill();
      ctx.closePath();
    }
  }
}
