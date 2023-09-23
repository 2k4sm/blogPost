package model

import "time"

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Post struct {
	Id          uint64
	Title       string
	Author      string
	Description string
	DateCreated time.Time `objectbox:"date"`
}
