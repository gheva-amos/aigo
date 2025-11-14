package maze

import (
	"math/rand"
	"os"
	"strings"
)

type Maze struct {
	Board [][]Node
	Start Point
	End   Point
}

type Point struct {
	Row int
	Col int
}

type Node struct {
	Coords Point
	IsWall bool
}

func (p Point) Equals(other Point) bool {
	return p.Row == other.Row && p.Col == other.Col
}

func (m *Maze) Height() int {
	return len(m.Board)
}

func (m *Maze) Width() int {
	return len(m.Board[0])
}

func (m *Maze) FromString(desc string) {
	lines := strings.Split(desc, "\n")
	m.Board = make([][]Node, len(lines))
	for row, line := range lines {
		m.Board[row] = make([]Node, len(line))
		for col, r := range line {
			tmp := Node{Coords: Point{Row: row, Col: col}}
			switch r {
			case '#':
				tmp.IsWall = true
			case 'A':
				m.Start = Point{Row: row, Col: col}
			case 'B':
				m.End = Point{Row: row, Col: col}
			}
			m.Board[row][col] = tmp
		}
	}
}

func (m *Maze) FromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	str := string(data)
	m.FromString(str)
	return nil
}

func (n *Node) Neighbours(m *Maze) []Point {
	candidates := []Point{
		Point{Row: n.Coords.Row - 1, Col: n.Coords.Col},
		Point{Row: n.Coords.Row + 1, Col: n.Coords.Col},
		Point{Row: n.Coords.Row, Col: n.Coords.Col - 1},
		Point{Row: n.Coords.Row, Col: n.Coords.Col + 1},
	}

	var ret []Point
	for _, p := range candidates {
		if p.Row >= 0 && p.Row < m.Height() && p.Col >= 0 && p.Col < m.Width() {
			if !m.Board[p.Row][p.Col].IsWall {
				ret = append(ret, p)
			}
		}
	}
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}
