package main

import (
	"github.comr/gheva-amos/aigo/mazes"
	"flag"
	"fmt"
	"os"
)

func main() {

	maze_file := flag.String("maze", "", "The path to the maze file")

	save_to := flag.String("save", "", "The path the resulting image")
	flag.Parse()

	var m maze.Maze
	err := m.FromFile(*maze_file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dfs := maze.NewDFS(&m)
	dfs.Solve()
	m.ToImage(dfs.Solution(), *save_to)
}
