package maze

import (
	"fmt"
	"slices"
	"strings"
)

type Solution struct {
	path []Point
}

func (s *Solution) Add(point Point) {
	s.path = append(s.path, point)
}

func (s *Solution) String() string {
	slices.Reverse(s.path)
	var sb strings.Builder
	for _, p := range s.path {
		sb.WriteString(fmt.Sprintf("(%d, %d)", p.Row, p.Col))
	}
	return sb.String()
}

func (s *Solution) Contains(p Point) bool {
	for _, c := range s.path {
		if c.Equals(p) {
			return true
		}
	}
	return false
}
