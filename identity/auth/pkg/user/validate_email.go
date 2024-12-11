package user

import (
	"context"
	// "fmt"
	// "strconv"
	// "time"

	"github.com/google/uuid"
)

type ValidateEmailRequest struct {
	UserID           uuid.UUID
	VerificationCode string
}

type ValidateEmailResponse struct {
	Validated bool
}

func (s *Service) ValidateEmail(ctx context.Context, req ValidateEmailRequest) (ValidateEmailResponse, error) {
	user, err := s.Repo.FindByID(ctx, req.UserID)
	if err != nil {
		return ValidateEmailResponse{}, err
	}

	// code := *user.VerificationCode

	// Only should do this validation if not expired.
	// If expired, client should ask for a new verification code.
	// now := time.Now()
	// codeExp := *user.VerificationCodeExpiration
	// expiration, _ := strconv.ParseInt(strconv.Itoa(codeExp), 10, 64)
	// expirationTime := time.Unix(expiration, 0)

	// comparison := now.Compare(expirationTime)
	// fmt.Println(comparison)

	// if code != req.VerificationCode || comparison == 1 {
	// 	return ValidateEmailResponse{false}, nil
	// }

	// If valid, update the user model
	s.Repo.UpdateByID(ctx, req.UserID, user)

	return ValidateEmailResponse{true}, nil
}
