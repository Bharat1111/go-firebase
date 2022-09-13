package main

import (
	"firebase/controller"
	router "firebase/http"
	"fmt"
	"net/http"
	"os"
)

var (
	postController controller.PostController = controller.NewPostController()
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "D:/Code/projects/todo-app-62a46-firebase-adminsdk-hk1m0-1bb41e1192.json")

	const port string = ":8000"
	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Hello World!")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
	// D:\Code\projects\todo-app-62a46-firebase-adminsdk-hk1m0-1bb41e1192.json
}
