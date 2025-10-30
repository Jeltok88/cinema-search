package main

import (
	"log"

	"github.com/Jeltok88/calendar/StradaGo/cinema-searsh/internal/app/server"
)

func main() {
	r := server.NewGin()
	log.Println("listen :8080")
	_ = r.Run(":8080")
}
