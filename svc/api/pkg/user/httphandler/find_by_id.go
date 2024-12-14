package httphandler

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jkitajima/efm/lib/responder"
	"github.com/jkitajima/efm/svc/api/pkg/user"
)

func (s *UserServer) handleUserFindByID() http.HandlerFunc {
	type response struct {
		Entity    string    `json:"entity"`
		ID        uuid.UUID `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("userID")
		uuid, err := uuid.Parse(id)
		if err != nil {
			// encoding.ErrorRespond(w, r, http.StatusBadRequest, err)
			responder.RespondInternalError(w, r)
			return
		}

		findResponse, err := s.service.FindByID(r.Context(), user.FindByIDRequest{ID: uuid})
		if err != nil {
			responder.RespondInternalError(w, r)
			return
		}

		resp := response{
			Entity:    s.entity,
			ID:        findResponse.User.ID,
			FirstName: findResponse.User.FirstName,
			LastName:  *findResponse.User.LastName,
			Role:      string(findResponse.User.Role),
			CreatedAt: findResponse.User.CreatedAt,
			UpdatedAt: findResponse.User.UpdatedAt,
		}

		if err := responder.Respond(w, r, http.StatusOK, &responder.DataField{Data: resp}); err != nil {
			responder.RespondInternalError(w, r)
			return
		}
	}
}
