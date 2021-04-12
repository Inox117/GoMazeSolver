package main

import (
	"fmt"
	"github.com/Inox117/GoMazeSolver/internal/model"
	"github.com/Inox117/GoMazeSolver/internal/service"
)

func main() {
	fmt.Println("Starting to solve a Maze")
	graph := model.CreateEmptyGraph()
	err := service.FindExit(graph)
	if err != nil {
		fmt.Printf("Something wrong happened : %s", err.Error())
	}
}
