// Основные настройки игры
export const GAME_CONFIG = {
    CANVAS_WIDTH: 700,
    CANVAS_HEIGHT: 500,
    DIFFICULTY_INTERVAL: 15000, // 15 секунд
    INITIAL_PADDLE_WIDTH: 130,
    INITIAL_BALL_SPEED: 3.5,
    BALL_RADIUS: 10,
    PADDLE_HEIGHT: 15,
    PADDLE_SPEED: 8,
    MAX_BALL_SPEED: 7,
    MIN_PADDLE_WIDTH: 40,
    DIFFICULTY_INCREASE_RATE: 1.1, // 10% увеличение
    BALL_BOUNCE_FACTOR: 0.15
  };
  
  // Настройки блоков
  export const BRICKS_CONFIG = {
    ROWS: 10,
    COLS: 12,
    HEIGHT: 20,
    PADDING: 1,
    TEXTURES_COUNT: 9
  };
  
  // Цвета для фолбэка
  export const COLORS = {
    BALL: 'white',
    PADDLE: 'white',
    BRICK: 'tomato',
    BRICK_BORDER: 'black'
  };
  
  // Тексты игры
  export const TEXTS = {
    WIN: 'You Win!',
    LOSE: 'Game Over',
    DIFFICULTY_INCREASE: 'Стало сложнее! ;)'
  };