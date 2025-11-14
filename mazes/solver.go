package maze

type PathFollower interface {
	Step(*Maze, Point)
	Done(*Maze, *Solution)
}

type Solver struct {
	maze       *Maze
	road_taken []Point
	solution   *Solution
	did_visit  map[Point]bool
	followers  []PathFollower
}

func (s *Solver) Maze() *Maze {
	return s.maze
}

func (s *Solver) Start() Point {
	return s.maze.Start
}

func (s *Solver) End() Point {
	return s.maze.End
}

func (s *Solver) At(row, col int) *Node {
	return &s.maze.Board[row][col]
}

func (s *Solver) Step(point Point) {
	s.road_taken = append(s.road_taken, point)
	s.did_visit[point] = true
	for _, f := range s.followers {
		f.Step(s.maze, point)
	}
}

func (s *Solver) Solved(solution *Solution) {
	if solution != nil {
		for _, f := range s.followers {
			f.Done(s.maze, solution)
		}
	}
}

func (s *Solver) Solution() *Solution {
	return s.solution
}

func (s *Solver) DidVisit(point Point) bool {
	_, exists := s.did_visit[point]
	return exists
}

func (s *Solver) RegisterFollower(follower PathFollower) {
	s.followers = append(s.followers, follower)
}
