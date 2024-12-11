package responder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Respond(w http.ResponseWriter, r *http.Request, status int, data any) error {
	if data == nil {
		noContent := http.StatusNoContent
		if status != noContent {
			return fmt.Errorf(`responder: is data is nil then status code must be %d %s`, noContent, http.StatusText(noContent))
		}

		w.WriteHeader(status)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("responder: encode json: %w", err)
	}
	return nil
}

func RespondInternalError(w http.ResponseWriter, r *http.Request) error {
	status := http.StatusInternalServerError
	field := NewMetaField(status, http.StatusText(status))
	return Respond(w, r, field.Meta.Status, field.Meta.Message)
}
