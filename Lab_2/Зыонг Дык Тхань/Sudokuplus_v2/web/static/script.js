let gameState = {
    id: null,
    grid: null,
    originalGrid: null,
    level: 1,
    completed: false,
    selectedCell: null
};
const elements = {
    screens: {
        start: document.getElementById('start-screen'),
        instructions: document.getElementById('instructions-screen'),
        game: document.getElementById('game-screen'),
        completion: document.getElementById('completion-screen')
    },
    game: {
        board: document.querySelector('.game-board'),
        levelDisplay: document.getElementById('level-display'),
        resetBtn: document.getElementById('reset-btn'),
        exitBtn: document.getElementById('exit-game-btn')
    },
    dialogs: {
        message: document.getElementById('message-dialog'),
        messageTitle: document.getElementById('message-title'),
        messageText: document.getElementById('message-text')
    },
    completion: {
        level: document.getElementById('completion-level'),
        menuBtn: document.getElementById('back-to-menu-btn')
    }
};

document.addEventListener('DOMContentLoaded', function() {
    setupEventListeners();
    showScreen('start');
});
function setupEventListeners() {
    document.getElementById('level-selection-btn').addEventListener('click', () => {
        startNewGame(1);
    });
    
    document.getElementById('instructions-btn').addEventListener('click', () => {
        showScreen('instructions');
    });
    
    document.getElementById('back-from-instructions-btn').addEventListener('click', () => showScreen('start'));
    
    document.getElementById('reset-btn').addEventListener('click', resetGame);
    document.getElementById('exit-game-btn').addEventListener('click', exitToMainMenu);
    
    document.getElementById('back-to-menu-btn').addEventListener('click', exitToMainMenu);
    
    document.getElementById('message-ok-btn').addEventListener('click', () => hideDialog('message'));
    
    document.addEventListener('keydown', handleKeyPress);
}

function showScreen(screenName) {
    Object.values(elements.screens).forEach(screen => screen.classList.add('hidden'));
    elements.screens[screenName].classList.remove('hidden');
}

function showDialog(dialogName) {
    elements.dialogs[dialogName].classList.remove('hidden');
}

function hideDialog(dialogName) {
    elements.dialogs[dialogName].classList.add('hidden');
}

async function startNewGame(level) {
    resetGameState();
    gameState.level = level;
    
    try {
        const formData = new FormData();
        formData.append('level', level);
        
        const response = await fetch('/api/game/new', {
            method: 'POST',
            body: formData
        });
        
        if (!response.ok) {
            throw new Error('Failed to create new game');
        }
        
        const data = await response.json();
        
        gameState.id = data.id;
        gameState.grid = data.grid;
        gameState.originalGrid = data.originalGrid;
        gameState.level = data.level;
        gameState.completed = data.completed;
        
        showScreen('game');
        renderGame();
    } catch (error) {
        console.error('Error starting new game:', error);
        showMessage('Error', 'Failed to start a new game. Please try again.');
    }
}
function resetGameState() {
    gameState = {
        id: null,
        grid: null,
        originalGrid: null,
        level: 1,
        completed: false,
        selectedCell: null
    };
}

function renderGame() {
    elements.game.board.innerHTML = '';
    updateGameInfo();
    
    for (let row = 0; row < 9; row++) {
        for (let col = 0; col < 9; col++) {
            const cell = document.createElement('div');
            cell.className = 'cell';
            cell.setAttribute('data-row', row);
            cell.setAttribute('data-col', col);
            
            if (row === 2 || row === 5) {
                cell.style.borderBottom = '2px solid #34495e';
            }
            if (col === 2 || col === 5) {
                cell.style.borderRight = '2px solid #34495e';
            }
            
            const cellValue = gameState.grid[row][col];
            
            if (gameState.originalGrid[row][col] !== 0) {
                cell.classList.add('original');
            }
            
            if (cellValue !== 0) {
                cell.textContent = cellValue;
            }
            
            cell.addEventListener('click', () => selectCell(row, col, cell));
            elements.game.board.appendChild(cell);
        }
    }
}

function updateGameInfo() {
    elements.game.levelDisplay.textContent = `Level: ${gameState.level}/10`;
}

function selectCell(row, col, cellElement) {
    const previouslySelected = document.querySelector('.cell.selected');
    if (previouslySelected) {
        previouslySelected.classList.remove('selected');
    }
    
    cellElement.classList.add('selected');
    gameState.selectedCell = { row, col };
}

async function makeMove(number) {
    if (!gameState.selectedCell) {
        return;
    }
    
    const { row, col } = gameState.selectedCell;
    
    if (gameState.originalGrid[row][col] !== 0) {
        return;
    }
    
    try {
        const formData = new FormData();
        formData.append('gameId', gameState.id);
        formData.append('row', row);
        formData.append('col', col);
        formData.append('num', number);
        
        const response = await fetch('/api/game/move', {
            method: 'POST',
            body: formData
        });
        
        if (!response.ok) {
            throw new Error('Failed to make move');
        }
        
        const data = await response.json();
        
        if (data.success) {
            gameState.grid[row][col] = number;
            
            const cell = getCellElement(row, col);
            cell.textContent = number !== 0 ? number : '';
            cell.classList.remove('wrong');
            
            if (data.completed) {
                gameState.completed = true;
                handleGameCompletion();
            }
        } else {
            const cell = getCellElement(row, col);
            if (number !== 0) {
                cell.textContent = number;
                cell.classList.add('wrong');
                
                setTimeout(() => {
                    if (gameState.grid[row][col] === 0) {
                        cell.textContent = '';
                        cell.classList.remove('wrong');
                    }
                }, 1000);
            }
        }
    } catch (error) {
        console.error('Error making move:', error);
        showMessage('Error', 'Failed to make the move. Please try again.');
    }
}

function handleGameCompletion() {
    elements.completion.level.textContent = `Level: ${gameState.level}/10`;
    
    if (gameState.level < 10) {
        setTimeout(() => {
            startNewGame(gameState.level + 1);
        }, 1500);
    } else {
        showScreen('completion');
    }
}



function resetGame() {
    gameState.grid = JSON.parse(JSON.stringify(gameState.originalGrid));
    gameState.completed = false;
    gameState.selectedCell = null;
    
    showScreen('game');
    renderGame();
}

function exitToMainMenu() {
    resetGameState();
    showScreen('start');
}

function showMessage(title, message) {
    elements.dialogs.messageTitle.textContent = title;
    elements.dialogs.messageText.textContent = message;
    showDialog('message');
}

function getCellElement(row, col) {
    return document.querySelector(`.cell[data-row="${row}"][data-col="${col}"]`);
}

function handleKeyPress(event) {
    if (elements.screens.game.classList.contains('hidden')) {
        return;
    }
    
    if (event.key >= '1' && event.key <= '9') {
        makeMove(parseInt(event.key));
    }
    else if (event.key === '0' || event.key === 'Delete' || event.key === 'Backspace') {
        makeMove(0);
    }
    else if (event.key.startsWith('Arrow')) {
        if (!gameState.selectedCell) {
            selectCell(0, 0, getCellElement(0, 0));
            return;
        }
        
        let { row, col } = gameState.selectedCell;
        
        switch (event.key) {
            case 'ArrowUp':
                row = Math.max(0, row - 1);
                break;
            case 'ArrowDown':
                row = Math.min(8, row + 1);
                break;
            case 'ArrowLeft':
                col = Math.max(0, col - 1);
                break;
            case 'ArrowRight':
                col = Math.min(8, col + 1);
                break;
        }
        
        const cell = getCellElement(row, col);
        selectCell(row, col, cell);
    }
}