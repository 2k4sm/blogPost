package internals

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ViewAll(c *fiber.Ctx) error {

	db, err := gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect Database.")
	}
	db.AutoMigrate(Post{})

	//db.Create(&Post{Title: "md test", Author: "md author", Description: "### working with md\n # this is a big tag\n - point1\n - point2\n [link](https://www.google.com)"})

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

func ProcessForm(c *fiber.Ctx) error {

	db, err := gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect Database.")
	}

	title := c.FormValue("title")
	author := c.FormValue("author")
	desc := c.FormValue("description")
	post := &Post{Title: title, Author: author, Description: desc}
	db.Create(post)

	fmt.Println("Form Processed Successfully.")

	return c.RedirectToRoute("/", nil)
}
