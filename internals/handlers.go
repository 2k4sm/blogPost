package internals

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ViewAll(c *fiber.Ctx) error {

	db, err := gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect Database.")
	}

	var posts []*Post
	rows, err := db.Model(&Post{}).Rows()

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		db.ScanRows(rows, &posts)
	}

	P := map[string][]*Post{
		"Posts": posts,
	}

	return c.Render("index", P)

}

func CreatePost(c *fiber.Ctx) error {
	return c.SendString("hello")
}
