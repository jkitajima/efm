package httphandler

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jkitajima/efm/lib/responder"
	"github.com/jkitajima/efm/svc/api/pkg/user"
)

func (s *UserServer) handleUserCreate() http.HandlerFunc {
	type request struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name"`
		Role      string `json:"role" validate:"oneof=default admin"`
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

	var contract = map[string]struct {
		field      string
		validation string
	}{
		"FirstName": {
			field:      "first_name",
			validation: "Field is required and cannot be an empty string.",
		},
		"LastName": {
			field:      "last_name",
			validation: "Field value cannot be an empty string.",
		},
		"Role": {
			field:      "role",
			validation: "Field value must be either 'default' or 'admin'.",
		},
	}

	validateInput := func(req request) ([]responder.ErrorObject, error) {
		errors := make([]responder.ErrorObject, 0, len(contract))

		err := s.inputValidator.Struct(req)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				structField := err.StructField()
				t := contract[structField].field
				d := contract[structField].validation
				errors = append(errors, responder.ErrorObject{
					Title:  t,
					Detail: &d,
				})
			}
		}
		return errors, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req, err := responder.Decode[request](r)
		if err != nil {
			responder.RespondMeta(w, r, http.StatusBadRequest)
			return
		}

		errors, err := validateInput(req)
		if err != nil {
			responder.RespondClientErrors(w, r, errors...)
			return
		}

		createResponse, err := s.service.Create(r.Context(), user.CreateRequest{
			User: &user.User{
				FirstName: req.FirstName,
				LastName:  &req.LastName,
				Role:      user.Role(req.Role),
			},
		})
		if err != nil {
			// TODO: error from business domain could be either client or server side
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
