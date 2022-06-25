package main

import (
	"TicTacToe/models"
	"TicTacToe/repo"
	"TicTacToe/service"
)

func main(){
	boardSize := 3
	// square board
	gameRepo := repo.NewGame(boardSize, boardSize)
	userRepo := addNewUsers(repo.NewUser())
	gameManager := service.NewGameService(gameRepo, userRepo)
	commands := getCommands()
	RunCommands(gameManager, userRepo, commands)
}

func getCommands() [][]int {
	commands1 := [][]int{
		{2, 2},
		{1, 3},
		{1, 1},
		{1, 2},
		{2, 2},
		{3, 3},
	}
	return commands1
	commands2 := [][]int{
		{2, 3},
		{1, 2},
		{2, 2},
		{2, 1},
		{1, 1},
		{3, 3},
		{3, 2},
		{3, 1},
		{1, 3},
	}
	return commands2
}

func addNewUsers(userRepo *repo.User) *repo.User {
	user1 := models.NewUser("Pankaj", "X")
	user2 := models.NewUser("Shivam", "O")
	userRepo.AddUser(user1)
	userRepo.AddUser(user2)
	return userRepo
}

