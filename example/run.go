package main
import "github.com/zubairhamed/lab/ogcsensorthings"

func main() {
	server := ogcsensorthings.NewServer()

	server.UseSensingProfile(&DefaultSensingProfileHandler{})
	server.UseTaskingProfile(&DefaultTaskingProfileHandler{})

	server.Start()
}

type DefaultSensingProfileHandler struct {

}

type DefaultTaskingProfileHandler struct {

}