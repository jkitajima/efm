package httphandler

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jkitajima/efm/lib/responder"
	"github.com/jkitajima/efm/svc/api/pkg/user"
)

func (s *UserServer) handleUserCreate() http.HandlerFunc {
	type request struct {
		FirstName string  `json:"first_name"`
		LastName  *string `json:"last_name"`
		Role      *string `json:"role"`
	}

	type response struct {
		Entity    string    `json:"entity"`
		ID        uuid.UUID `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  *string   `json:"last_name"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req, err := responder.Decode[request](r)
		if err != nil {
			responder.RespondInternalError(w, r)
			return
		}

		createResponse, err := s.service.Create(r.Context(), user.CreateRequest{
			User: &user.User{
				FirstName: req.FirstName,
				LastName:  req.LastName,
				Role:      user.Role(*req.Role),
			},
		})
		if err != nil {
			responder.RespondInternalError(w, r)
			return
		}

		resp := response{
			Entity:    s.entity,
			ID:        createResponse.User.ID,
			FirstName: createResponse.User.FirstName,
			LastName:  createResponse.User.LastName,
			Role:      string(createResponse.User.Role),
			CreatedAt: createResponse.User.CreatedAt,
			UpdatedAt: createResponse.User.UpdatedAt,
		}

		if err := responder.Respond(w, r, http.StatusCreated, &responder.DataField{Data: resp}); err != nil {
			responder.RespondInternalError(w, r)
			return
		}
	}
}
