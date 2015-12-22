package object

import (
	"github.com/disintegration/imaging"
	"image"
)

type Piece struct {
	Id          int
	south       int
	east        int
	north       int
	west        int
	Orientation int
	image       image.Image
	Score       float64
}

func (p *Piece) North() int {
	if p.Orientation == 1 {
		return p.west
	} else if p.Orientation == 2 {
		return p.south
	} else if p.Orientation == 3 {
		return p.east
	}
	return p.north
}

func (p *Piece) East() int {
	if p.Orientation == 1 {
		return p.north
	} else if p.Orientation == 2 {
		return p.west
	} else if p.Orientation == 3 {
		return p.south
	}
	return p.east
}

func (p *Piece) South() int {
	if p.Orientation == 1 {
		return p.east
	} else if p.Orientation == 2 {
		return p.north
	} else if p.Orientation == 3 {
		return p.west
	}
	return p.south
}

func (p *Piece) West() int {
	if p.Orientation == 1 {
		return p.south
	} else if p.Orientation == 2 {
		return p.east
	} else if p.Orientation == 3 {
		return p.north
	}
	return p.west
}

func (p *Piece) HasBorder() bool {
	if p.north == 0 || p.east == 0 ||
		p.south == 0 || p.west == 0 {
		return true
	}
	return false
}

func (p *Piece) Rotate(angle int) {
	p.Orientation = (p.Orientation + angle) % 4
}

func (p *Piece) Render() image.Image {

	if p.Orientation == 1 {
		return imaging.Rotate270(p.image)
	} else if p.Orientation == 2 {
		return imaging.Rotate180(p.image)
	} else if p.Orientation == 3 {
		return imaging.Rotate90(p.image)
	}
	return p.image
}
