package transport

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"ranzhouol/go_study/microService/go-kit/pkg/endpoint"
)

func NewHttpHandler(endpoints endpoint.Set) http.Handler {
	m := http.NewServeMux()
	m.Handle("/sum", httptransport.NewServer(endpoints.SumEndpoint, decodeHTTPSumRequest, encodeHTTPGenericResponse))
	m.Handle("/concat", httptransport.NewServer(endpoints.ConcatEndpoint, decodeHTTPConcatRequest, encodeHTTPGenericResponse))
	return m
}

// http解码-反序列化
func decodeHTTPSumRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.SumRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPConcatRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.ConcatRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// http编码-序列化
func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Context-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
