package httphandler

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jkitajima/efm/identity/auth/pkg/user"
	"github.com/jkitajima/efm/lib/responder"
)

func (s *UserServer) handleUserRegister() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		Entity    string     `json:"entity"`
		ID        uuid.UUID  `json:"id"`
		Email     string     `json:"email"`
		Password  string     `json:"password"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Check if "email" field is actually an email formatted string

		req, err := responder.Decode[request](r)
		if err != nil {
			responder.RespondInternalError(w, r)
			return
		}

		registerResponse, err := s.service.Register(r.Context(), user.RegisterRequest{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			responder.RespondInternalError(w, r)
			return
		}

		resp := response{
			Entity:    s.entity,
			ID:        registerResponse.User.ID,
			Email:     registerResponse.User.Email,
			Password:  registerResponse.User.Password,
			CreatedAt: registerResponse.User.CreatedAt,
			UpdatedAt: registerResponse.User.UpdatedAt,
			DeletedAt: registerResponse.User.DeletedAt,
		}

		if err := responder.Respond(w, r, http.StatusCreated, &responder.DataField{Data: resp}); err != nil {
			responder.RespondInternalError(w, r)
			return
		}
	}
}
