package repo

type Game struct {
	mp map[int]string
	row int
	col int
}

func NewGame(row int, col int)*Game{
	return &Game{
		mp: make(map[int]string),
		row: row,
		col: col,
	}
}

func (g *Game)AddPieceToPos(pos int, piece string) {
	g.mp[pos] = piece
}

func (g *Game)GetPiece(pos int)(string, bool) {
	piece,ok := g.mp[pos]
	if !ok {
		return "", false
	}
	return piece, true
}

func (g *Game)IsPositionAlreadyFilled(pos int) bool {
	_, ok := g.mp[pos]
	if !ok {
		return false
	}
	return true
}

func (g *Game)GetRow() int {
	return g.row
}

func (g *Game)GetCol() int {
	return g.col
}
