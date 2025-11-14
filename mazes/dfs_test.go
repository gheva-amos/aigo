package maze_test

import (
	"github.com/gheva-amos/aigo/mazes"
	"testing"
)

func TestPopNode(t *testing.T) {
	str := `##B   #
## ## #
#  #  #
# ## ##
     ##
A######`
	var m maze.Maze
	m.FromString(str)
	dfs := maze.NewDFS(&m)
	dfs.Frontier = append(dfs.Frontier, dfs.Start())
	node, err := dfs.PopNode()
	if err != nil {
		t.Errorf("Got error instead of node %e", err)
	}
	if !node.Coords.Equals(dfs.Start()) {
		t.Errorf("Nodes are not in the same place")
	}
	node, err = dfs.PopNode()
	if err == nil {
		t.Errorf("Expected to get an error")
	}
}

func TestSolve(t *testing.T) {
	str := `##B   #
## ## #
#  #  #
# ## ##
     ##
A######`
	var m maze.Maze
	m.FromString(str)
	dfs := maze.NewDFS(&m)
	dfs.Solve()
	m.ToImage(dfs.Solution(), "out.png")
}
