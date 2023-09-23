package main

import (
	"encoding/json"
	"fmt"

	"github.com/2k4sm/blogPost/internals/model"
	"github.com/gofiber/fiber/v2"
	"github.com/objectbox/objectbox-go/objectbox"
)

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

	p, err := json.Marshal(posts)

	if err != nil {
		panic(err)
	}

	c.Response().BodyWriter().Write([]byte(p))

	return nil
}

func homeHandler(c *fiber.Ctx) error {
	return c.Render("index", nil)
}
