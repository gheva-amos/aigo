package main

import (
	"flag"
	"fmt"
	"github.com/gheva-amos/aigo/mazes"
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
	a := maze.NewAnimator()
	dfs := maze.NewDFS(&m)
	dfs.RegisterFollower(a)
	dfs.Solve()
	a.Save(*save_to)
	//m.ToImage(dfs.Solution(), *save_to)
}
