package tetris

import "time"

type gameSate int

const (
	gameInit gameSate = iota
	gameRunning
	gameOver
)

type Game struct {
	bdHeight  int
	bdWidth   int
	board     [][]int
	position  vector
	piece     piece
	state     gameSate
	FallSpeed *time.Timer
}

func (g *Game) collision() bool {
	for _, v := range g.piece.shape {
		b := g.blockOnBoardByPosition(v)
		if b.y >= g.bdHeight || b.x >= g.bdWidth || b.x < 0 || b.y < 0 || g.board[b.y][b.x] > 0 {
			return true
		}
	}

	return false
}

func (g *Game) Rotate() {
	g.piece.rotate()
	if g.collision() {
		g.piece.rotateBack()
	}
}

func (g *Game) SpeedUp() {
	if g.state != gameRunning {
		return
	}
	g.FallSpeed.Reset(50)
}

func (g *Game) MoveLeft() {
	g.moveIfPossible(vector{0, -1})
}

func (g *Game) MoveRight() {
	g.moveIfPossible(vector{0, 1})
}

func (g *Game) Fall() {
	if g.state != gameRunning {
		return
	}
	for g.moveIfPossible(vector{1, 0}) {
		if g.collision() {
			break
		}
		g.FallSpeed.Reset(1 * time.Millisecond)
	}
}

func (g *Game) moveIfPossible(v vector) bool {
	if g.state != gameRunning {
		return false
	}
	g.position.x += v.x
	g.position.y += v.y

	if g.collision() {
		g.position.y -= v.y
		g.position.x -= v.x
		return false
	}

	return true
}

func (g *Game) LockPiece() {
	g.board = g.GetBoard()
}

func (g *Game) isLineFull(y int) bool {
	for _, v := range g.board[y] {
		if v == 0 {
			return false
		}
	}

	return true
}

func (g *Game) removeLine(y int) {
	for i := y; i > 0; i-- {
		for j := 0; j < g.bdWidth; j++ {
			g.board[i][j] = g.board[i-1][j]
		}
	}

	for j := 0; j < g.bdWidth; j++ {
		g.board[0][j] = 0
	}
}

func (g *Game) removeLines() {
	for i := len(g.board) - 1; i >= 0; i-- {
		if g.isLineFull(i) {
			g.removeLine(i)
			i++
		}
	}
}

func (g *Game) GameLoop() {
	if !g.moveIfPossible(vector{1, 0}) {
		g.LockPiece()
		g.removeLines()
		g.getPiece()
		if g.collision() {
			g.state = gameOver
			g.FallSpeed.Stop()
			return
		}
	}
	g.ResetFallSpeed()
}

func (g *Game) blockOnBoardByPosition(v vector) vector {
	return vector{g.position.y + v.y, g.position.x + v.x}
}

func (g *Game) getPiece() {
	g.piece = randomPieces()
	g.position = vector{1, g.bdWidth / 2}
}

func (g *Game) ResetFallSpeed() {
	g.FallSpeed.Reset(700 * time.Millisecond)
}

func (g *Game) Start() {
	g.state = gameRunning
	g.getPiece()
	g.ResetFallSpeed()
}

func (g *Game) GetBoard() [][]int {
	cBoard := make([][]int, len(g.board))
	for i := range g.board {
		cBoard[i] = make([]int, len(g.board[i]))
		for j := range g.board[i] {
			cBoard[i][j] = g.board[i][j]
		}
	}

	for _, v := range g.piece.shape {
		b := g.blockOnBoardByPosition(v)
		cBoard[b.y][b.x] = g.piece.color
	}

	return cBoard
}

func (g *Game) init(bdHeight, bdWidth int) {
	g.board = make([][]int, bdHeight)
	for i := range g.board {
		g.board[i] = make([]int, bdWidth)
		for j := range g.board[i] {
			g.board[i][j] = 0
		}
	}
	g.bdWidth = bdWidth
	g.bdHeight = bdHeight
	g.position = vector{1, bdWidth / 2}
	g.piece = pieces[0]
	g.FallSpeed = time.NewTimer(10 * time.Millisecond)
	g.FallSpeed.Stop()
	g.state = gameInit
}

func NewGame(bdHeight, bdWidth int) *Game {
	g := &Game{}
	g.init(bdHeight, bdWidth)
	return g
}
