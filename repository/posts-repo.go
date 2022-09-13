package repository

import (
	"firebase/constants"
)

type PostRepo interface {
	Save(post *constants.Post) (*constants.Post, error)
	FindAll() ([]constants.Post, error)
}
