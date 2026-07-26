package main

import (
	"log"
	"os"

	"github.com/evgeni/ftpput/driver/fileput"
	"goftp.io/server/v2"
)

func main() {
	dir := os.Getenv("FTPPUT_DIR")
	if dir == "" {
		dir = "./"
	}
	driver, err := fileput.NewDriver(dir)
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.NewServer(&server.Options{
		Driver: driver,
		Auth: &server.SimpleAuth{
			Name:     "admin",
			Password: "admin",
		},
		Perm:      server.NewSimplePerm("root", "root"),
		RateLimit: 1000000, // 1MB/s limit
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
