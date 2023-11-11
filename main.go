package main

import (
	_ "github.com/go-sql-driver/mysql"
	"hackathon/controller/post_controller"
	"hackathon/controller/user_controller"
	"hackathon/dao"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dao.Init()

	http.HandleFunc("/users", user_controller.GetUserListHandler)
	http.HandleFunc("/signup", user_controller.PostNewUserHandler)
	http.HandleFunc("/user/update", user_controller.PostUpdateUserHandler)
	http.HandleFunc("/user/delete", user_controller.PostDeleteUserHandler)
	http.HandleFunc("/user", user_controller.GetUserDetailHandler)

	http.HandleFunc("/posts/categories", post_controller.GetPostsByCategoryHandler)
	http.HandleFunc("/posts/curriculums/all", post_controller.GetAllCurriculumsHandler)
	http.HandleFunc("/posts/curriculums", post_controller.GetPostsByCurriculumHandler)
	http.HandleFunc("/posts/tags/all", post_controller.GetAllTagsHandler)
	http.HandleFunc("/posts/user", post_controller.GetPostsByUserHandler)
	http.HandleFunc("/posts/date", post_controller.GetPostsByDateHandler)
	http.HandleFunc("/posts/tags", post_controller.GetPostsByTagHandler)
	http.HandleFunc("/post/like", post_controller.PostLikePostHandler)
	http.HandleFunc("/post/unlike", post_controller.PostUnlikePostHandler)
	http.HandleFunc("/posts/likes", post_controller.GetLikedPostsHandler)
	http.HandleFunc("/post/new", post_controller.PostNewPostHandler)
	http.HandleFunc("/post/update", post_controller.PostUpdatePostHandler)
	http.HandleFunc("/post/delete", post_controller.PostDeletePostHandler)
	http.HandleFunc("/post/comment", post_controller.PostCommentHandler)
	http.HandleFunc("/post/comments", post_controller.GetCommentsHandler)
	http.HandleFunc("/post/comment/delete", post_controller.PostDeleteCommentHandler)
	http.HandleFunc("/post/comment/update", post_controller.PostUpdateCommentHandler)
	http.HandleFunc("/post", post_controller.GetPostHandler)

	closeDBWithSysCall()
	port := os.Getenv("PORT")
	log.Println("Listening...")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := dao.CloseDb(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
