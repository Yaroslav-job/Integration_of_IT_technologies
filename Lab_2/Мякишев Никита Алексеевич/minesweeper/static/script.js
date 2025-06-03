let startTime, timerInterval;

function initGame() {
  document.getElementById('restart').onclick = newGame;
  newGame(); // сразу стартует новая игра
}

async function loadBoard() {
  const res = await fetch('/api/board');
  const board = await res.json();
  renderBoard(board);
  updateFlagCount(board);
}

function renderBoard(board) {
  const container = document.getElementById('game');
  container.innerHTML = '';
  const table = document.createElement('table');

  board.forEach((row, y) => {
    const tr = document.createElement('tr');
    row.forEach((cell, x) => {
      const td = document.createElement('td');
      if (cell.IsRevealed) {
        td.className = 'revealed';
        td.textContent = cell.IsMine ? '💣' : (cell.NeighborMines || '');
      } else if (cell.IsFlagged) {
        td.textContent = '🚩';
      }

      td.oncontextmenu = (e) => {
        e.preventDefault();
        toggleFlag(x, y);
      };

      if (!cell.IsRevealed && !cell.IsFlagged) {
        td.addEventListener('click', () => reveal(x, y));
      }

      tr.appendChild(td);
    });
    table.appendChild(tr);
  });

  container.appendChild(table);
}

async function reveal(x, y) {
  await fetch('/api/reveal', {
    method: 'POST',
    body: JSON.stringify({ X: x, Y: y })
  });
  loadBoard();
}

async function toggleFlag(x, y) {
  await fetch('/api/flag', {
    method: 'POST',
    body: JSON.stringify({ X: x, Y: y })
  });
  loadBoard();
}

function updateFlagCount(board) {
  let flags = 0;
  board.forEach(row =>
    row.forEach(cell => {
      if (cell.IsFlagged) flags++;
    })
  );
  document.getElementById('flags').textContent = `🚩 Флаги: ${flags}`;
}

function startTimer() {
  startTime = Date.now();
  clearInterval(timerInterval); // сбрасываем предыдущий интервал
  timerInterval = setInterval(() => {
    const seconds = Math.floor((Date.now() - startTime) / 1000);
    document.getElementById('timer').textContent = `⏱ Время: ${seconds} сек`;
  }, 1000);
}

async function newGame() {
  await fetch('/api/new');
  startTimer();
  loadBoard();
}

// Ждём полной загрузки DOM
window.addEventListener('DOMContentLoaded', initGame);
