package tetris

import (
	"math"
	"math/rand"
)

type piece struct {
	shape     []vector
	canRotate bool
	color     int
}

func (p *piece) rotateBack() {
	ang := math.Pi / 2 * 3
	p.rotateWithAngle(ang)
}

func (p *piece) rotate() {
	ang := math.Pi / 2
	p.rotateWithAngle(ang)
}

func (p *piece) rotateWithAngle(ang float64) {
	if !p.canRotate {
		return
	}
	cos := int(math.Round(math.Cos(ang)))
	sin := int(math.Round(math.Sin(ang)))

	for i, e := range p.shape {
		ny := e.y*cos - e.x*sin
		nx := e.y*sin - e.x*cos

		p.shape[i] = vector{ny, nx}
	}
}

var pieces = []piece{
	{
		shape: []vector{
			{0, 0},
		},
		canRotate: false,
		color:     0,
	},
	{
		// L shape
		shape: []vector{
			{0, -1},
			{0, 0},
			{0, 1},
			{1, 1},
		},
		canRotate: true,
		color:     1,
	},
	{
		// opposite L shape
		shape: []vector{
			{0, -1},
			{0, 0},
			{0, 1},
			{-1, 1},
		},
		canRotate: true,
		color:     2,
	},
	{
		// I shape
		shape: []vector{
			{0, -1},
			{0, 0},
			{0, 1},
			{0, 2},
		},
		canRotate: true,
		color:     3,
	},
	{
		// O shape
		shape: []vector{
			{1, -1},
			{1, 0},
			{0, -1},
			{0, 0},
		},
		canRotate: false,
		color:     4,
	},
	{
		// T shape
		shape: []vector{
			{0, -1},
			{0, 0},
			{0, 1},
			{1, 0},
		},
		canRotate: true,
		color:     5,
	},
	{
		// S shape
		shape: []vector{
			{0, -1},
			{0, 0},
			{1, 0},
			{1, 1},
		},
		canRotate: true,
		color:     6,
	},
	{
		// Z shape
		shape: []vector{
			{0, 0},
			{0, 1},
			{1, -1},
			{1, 0},
		},
		canRotate: true,
		color:     7,
	},
}

func randomPieces() piece {
	idx := rand.Intn(len(pieces)-1) + 1
	pc := pieces[idx]
	return piece{
		shape:     append([]vector(nil), pc.shape...),
		canRotate: pc.canRotate,
		color:     pc.color,
	}
}
