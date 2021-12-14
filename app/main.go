package main

import (
	"cleango/app/infrastuctures"
	"log"
	"time"
)

const timeout time.Duration = 10 * time.Second

func main() {
	db := infrastuctures.NewDB()

	connMySQL, err := db.GetConnMySQL()
	if err != nil {
		log.Println(err.Error())
	}
	r := infrastuctures.Route(connMySQL)
	/*
		s := http.Server{
			Addr:    "8080",
			Handler: r,
		}
	*/
	if err := r.Run(); err != nil {
		log.Println("startup service failed, err:\n", err.Error())
	}
	/*
		go func() {

		}()
		q := make(chan os.Signal, 1)
		signal.Notify(q, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-q

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		if err := http.Shutdown(ctx); err != nil {
			log.Println(err.Error())
		}

		log.Println("\n Gracefully shutdown")
	*/
}
