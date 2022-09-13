package repository

import (
	"context"
	"firebase/constants"
	"log"

	"cloud.google.com/go/firestore"
)

type postRepo struct {
}

func NewFirestoreRepo() PostRepo {
	return &postRepo{}
}

func (p *postRepo) Save(post *constants.Post) (*constants.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "todo-app-62a46")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection("posts").Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (p *postRepo) FindAll() ([]constants.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "todo-app-62a46")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []constants.Post
	iter := client.Collection("posts").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			log.Println("Failed to iterate the list of posts: %v", err)
			break
		}
		log.Println(doc.Data(), "1", iter)
		post := constants.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
