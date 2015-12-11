package main

import (
	"github.com/zubairhamed/gossamer"
)

func main() {
	server := gossamer.NewServer()
	
	server.UseStore(gossamer.NewMongoStore("localhost"))

	server.Start()
}

type DefaultSensingProfileHandler struct {
}

type DefaultTaskingProfileHandler struct {
}
