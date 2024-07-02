package main

import (
	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Password string
}

type Todo struct {
	ID     int
	UserID int
	Text   string
}
