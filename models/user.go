package models

type User struct {
	name string
	piece string
}

func NewUser(name string, piece string) *User{
	return &User{
		name: name,
		piece: piece,
	}
}

func (u *User)GetName()string{
	return u.name
}

func (u *User)GetPiece()string{
	return u.piece
}
