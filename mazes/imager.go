package maze

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

const cell_size = 60

var (
	empty   = color.RGBA{B: 255, A: 255}
	in_path = color.RGBA{G: 150, A: 150}
	wall    = color.RGBA{A: 255}
	start   = color.RGBA{G: 255, A: 255}
	end     = color.RGBA{R: 255, A: 255}
)

func (m *Maze) ToImage(solution *Solution, path string) {
	canvas := image.NewRGBA(image.Rect(0, 0, m.Height()*cell_size, m.Width()*cell_size))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{color.NRGBA{30, 30, 30, 255}}, image.Point{}, draw.Src)
	for row, row_elements := range m.Board {
		for col, cell := range row_elements {
			color := empty
			if solution.Contains(cell.Coords) {
				color = in_path
			}
			if cell.IsWall {
				color = wall
			}
			if cell.Coords.Equals(m.Start) {
				color = start
			}
			if cell.Coords.Equals(m.End) {
				color = end
			}
			m.draw_rect(col, row, cell_size, canvas, color)
		}
	}
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	png.Encode(f, canvas)
}

func (m *Maze) draw_rect(x, y, size int, canvas *image.RGBA, color color.RGBA) {
	cell := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(cell, cell.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)
	draw.Draw(canvas, image.Rect(x*size, y*size, (x+1)*size, (y+1)*size), cell, image.Point{}, draw.Src)
}
