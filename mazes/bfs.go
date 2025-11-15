package maze

import (
	"fmt"
)

type BreadthFistSearch struct {
	Solver
	Frontier []Point
	Parents  map[*Node]*Node
	current  *Node
}

func NewBFS(maze *Maze) BreadthFistSearch {
	return BreadthFistSearch{Solver: Solver{
		maze:      maze,
		solution:  &Solution{},
		did_visit: make(map[Point]bool),
	},
		current: nil, Parents: make(map[*Node]*Node)}
}

func (bfs *BreadthFistSearch) Solve() error {
	bfs.Frontier = append(bfs.Frontier, bfs.Start())
	bfs.Parents[&bfs.Maze().Board[bfs.Start().Row][bfs.Start().Col]] = bfs.current

	for {
		if len(bfs.Frontier) == 0 {
			break
		}

		node, err := bfs.PopNode()
		if err != nil {
			return err
		}
		if node.Coords.Equals(bfs.End()) {
			fmt.Println("solved")
			for {
				if node == nil {
					break
				}
				bfs.solution.Add(node.Coords)
				node = bfs.Parents[node]
			}
			fmt.Printf("---\n%s\n", bfs.solution)
			bfs.Solved(bfs.solution)
			return nil
		}
		bfs.current = node
		bfs.Step(node.Coords)
		neighbours := node.Neighbours(bfs.Maze())
		for _, n := range neighbours {
			as_node := bfs.At(n.Row, n.Col)
			if !bfs.DidVisit(n) {
				if !bfs.in_frontier(n) {
					bfs.Frontier = append(bfs.Frontier, n)
					bfs.Parents[as_node] = bfs.current
				}
			}
		}
	}
	return nil
}

func (bfs *BreadthFistSearch) in_frontier(point Point) bool {
	for _, n := range bfs.Frontier {
		if point.Equals(n) {
			return true
		}
	}
	return false
}

func (bfs *BreadthFistSearch) PopNode() (*Node, error) {
	if len(bfs.Frontier) == 0 {
		return nil, fmt.Errorf("Trying to get a node from an empty frontier")
	}
	p := bfs.Frontier[0]
	ret := bfs.At(p.Row, p.Col)
	bfs.Frontier = bfs.Frontier[1:]

	return ret, nil
}
