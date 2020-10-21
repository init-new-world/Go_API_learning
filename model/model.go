package model

import "sync"

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint]*User
}

type Token struct {
	Token string `json:"token"`
}
