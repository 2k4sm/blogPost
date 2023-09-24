package internals

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Description string
}
