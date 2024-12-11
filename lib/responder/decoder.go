package responder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Decode[T any](r *http.Request) (v T, err error) {
	if err = json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("responder: decode json: %w", err)
	}
	return v, nil
}
