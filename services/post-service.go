package services

import (
	"math/rand"
	"errors"
	"firebase/constants"
	"firebase/repository"
)

type PostService interface {
	ValidatePost(post *constants.Post) error
	Create(post *constants.Post) (*constants.Post, error)
	FindAll() ([]constants.Post, error)
}

type postService struct {}

var (
	repo repository.PostRepo = repository.NewFirestoreRepo()
)

func NewPostService() PostService {
	return &postService{}
}

func (p *postService) ValidatePost(post *constants.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (p *postService) Create(post *constants.Post) (*constants.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (p *postService) FindAll() ([]constants.Post, error) {
	return repo.FindAll()
}