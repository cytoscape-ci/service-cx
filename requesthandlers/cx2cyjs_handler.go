package handlers

import (
	"net/http"
	"github.com/cytoscape-ci/cxtool/converter"
)

func Cx2CyjsHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case POST:
		cx2cyjs(w, r)
	default:
		unsupported(w, r)
	}
}

func cx2cyjs(w http.ResponseWriter, r *http.Request) {
	conv := converter.Cx2Cyjs{}
	conv.Convert(r.Body, w)
}

