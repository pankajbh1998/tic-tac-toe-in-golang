package service

import (
	"TicTacToe/repo"
	"fmt"
)

type GameManager struct{
	g *repo.Game
	u *repo.User
}

func NewGameService(g *repo.Game, u *repo.User)*GameManager{
	return &GameManager{
		g: g,
		u: u,
	}
}

func( g *GameManager)GetPosition(row int, col int) int {
	return (row-1) * g.GetHorizontalSize() + (col-1)
}

func (g *GameManager)AddPieceToPosition(pos int, piece string) error {
	if g.g.IsPositionAlreadyFilled(pos) {
		return fmt.Errorf("Invalid Move")
	}

	g.g.AddPieceToPos(pos, piece)
	return nil
}

func (g *GameManager)GetHorizontalSize() int {
	return g.g.GetRow()
}

func (g *GameManager)GetPiece(pos int) (string, bool) {
	return g.g.GetPiece(pos)
}

func (g *GameManager)GetVerticalSize() int {
	return g.g.GetCol()
}

func (g *GameManager)CheckForWinner(pos int, piece string) bool {
	size := g.GetVerticalSize()
	if g.checkForHorizontalWinner(pos, size, piece) ||
		g.checkForVerticalWinner(pos, size, piece) ||
		g.checkForDiagonalWinner(0, size+1, piece) ||
		g.checkForDiagonalWinner(size-1, size-1, piece) {
		return true
	}

	return false

}

func (g *GameManager)checkForHorizontalWinner(pos int, size int, piece string) bool {
	row := pos/size
	for i:=0; i<size;i++ {
		currPiece, ok := g.g.GetPiece(row*size + i)
		if !ok || currPiece != piece{
			return false
		}
	}
	return true
}

func (g *GameManager)checkForVerticalWinner(pos int, size int, piece string) bool {
	row := pos%size
	for i:=0; i<size;i++ {
		currPiece, ok := g.g.GetPiece(i*size + row)
		if !ok || currPiece != piece{
			return false
		}
	}
	return true
}

func (g *GameManager)checkForDiagonalWinner(pos int, size int, piece string) bool {
	for i:=0; i<size;i++ {
		currPiece, ok := g.g.GetPiece(pos)
		if !ok || currPiece != piece{
			return false
		}
		pos += size
	}
	return true
}