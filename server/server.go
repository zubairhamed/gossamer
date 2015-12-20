package server

import (
	"encoding/json"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	. "github.com/zubairhamed/gossamer"
	"log"
	"net/http"
	"reflect"
)

func NewServer() Server {
	return &GossamerServer{}
}

type GossamerServer struct {
	dataStore Datastore
}

func (s *GossamerServer) handleNotImplemented(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("Not implemented")
}

func (s *GossamerServer) UseStore(ds Datastore) {
	s.dataStore = ds
}

func (s *GossamerServer) Start() {
	goji.Get("/v1.0", s.handleRootResource)
	goji.Get("/v1.0/", s.handleRootResource)

	goji.Get("/v1.0/*", s.HandleGet)
	goji.Post("/v1.0/*", s.HandlePost)
	goji.Put("/v1.0/*", s.HandlePut)
	goji.Delete("/v1.0/*", s.HandleDelete)
	goji.Patch("/v1.0/*", s.HandlePatch)

	log.Println("Start Server")
	goji.Serve()
}

func (s *GossamerServer) Stop() {
	log.Println("Stopped Server")
}

func (s *GossamerServer) handleRootResource(c web.C, w http.ResponseWriter, r *http.Request) {
	data := []ResourceUrlType{
		{"Things", ResolveSelfLinkUrl("", ENTITY_THINGS)},
		{"Locations", ResolveSelfLinkUrl("", ENTITY_LOCATIONS)},
		{"Datastreams", ResolveSelfLinkUrl("", ENTITY_DATASTREAMS)},
		{"Sensors", ResolveSelfLinkUrl("", ENTITY_SENSORS)},
		{"Observations", ResolveSelfLinkUrl("", ENTITY_OBSERVATIONS)},
		{"ObservedProperties", ResolveSelfLinkUrl("", ENTITY_OBSERVEDPROPERTIES)},
		{"FeaturesOfInterest", ResolveSelfLinkUrl("", ENTITY_FEATURESOFINTEREST)},
	}

	v := &ValueList{
		Value: data,
	}

	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
		return
	}
	w.Write(out)
}

func (s *GossamerServer) HandleGet(c web.C, w http.ResponseWriter, r *http.Request) {
	req, err := CreateIncomingRequest(r.URL, HTTP)
	if err != nil {
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
	}

	rp := req.GetResourcePath()
	result, err := s.dataStore.Query(rp, req.GetQueryOptions())
	if err != nil {
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
		return
	}

	var jsonOut interface{}
	if reflect.TypeOf(result).Kind() == reflect.Slice {
		count := reflect.ValueOf(result).Len()
		if req.GetQueryOptions().CountSet() {
			opt := req.GetQueryOptions().GetCountOption()
			if !opt.GetValue() {
				count = 0
			}
		}
		// Check for $count and include

		jsonOut = &ValueList{
			Count: count,
			Value: result,
		}
	} else {
		jsonOut = result
	}

	b, err := json.MarshalIndent(jsonOut, "", "  ")
	if err != nil {
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
		return
	}
}

func (s *GossamerServer) HandlePost(c web.C, w http.ResponseWriter, r *http.Request) {
	var err error
	var req Request

	if r.Header.Get("Content-Type") != "application/json" {
		ThrowNotAcceptable("Format of request is not JSON", w)
	}

	req, err = CreateIncomingRequest(r.URL, HTTP)
	if err != nil {
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
	}

	if err = ValidatePostRequestUrl(req); err != nil {
		ThrowHttpMethodNotAllowed(err.Error(), w)
		return
	}

	rp := req.GetResourcePath()
	ent := rp.Last().GetEntity()
	cont := rp.Containing()

	if !IsSingularEntity(string(ent)) {
		decoder := json.NewDecoder(r.Body)
		e, err := DecodeJsonToEntityStruct(decoder, ent)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+err.Error(), w)
			return
		}

		if cont != nil {
			SetAssociatedEntityId(e, cont.GetEntity(), cont.GetId())
		}

		rp := req.GetResourcePath()

		err = ValidateMandatoryProperties(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+err.Error(), w)
			return
		}

		err = ValidateIntegrityConstraints(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+err.Error(), w)
			return
		}

		err = s.dataStore.Insert(rp, e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+err.Error(), w)
			return
		}

		// TODO
		w.Header().Add("Location", "url-to-entity")
		ThrowHttpCreated("Entity Created", w)
	}
}

func (s *GossamerServer) HandleDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	var err error
	var req Request

	req, err = CreateIncomingRequest(r.URL, HTTP)
	if err != nil {
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST+err.Error(), w)
	}

	if err = ValidateDeleteRequestUrl(req); err != nil {
		ThrowHttpMethodNotAllowed(err.Error(), w)
		return
	}

	rp := req.GetResourcePath()
	last := rp.Last()
	ent := last.GetEntity()
	if !IsSingularEntity(string(ent)) {
		err = s.dataStore.Delete(ent, last.GetId())
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_DELETING_ENTITY+err.Error(), w)
			return
		}
	}
}

func (s *GossamerServer) HandlePut(c web.C, w http.ResponseWriter, r *http.Request) {

}

func (s *GossamerServer) HandlePatch(c web.C, w http.ResponseWriter, r *http.Request) {

}
