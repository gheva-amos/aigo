package maze

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

const (
	Empty uint8 = iota
	Wall
	Current
	Visited
	Start
	End
	SolutionColor
)

type Animator struct {
	palette   color.Palette
	did_visit map[Point]bool
	frames    []*image.Paletted
	delays    []int
	cell_size int
	current   Point
	solution  *Solution
}

func NewAnimator() *Animator {
	return &Animator{
		palette: color.Palette{
			color.White,                    // Empty
			color.Black,                    // Wall
			color.RGBA{255, 0, 0, 255},     // Current
			color.RGBA{0, 255, 0, 120},     // Visited
			color.RGBA{0, 255, 255, 255},   // Start
			color.RGBA{0, 0, 255, 255},     // End
			color.RGBA{120, 120, 255, 255}, // SolutionColor
		},
		did_visit: make(map[Point]bool),
		cell_size: 60,
	}
}

func (a *Animator) Step(maze *Maze, point Point) {
	a.current = point

	a.add_frame(maze)
	a.did_visit[point] = true
}

func (a *Animator) add_frame(maze *Maze) {
	rect := image.Rect(0, 0, maze.Width()*a.cell_size, maze.Height()*a.cell_size)
	img := image.NewPaletted(rect, a.palette)
	for row, row_elements := range maze.Board {
		for col, cell := range row_elements {
			color := Empty
			if cell.IsWall {
				color = Wall
			}
			if _, exists := a.did_visit[cell.Coords]; exists {
				color = Visited
			}
			if cell.Coords.Equals(maze.Start) {
				color = Start
			}
			if a.solution != nil && a.solution.Contains(cell.Coords) {
				color = SolutionColor
			}
			if cell.Coords.Equals(maze.End) {
				color = End
			}
			if cell.Coords.Equals(a.current) {
				color = Current
			}
			fill_rect(img, col*a.cell_size, row*a.cell_size, a.cell_size, a.cell_size, color)
		}
	}
	a.frames = append(a.frames, img)
	a.delays = append(a.delays, 5)
}

func fill_rect(img *image.Paletted, x, y, w, h int, idx uint8) {
	for yy := y; yy < y+h; yy++ {
		for xx := x; xx < x+w; xx++ {
			img.SetColorIndex(xx, yy, idx)
		}
	}
}

func (a *Animator) Done(maze *Maze, solution *Solution) {
	a.solution = solution
	a.add_frame(maze)
}

func (a *Animator) Save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	a.delays[len(a.delays)-1] = 300
	anim := gif.GIF{
		Image: a.frames,
		Delay: a.delays,
	}
	return gif.EncodeAll(f, &anim)
}
