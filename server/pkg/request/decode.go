package request

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		// response.Json(w, err.Error(), http.StatusBadRequest)
		return payload, err
	}
	return payload, nil
}
