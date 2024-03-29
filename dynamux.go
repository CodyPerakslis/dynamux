package dynamux

import (
  "net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)
type DynamicRequest func(*http.Request) interface{}
type DynamicResponse func(http.ResponseWriter)
type Method func(data interface{}) DynamicResponse

func New(header map[string][]string, req DynamicRequest, method Method) Handler {
  return func(w http.ResponseWriter, r *http.Request) {
    for k, v := range header {for _, elem := range v {w.Header().Set(k, elem)}}
    method(req(r))(w)
  }
}
