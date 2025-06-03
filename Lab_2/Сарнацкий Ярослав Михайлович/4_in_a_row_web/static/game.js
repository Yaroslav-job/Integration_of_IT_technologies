let boardDiv = document.getElementById('board');
let statusP = document.getElementById('status');
let gameChoiceDiv = document.getElementById('game-choice');

function startGame(type) {
  fetch(`/start?type=${type}`)
    .then(res => res.json())
    .then(data => {
      gameChoiceDiv.style.display = 'none'; // Скрыть выбор игры после начала
      document.getElementById("board").classList.add("visible");
      renderBoard(data.board);
      updateGame(data);
    });
}

function renderBoard(board) {
  boardDiv.innerHTML = ''; // Очищаем доску перед отрисовкой
  for (let i = 0; i < 6; i++) {
    for (let j = 0; j < 7; j++) {
      const cell = document.createElement('div');
      cell.className = 'cell';
      cell.textContent = board[i][j] || ''; // Показать пустую ячейку, если нет фишки
      cell.dataset.column = j;
      cell.addEventListener('click', () => move(j));
      boardDiv.appendChild(cell);
    }
  }
  boardDiv.style.display = 'grid'; // Убедитесь, что стиль "grid" применён
}

function updateGame(data) {
  renderBoard(data.board);
  if (data.game_over) {
    if (data.winner) {
      statusP.textContent = `Победил: ${data.winner}`;
    } else {
      statusP.textContent = 'Ничья!';
    }
    gameChoiceDiv.style.display = 'block'; // Показать выбор типа игры после завершения
  } else {
    statusP.textContent = `Ходит: ${data.turn}`;
  }
}

function move(column) {
  fetch(`/move?column=${column}`)
    .then(res => res.json())
    .then(updateGame);
}
