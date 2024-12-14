package httphandler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jkitajima/efm/lib/responder"
	"github.com/jkitajima/efm/svc/api/pkg/user"
)

func (s *UserServer) handleUserSoftDeleteByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("userID")
		uuid, err := uuid.Parse(id)
		if err != nil {
			// encoding.ErrorRespond(w, r, http.StatusBadRequest, err)
			responder.RespondInternalError(w, r)
			return
		}

		err = s.service.SoftDeleteByID(
			r.Context(),
			user.SoftDeleteByIDRequest{ID: uuid},
		)
		if err != nil {
			responder.RespondInternalError(w, r)
			return
		}

		if err := responder.Respond(w, r, http.StatusNoContent, nil); err != nil {
			responder.RespondInternalError(w, r)
			return
		}
	}
}
