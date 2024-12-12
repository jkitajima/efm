package httphandler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jkitajima/efm/identity/auth/pkg/user"
	"github.com/jkitajima/efm/lib/responder"
)

func (s *UserServer) handleValidateEmailService() http.HandlerFunc {
	type request struct {
		UserID           string `json:"user_id"` // TODO: receive id from claims
		VerificationCode string `json:"verification_code"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		req, _ = responder.Decode[request](r)
		// if err != nil {
		// 	encoding.ErrorRespond(w, r, http.StatusBadRequest, err)
		// 	return
		// }

		// id := chi.URLParam(r, "fileID")
		// uuid, err := uuid.Parse(id)
		uuid, err := uuid.Parse(req.UserID)
		if err != nil {
			// encoding.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}

		s.service.ValidateEmail(r.Context(), user.ValidateEmailRequest{
			UserID:           uuid,
			VerificationCode: req.VerificationCode,
		})

		w.Write([]byte("user email validation response"))
	}
}
