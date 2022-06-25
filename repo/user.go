package repo

import "TicTacToe/models"

type User struct {
	mp map[int]*models.User
	size int
}

func NewUser()*User{
	return &User{
		mp: make(map[int]*models.User),
		size: 0,
	}
}

func (u *User)AddUser(user *models.User) {
	u.mp[u.size] = user
	u.size++
}

func (u *User)GetUser(id int) (*models.User, bool) {
	user,ok := u.mp[id]
	if !ok {
		return nil, false
	}
	return user, true
}

func (u *User)GetUserCount() int {
	return u.size
}