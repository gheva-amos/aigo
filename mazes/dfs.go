package maze

import (
	"fmt"
)

type DepthFirstSearch struct {
	maze     *Maze
	Frontier []Point
	Parents  map[*Node]*Node
	visited  []Point
	current  *Node
	solution *Solution
}

func NewDFS(maze *Maze) DepthFirstSearch {
	return DepthFirstSearch{maze: maze, current: nil, Parents: make(map[*Node]*Node), solution: &Solution{}}
}

func (dfs *DepthFirstSearch) Solve() error {
	dfs.Frontier = append(dfs.Frontier, dfs.maze.Start)
	dfs.Parents[&dfs.maze.Board[dfs.maze.Start.Row][dfs.maze.Start.Col]] = dfs.current

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
		dfs.visited = append(dfs.visited, node.Coords)
		neighbours := node.Neighbours(dfs.maze)
		for _, n := range neighbours {
			as_node := &dfs.maze.Board[n.Row][n.Col]
			if !dfs.did_visit(as_node) {
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

func (dfs *DepthFirstSearch) did_visit(node *Node) bool {
	for _, n := range dfs.visited {
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
	ret := &dfs.maze.Board[p.Row][p.Col]
	dfs.Frontier = dfs.Frontier[:len(dfs.Frontier)-1]

	return ret, nil
}

func (dfs *DepthFirstSearch) Start() Point {
	return dfs.maze.Start
}

func (dfs *DepthFirstSearch) End() Point {
	return dfs.maze.End
}

func (dfs *DepthFirstSearch) Solution() *Solution {
	return dfs.solution
}
