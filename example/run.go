package main
import (
	"github.com/zubairhamed/gossamer"
)

func main() {
	server := gossamer.NewServer()

	server.UseSensingProfile(&DefaultSensingProfileHandler{})
	server.UseTaskingProfile(&DefaultTaskingProfileHandler{})

	server.Start()
}

type DefaultSensingProfileHandler struct {

}

type DefaultTaskingProfileHandler struct {

}