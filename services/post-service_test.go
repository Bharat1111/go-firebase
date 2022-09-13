package services

import (
	"firebase/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type MockRepository struct {
// 	mock.Mock
// }

// func (mock *MockRepository) Save(post *constants.Post) (*constants.Post, error) {
// 	args := mock.Called()
// 	result := args.Get(0)
// 	return result.(*constants.Post), args.Error(1)
// }

// func (mock *MockRepository) FindAll() ([]constants.Post, error) {
// 	args := mock.Called()
// 	result := args.Get(0)
// 	return result.([]constants.Post), args.Error(1)
// }

// func TestFindAll(t *testing.T) {
// 	mockRepo := new(MockRepository)

// 	// var mockList []constants.Post
// 	post := constants.Post{ID: 1, Title: "A", Text: "B"}

// 	mockRepo.On("FindAll").Return([]constants.Post{post}, nil)
// 	testService := NewPostService(mockRepo)
// 	result, _ := testService.FindAll()
// 	mockRepo.AssertExpectations(t)
// }

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService()
	err := testService.ValidatePost(nil)

	assert.NotNil(t, err)

	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidatePostTitle(t *testing.T) {
	post := constants.Post{ID: 1, Title: "", Text: ""}
	testService := NewPostService()
	err := testService.ValidatePost(&post)

	assert.NotNil(t, err)

	assert.Equal(t, "The post title is empty", err.Error())
}
