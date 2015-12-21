package main

import (
	"github.com/zubairhamed/gossamer/server"
)

func main() {
	s := server.NewServer()
	s.UseStore(server.NewMongoStore("localhost", "sensorthings"))
	s.Start()
}
