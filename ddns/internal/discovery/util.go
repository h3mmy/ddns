package discovery

import (
	"encoding/json"
	"net/http"

)


func parseResponse(r *http.Response, target interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
