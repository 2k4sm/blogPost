package main

import (
	"fmt"
	"time"

	"github.com/2k4sm/blogPost/internals/model"
	"github.com/gofiber/fiber/v2"
	"github.com/objectbox/objectbox-go/objectbox"
)

type Post struct {
	Id          uint64
	Title       string
	Author      string
	Description string
	DateCreated time.Time
}

func initObjectBox() *objectbox.ObjectBox {
	objectBox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Build()

	if err != nil {
		panic(fmt.Sprintf("Error initialising Object Box: %v", err))
	}

	return objectBox
}

func viewAll(c *fiber.Ctx) error {
	ob := initObjectBox()
	defer ob.Close()

	box := model.BoxForPost(ob)

	posts, _ := box.GetAll()

	P := map[string][]*model.Post{
		"Posts": posts,
	}

	return c.Render("index", P)

}
