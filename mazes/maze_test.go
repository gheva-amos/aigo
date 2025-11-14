package maze_test

import (
	"github.com/gheva-amos/aigo/mazes"
	"testing"
)

func TestPointEquals(t *testing.T) {
	p1 := maze.Point{Row: 0, Col: 5}
	p2 := maze.Point{Row: 0, Col: 5}
	if !p1.Equals(p2) {
		t.Errorf("they are not equal????")
	}
	if !p2.Equals(p1) {
		t.Errorf("they are not equal????")
	}
}

func TestFromString(t *testing.T) {

	str := `##B   #
## ## #
#  #  #
# ## ##
     ##
A######`
	var m maze.Maze
	m.FromString(str)
	if !m.Start.Equals(maze.Point{5, 0}) {
		t.Errorf("they are not equal????")
	}
	if !m.End.Equals(maze.Point{0, 2}) {
		t.Errorf("they are not equal????")
	}
	if !m.Board[0][0].IsWall {
		t.Errorf("Not a wall")
	}
	if m.Board[0][3].IsWall {
		t.Errorf("Not a wall")
	}
}

func TestNeighbours(t *testing.T) {

	str := `##B   #
## ## #
#  #  #
# ## ##
     ##
A######`
	var m maze.Maze
	m.FromString(str)
	n := &m.Board[5][0]
	nn := n.Neighbours(&m)
	if len(nn) != 1 {
		t.Errorf("too many neighbours")
	}
	if nn[0].Col != 0 || nn[0].Row != 4 {
		t.Errorf("wrong neighbour")
	}
}
