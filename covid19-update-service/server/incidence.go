package server

import (
	"covid19-update-service/model"
	"fmt"
	"net/http"

	"github.com/pmoule/go2hal/hal"

	"github.com/gorilla/mux"
)

type Incidence struct {
	Incidence float64 `json:"incidence"`
}

func (i Incidence) ToHAL(path string) hal.Resource {
	root := hal.NewResourceObject()
	root.AddData(i)

	selfRel := hal.NewSelfLinkRelation()
	selfLink := &hal.LinkObject{Href: path}
	selfRel.SetLink(selfLink)
	root.AddLink(selfRel)

	return root
}

func (ws *Covid19UpdateWebServer) registerIncidenceRoutes(r *mux.Router) {
	incidenceRouter := r.Path(incidenceRoute).Subrouter().StrictSlash(strictSlash)
	incidenceRouter.HandleFunc("", ws.checkAcceptType(ws.getIncidence)).Methods(http.MethodGet)
	incidenceRouter.HandleFunc("", nil).Methods(http.MethodOptions)
	incidenceRouter.Use(newCorsHandler(incidenceRouter))
	incidenceRouter.MethodNotAllowedHandler = ws.createNotAllowedHandler(incidenceRouter)
}

func (ws *Covid19UpdateWebServer) getIncidence(w http.ResponseWriter, r *http.Request) {
	t, ok := findTopic(w, r)
	if !ok {
		return
	}
	c, err := model.GetCovid19Region(t.Covid19RegionID)
	if err != nil {
		writeHTTPResponse(model.NewError(fmt.Sprintf("Could not load incidence: %v.", err)), http.StatusInternalServerError, w, r)
		return
	}
	if c == nil {
		writeHTTPResponse(model.NewError("No incidence available"), http.StatusInternalServerError, w, r)
		return
	}
	rsp := Incidence{c.Incidence}
	writeHTTPResponse(rsp, http.StatusOK, w, r)
}
