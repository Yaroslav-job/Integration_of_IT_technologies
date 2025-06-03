package game

import (
	"sync"
	"time"

	"fyne.io/fyne/v2"
)

var WindowSize = fyne.NewSize(800, 600)
var ExternalWindowSize = fyne.NewSize(813, 613)

var BallSpeed float64 = 4.0

var activeBalls []*Ball

var activeBallsLock sync.Mutex

var ballSpawnTicker *time.Ticker

var maxBalls = 5
var BrickRows = 8

var BrickColumns = 10

var BallSize = fyne.NewSize(25, 25)

var PaddleSize = fyne.NewSize(170, 20)
