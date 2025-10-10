package networks

import (
	"net/http"
)

// HTTPHandler handles HTTP requests
type HTTPHandler struct {
}

func (f *HTTPHandler) Register(Path string, Method interface{}) {

}

func (f *HTTPHandler) HandlingRequest(w http.ResponseWriter, r *http.Request) {

}
