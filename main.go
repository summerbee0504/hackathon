package main

import (
	_ "github.com/go-sql-driver/mysql"
	"hackathon/controller"
	"hackathon/dao"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dao.Init()

	http.HandleFunc("/users", controller.GetHandler)
	http.HandleFunc("/user", controller.PostHandler)

	closeDBWithSysCall()

	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
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
