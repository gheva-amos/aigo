package maze

import (
	"fmt"
)

type DijkstraSearch struct {
	Solver
	Frontier *PriorityQueue
	Parents  map[*Node]*Node
	current  *Node
}

func NewDijkstra(maze *Maze) DijkstraSearch {
	return DijkstraSearch{Solver: Solver{
		maze:      maze,
		solution:  &Solution{},
		did_visit: make(map[Point]bool),
	},
		Frontier: NewPriorityQueue(false),
		current: nil, Parents: make(map[*Node]*Node)}
}

func (bfs *DijkstraSearch) Solve() error {
	bfs.Frontier.Push(bfs.Start(), 0)
	bfs.Parents[&bfs.Maze().Board[bfs.Start().Row][bfs.Start().Col]] = bfs.current

	for {
		if bfs.Frontier.Len() == 0 {
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
					bfs.Frontier.Push(n, bfs.Distance(n))
					bfs.Parents[as_node] = bfs.current
				}
			}
		}
	}
	return nil
}

func (bfs *DijkstraSearch) in_frontier(point Point) bool {
	return bfs.Frontier.Find(point)
}

func (bfs *DijkstraSearch) PopNode() (*Node, error) {
	if bfs.Frontier.Len() == 0 {
		return nil, fmt.Errorf("Trying to get a node from an empty frontier")
	}
	p := bfs.Frontier.Pop()
	ret := bfs.At(p.Row, p.Col)

	return ret, nil
}

func (bfs *DijkstraSearch) Distance(point Point) int {
	return abs(bfs.Start().Col - point.Col) + abs(bfs.Start().Row - point.Row)
}

func abs(x int) int {
	 if x < 0 {
		 return -x
	 }
	 return x
}

