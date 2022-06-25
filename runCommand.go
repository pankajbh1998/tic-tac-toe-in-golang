package main

import (
	"TicTacToe/models"
	"TicTacToe/repo"
	"TicTacToe/service"
	"fmt"
)

func RunCommands(manager *service.GameManager, userManager *repo.User, commands [][]int){
	users,userSize := getAndPrintUsers(userManager)
	userTurn := 0
	for _,command := range commands{
		err := manager.AddPieceToPosition(manager.GetPosition(command[0], command[1]), users[userTurn].GetPiece())
		if err != nil {
			fmt.Println(err)
			printBoard(manager)
			continue
		}
		if manager.CheckForWinner(manager.GetPosition(command[0], command[1]), users[userTurn].GetPiece()) {
			fmt.Printf("%v won the game\n", users[userTurn].GetName())
			return
		}
		printBoard(manager)
		userTurn = (userTurn+1)%userSize
	}
	fmt.Println("Game Over")
}

func getAndPrintUsers(manager *repo.User)([]*models.User, int) {
	userCnt := manager.GetUserCount()
	users := make([]*models.User,0)
	for i:=0;i<userCnt;i++ {
		user,_ := manager.GetUser(i)
		fmt.Println(user.GetName(), user.GetPiece())
		users = append(users, user)
	}
	return users, userCnt
}

func printBoard(manager *service.GameManager){
	row := manager.GetHorizontalSize()
	col := manager.GetVerticalSize()
	for i:=1;i<=row;i++ {
		str := ""
		for j:=1;j<=col;j++{
			piece,ok := manager.GetPiece(manager.GetPosition(i, j))
			if !ok {
				piece = "-"
			}
			str = fmt.Sprintf("%v %v", str, piece)
		}
		fmt.Println(str)
	}
}
