package main

import (
	"log"
	"os"

	"github.com/flabio/Infeastructure/routers"
)

func main() {
	rd := os.Getenv("MYSQL_HOST")
	log.Println(rd)
	routers.NewRouter()

}
