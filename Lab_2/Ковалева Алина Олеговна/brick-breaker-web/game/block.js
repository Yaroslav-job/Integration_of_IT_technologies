import { COLORS, BRICKS_CONFIG } from './constants.js';

export class Block {
  // Конструктор блока (кирпича)
  constructor(x, y, width, height, destroyed = false, textureIndex = 1) {
    this.x = x;                 // Позиция по X
    this.y = y;                 // Позиция по Y
    this.width = width;         // Ширина блока
    this.height = height;       // Высота блока
    this.destroyed = destroyed; // Состояние блока (разрушен или нет)
    
    // Индекс текстуры, ограниченный числом доступных текстур
    this.textureIndex = textureIndex % BRICKS_CONFIG.TEXTURES_COUNT + 1;

    this.initImage();           // Загружаем изображение кирпича
  }

  // Загружаем изображение кирпича
  initImage() {
    this.image = new Image();
    this.image.src = `assets/brick${this.textureIndex}.png`; // путь к текстуре
    this.image.onload = () => this.imageLoaded = true;       // отметка о загрузке
    this.imageLoaded = false;
  }

  // Сброс состояния кирпича (например, при перезапуске уровня)
  reset() {
    this.destroyed = false;
  }

  // Отрисовка кирпича на холсте
  draw(ctx) {
    if (this.destroyed) return; // Не рисуем, если кирпич уже разрушен

    if (this.imageLoaded) {
      // Рисуем изображение кирпича
      ctx.drawImage(this.image, this.x, this.y, this.width, this.height);
    } else {
      // Временный вариант отрисовки, если изображение не загружено
      ctx.fillStyle = COLORS.BRICK;
      ctx.fillRect(this.x, this.y, this.width, this.height);
      ctx.strokeStyle = COLORS.BRICK_BORDER;
      ctx.strokeRect(this.x, this.y, this.width, this.height);
    }
  }

  // Проверка столкновения мяча с кирпичом
  checkCollision(ball) {
    if (this.destroyed) return false; // Пропускаем, если кирпич уже разрушен

    // Проверка пересечения мяча и блока по X и Y
    const withinX = ball.x + ball.radius > this.x && ball.x - ball.radius < this.x + this.width;
    const withinY = ball.y + ball.radius > this.y && ball.y - ball.radius < this.y + this.height;

    if (withinX && withinY) {
      // Если столкновение произошло:
      this.destroyed = true;      // Удаляем кирпич
      ball.speedY *= -1;          // Меняем направление мяча по вертикали
      return true;
    }

    return false;
  }
}
