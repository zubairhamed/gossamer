package server

import (
	"encoding/json"
	"flag"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	. "github.com/zubairhamed/gossamer"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"html/template"
	"bytes"
)

func NewServer(host string, port int) Server {
	GLOB_ENV_HOST = host + ":" + strconv.Itoa(port)

	tpl := CreateTemplateCache()

	return &GossamerServer{
		tpl: tpl,
		host: host,
		port: port,
	}
}

func CreateTemplateCache() *template.Template {
	tplBuf := bytes.NewBuffer([]byte{})


	var tpls = []string {
		"head",
		"menu",
		"entity_list",
		"entity_content",

		"things",
		"datastreams",
		"featuresofinterest",
		"index",
		"locations",
		"observations",
		"observedproperties",
		"sensors",
	}
	var tpl []byte

	for _, v := range tpls {
		tpl, _ = AssetContent("tpl/" + v + ".html")
		tplBuf.Write(tpl)
	}

	t, _ := template.New("tpls").Delims("#{", "}#").Parse(tplBuf.String())

	return t
}

type GossamerServer struct {
	tpl *template.Template
	dataStore Datastore
	host      string
	port      int
}

func (s *GossamerServer) handleNotImplemented(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("Not implemented")
}

func (s *GossamerServer) UseStore(ds Datastore) {
	s.dataStore = ds
}

func (s *GossamerServer) Start() {
	goji.Get("/", s.HandleWebUiIndex)
	goji.Get("/things.html", s.HandleWebUiThings)
	goji.Get("/sensors.html", s.HandleWebUiSensors)
	goji.Get("/observations.html", s.HandleWebUiObservations)
	goji.Get("/observedproperties.html", s.HandleWebUiObservedProperties)
	goji.Get("/locations.html", s.HandleWebUiLocations)
	goji.Get("/datastreams.html", s.HandleWebUiDatastreams)
	goji.Get("/featuresofinterest.html", s.HandleWebUiFeaturesOfInterest)
	goji.Get("/historiclocations.html", s.HandleWebUiHistoricLocations)
	goji.Get("/css/*", s.HandleWebUiResources)
	goji.Get("/img/*", s.HandleWebUiResources)
	goji.Get("/js/*", s.HandleWebUiResources)

	goji.Get("/v1.0", s.handleRootResource)
	goji.Get("/v1.0/", s.handleRootResource)

	goji.Get("/v1.0/*", s.HandleGet)
	goji.Post("/v1.0/*", s.HandlePost)
	goji.Put("/v1.0/*", s.HandlePut)
	goji.Delete("/v1.0/*", s.HandleDelete)
	goji.Patch("/v1.0/*", s.HandlePatch)

	flag.Set("bind", ":"+strconv.Itoa(s.port))

	log.Println("Start Server on port ", s.port)
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
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+": "+err.Error(), w)
		return
	}
	w.Write(out)
}

func (s *GossamerServer) HandleGet(c web.C, w http.ResponseWriter, r *http.Request) {
	req, err := CreateIncomingRequest(r.URL, HTTP)
	if err != nil {
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST+": "+err.Error(), w)
	}

	rp := req.GetResourcePath()
	result, err := s.dataStore.Query(rp, req.GetQueryOptions())
	if err != nil {
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+": "+err.Error(), w)
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
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+": "+err.Error(), w)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		ThrowHttpInternalServerError(MSG_ERR_HANDLING_REQUEST+": "+err.Error(), w)
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
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST+": "+err.Error(), w)
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
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
			return
		}

		if cont != nil {
			SetAssociatedEntityId(e, cont.GetEntity(), cont.GetId())
		}

		rp := req.GetResourcePath()

		err = ValidateMandatoryProperties(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
			return
		}

		err = ValidateIntegrityConstraints(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
			return
		}

		err = s.dataStore.Insert(rp, e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
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
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST+": "+err.Error(), w)
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
			ThrowHttpBadRequest(MSG_ERR_DELETING_ENTITY+": "+err.Error(), w)
			return
		}
	}
}

func (s *GossamerServer) HandlePut(c web.C, w http.ResponseWriter, r *http.Request) {
	var err error
	var req Request

	if r.Header.Get("Content-Type") != "application/json" {
		ThrowNotAcceptable("Format of request is not JSON", w)
	}

	req, err = CreateIncomingRequest(r.URL, HTTP)
	if err != nil {
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST+": "+err.Error(), w)
	}

	if err = ValidatePutRequestUrl(req); err != nil {
		ThrowHttpMethodNotAllowed(err.Error(), w)
		return
	}

	rp := req.GetResourcePath()
	lastEnt := rp.Last()
	ent := lastEnt.GetEntity()

	if lastEnt.GetId() == "" {
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST+": Missing @iot.id", w)
	}

	cont := rp.Containing()

	if !IsSingularEntity(string(ent)) {
		decoder := json.NewDecoder(r.Body)
		e, err := DecodeJsonToEntityStruct(decoder, ent)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
			return
		}

		if cont != nil {
			SetAssociatedEntityId(e, cont.GetEntity(), cont.GetId())
		}

		err = ValidateMandatoryProperties(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
			return
		}

		err = ValidateIntegrityConstraints(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
			return
		}

		err = s.dataStore.Patch(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
			return
		}

		// TODO
		w.Header().Add("Location", "url-to-entity")
		ThrowHttpOk("Entity Updated", w)
	}
}

func (s *GossamerServer) HandlePatch(c web.C, w http.ResponseWriter, r *http.Request) {
	var err error
	var req Request

	if r.Header.Get("Content-Type") != "application/json" {
		ThrowNotAcceptable("Format of request is not JSON", w)
	}

	req, err = CreateIncomingRequest(r.URL, HTTP)
	if err != nil {
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST+": "+err.Error(), w)
	}

	if err = ValidatePutRequestUrl(req); err != nil {
		ThrowHttpMethodNotAllowed(err.Error(), w)
		return
	}

	rp := req.GetResourcePath()
	lastEnt := rp.Last()
	ent := lastEnt.GetEntity()

	if lastEnt.GetId() == "" {
		ThrowHttpBadRequest(MSG_ERR_HANDLING_REQUEST+": Missing @iot.id", w)
	}

	cont := rp.Containing()

	if !IsSingularEntity(string(ent)) {
		decoder := json.NewDecoder(r.Body)
		e, err := DecodeJsonToEntityStruct(decoder, ent)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
			return
		}

		if cont != nil {
			SetAssociatedEntityId(e, cont.GetEntity(), cont.GetId())
		}

		err = ValidateMandatoryProperties(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
			return
		}

		err = ValidateIntegrityConstraints(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
			return
		}

		err = s.dataStore.Update(e)
		if err != nil {
			ThrowHttpBadRequest(MSG_ERR_INSERTING_ENTITY+": "+err.Error(), w)
			return
		}

		// TODO
		w.Header().Add("Location", "url-to-entity")
		ThrowHttpOk("Entity Updated", w)
	}
}

func (s *GossamerServer) HandleWebUiIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	s.tpl.ExecuteTemplate(w, "page_index", nil)
}

func (s *GossamerServer) HandleWebUiThings(c web.C, w http.ResponseWriter, r *http.Request) {
	err := s.tpl.ExecuteTemplate(w, "page_things", map[string]string {
		"Title": "Things",
	})
	if err != nil {
		log.Println(err)
	}
}

func (s *GossamerServer) HandleWebUiSensors(c web.C, w http.ResponseWriter, r *http.Request) {
	s.tpl.ExecuteTemplate(w, "page_sensors", map[string]string {
		"Title": "Sensors",
	})
}

func (s *GossamerServer) HandleWebUiObservations(c web.C, w http.ResponseWriter, r *http.Request) {
	s.tpl.ExecuteTemplate(w, "page_observations", map[string]string {
		"Title": "Observations",
	})
}

func (s *GossamerServer) HandleWebUiObservedProperties(c web.C, w http.ResponseWriter, r *http.Request) {
	s.tpl.ExecuteTemplate(w, "page_obsprops", map[string]string {
		"Title": "Observed Properties",
	})
}

func (s *GossamerServer) HandleWebUiLocations(c web.C, w http.ResponseWriter, r *http.Request) {
	s.tpl.ExecuteTemplate(w, "page_locs", map[string]string {
		"Title": "Locations",
	})
}

func (s *GossamerServer) HandleWebUiDatastreams(c web.C, w http.ResponseWriter, r *http.Request) {
	s.tpl.ExecuteTemplate(w, "page_ds", map[string]string {
		"Title": "Datastreams",
	})
}

func (s *GossamerServer) HandleWebUiFeaturesOfInterest(c web.C, w http.ResponseWriter, r *http.Request) {
	s.tpl.ExecuteTemplate(w, "page_foi", map[string]string {
		"Title": "Features of Interest",
	})
}

func (s *GossamerServer) HandleWebUiHistoricLocations(c web.C, w http.ResponseWriter, r *http.Request) {
	s.tpl.ExecuteTemplate(w, "page_histolocs", map[string]string {
		"Title": "Historic Locations",
	})
}

func (s *GossamerServer) HandleWebUiResources(c web.C, w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case strings.HasPrefix(path, "/css"):
		w.Header().Set("Content-Type", "text/css")

	case strings.HasPrefix(path, "/js"):
		w.Header().Set("Content-Type", "text/javascript")
	}

	data, _ := AssetContent(path)

	w.Write(data)
}

func AssetContent(path string) ([]byte, error) {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	data, err := Asset("resources/" + path)

	if err != nil {
		log.Println(err)
	}

	return data, err
}

type PageModel struct {
	Title 	string
}
