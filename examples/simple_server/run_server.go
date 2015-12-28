package main

import (
	"github.com/zubairhamed/gossamer/server"
)

func main() {

	s := server.NewServer("localhost", 8000)
	s.UseStore(server.NewMongoStore("localhost", "sensorthings"))
	s.Start()
}
