package maze

import (
	"fmt"
)

type DepthFirstSearch struct {
	Solver
	Frontier []Point
	Parents  map[*Node]*Node
	current  *Node
}

func NewDFS(maze *Maze) DepthFirstSearch {
	return DepthFirstSearch{Solver: Solver{
		maze:      maze,
		solution:  &Solution{},
		did_visit: make(map[Point]bool),
	},
		current: nil, Parents: make(map[*Node]*Node)}
}

func (dfs *DepthFirstSearch) Solve() error {
	dfs.Frontier = append(dfs.Frontier, dfs.Start())
	dfs.Parents[&dfs.Maze().Board[dfs.Start().Row][dfs.Start().Col]] = dfs.current

	for {
		if len(dfs.Frontier) == 0 {
			break
		}

		node, err := dfs.PopNode()
		if err != nil {
			return err
		}
		if node.Coords.Equals(dfs.End()) {
			dfs.Parents[node] = dfs.current
			fmt.Println("solved")
			for {
				if node == nil {
					break
				}
				dfs.solution.Add(node.Coords)
				node = dfs.Parents[node]
			}
			fmt.Printf("---\n%s\n", dfs.solution)
			return nil
		}
		dfs.current = node
		dfs.Step(node.Coords)
		neighbours := node.Neighbours(dfs.Maze())
		for _, n := range neighbours {
			as_node := &dfs.Maze().Board[n.Row][n.Col]
			if !dfs.DidVisit(n) {
				if !dfs.in_frontier(as_node) {
					dfs.Frontier = append(dfs.Frontier, n)
					dfs.Parents[as_node] = dfs.current
				}
			}
		}
	}
	return nil
}

func (dfs *DepthFirstSearch) in_frontier(node *Node) bool {
	for _, n := range dfs.Frontier {
		if node.Coords.Equals(n) {
			return true
		}
	}
	return false
}

func (dfs *DepthFirstSearch) PopNode() (*Node, error) {
	if len(dfs.Frontier) == 0 {
		return nil, fmt.Errorf("Trying to get a node from an empty frontier")
	}
	p := dfs.Frontier[len(dfs.Frontier)-1]
	ret := &dfs.Maze().Board[p.Row][p.Col]
	dfs.Frontier = dfs.Frontier[:len(dfs.Frontier)-1]

	return ret, nil
}
