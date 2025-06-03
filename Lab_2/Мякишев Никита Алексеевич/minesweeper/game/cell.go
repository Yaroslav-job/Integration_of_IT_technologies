package game

// Cell представляет ячейку на поле
type Cell struct {
	IsMine        bool // Мина или нет
	IsRevealed    bool // Открыта ли
	IsFlagged     bool // Помечена ли флагом
	NeighborMines int  // Кол-во мин вокруг
}
