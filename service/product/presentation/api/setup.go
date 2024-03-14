package api

import (
	"net/http"
	apihandler "product/app/handler/api"
)

func Set(mux *http.ServeMux) {
	apihandler.SetAPIHandler(mux)
}
