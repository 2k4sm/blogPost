package main

import (
	"github.com/2k4sm/blogPost/internals/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func viewAll(c *fiber.Ctx) error {

	db, err := gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect Database.")
	}

	var posts []*model.Post
	rows, err := db.Model(&model.Post{}).Rows()

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		db.ScanRows(rows, &posts)
	}

	P := map[string][]*model.Post{
		"Posts": posts,
	}

	return c.Render("index", P)

}

// func createPost(c *fiber.Ctx) error {

// }
