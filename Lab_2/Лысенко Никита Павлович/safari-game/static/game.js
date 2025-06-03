const canvas = document.getElementById('gameCanvas');
const ctx = canvas.getContext('2d');

const playerImage = new Image();
playerImage.src = '/static/player.png'; // Путь к изображению игрока
const meatImage = new Image();
meatImage.src = '/static/meat.png'; // Путь к изображению мяса
const animalImage = new Image();
animalImage.src = '/static/animal.png'; // Путь к изображению животного

const playerSize = 40;
let playerX = canvas.width / 2 - playerSize / 2;
let playerY = canvas.height - playerSize - 10;

let score = 0;
let gameOver = false;

function updateGame() {
    fetch('/game')
        .then(response => response.json())
        .then(data => {
            playerX = data.player_x;
            playerY = data.player_y;
            score = data.score;
            gameOver = data.game_over;

            ctx.clearRect(0, 0, canvas.width, canvas.height);
            drawPlayer();
            drawItems(data.meat, meatImage);
            drawItems(data.animals, animalImage);

            document.getElementById('score').textContent = score;

            if (gameOver) {
                ctx.fillStyle = 'rgba(0, 0, 0, 0.7)';
                ctx.fillRect(0, 0, canvas.width, canvas.height);
                ctx.fillStyle = 'white';
                ctx.font = '48px Arial';
                ctx.fillText('Game Over!', canvas.width / 2 - 130, canvas.height / 2);
                ctx.font = '24px Arial';
                ctx.fillText('Общий счет: ' + score, canvas.width / 2 - 80, canvas.height / 2 + 40);
            }
        })
        .catch(err => {
            console.error("Ошибка при обновлении игры:", err);
        });
}

function drawPlayer() {
    ctx.drawImage(playerImage, playerX, playerY, playerSize, playerSize);
}

function drawItems(items, image) {
    items.forEach(item => ctx.drawImage(image, item.x, item.y, item.width, item.height));
}

function movePlayer(dx, dy) {
    if (gameOver) return;

    playerX += dx;
    playerY += dy;

    // Ограничение по границам канваса
    if (playerX < 0) playerX = 0;
    if (playerX > canvas.width - playerSize) playerX = canvas.width - playerSize;
    if (playerY < 0) playerY = 0;
    if (playerY > canvas.height - playerSize) playerY = canvas.height - playerSize;

    fetch('/game', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ player_x: playerX, player_y: playerY })
    });
}

window.addEventListener('keydown', (e) => {
    if (gameOver) return;
    switch (e.key) {
        case 'ArrowLeft':
            movePlayer(-5, 0);
            break;
        case 'ArrowRight':
            movePlayer(5, 0);
            break;
        case 'ArrowUp':
            movePlayer(0, -5);
            break;
        case 'ArrowDown':
            movePlayer(0, 5);
            break;
    }
});

document.getElementById('startBtn').addEventListener('click', () => {
    score = 0;
    gameOver = false;
    document.getElementById('restartBtn').disabled = false;
    gameLoop();
});

document.getElementById('restartBtn').addEventListener('click', () => {
    score = 0;
    gameOver = false;
    document.getElementById('restartBtn').disabled = false;
    gameLoop();
});

function gameLoop() {
    if (!gameOver) {
        updateGame();
        requestAnimationFrame(gameLoop);
    }
}

gameLoop();
