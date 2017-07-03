package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"context"
	"net/http"
	"strconv"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(ctx context.Context, e Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter().StrictSlash(false)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
	}
	r.Methods("GET").Path("/deals").Handler(httptransport.NewServer(
		e.GetDealEndpoint,
		DecodeRequest,
		EncodeResponse,
		options...,
	))
	return r
}

func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idStr := r.FormValue("id")
	var id int
	var err error
	if len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
	} else {
		id = -1
	}
	return getDealRequest{
		ID: id,
	}, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
