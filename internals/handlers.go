package internals

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RenderHome(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func ViewAll(c *fiber.Ctx) error {

	db, err := gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect Database.")
	}
	db.AutoMigrate(Post{})

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

	return c.JSON(P)

}

func CreatePost(c *fiber.Ctx) error {
	return c.Render("publish", nil)
}

func ProcessForm(c *fiber.Ctx) error {

	db, err := gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect Database.")
	}

	desc := c.FormValue("description")
	post := &Post{Description: desc}
	db.Create(post)

	fmt.Println("Form Processed Successfully.")

	return c.Redirect("/")
}
