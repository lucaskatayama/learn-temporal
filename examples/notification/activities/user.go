package activities

import "context"

type User struct {
	ID       string
	MobileID string
	Name     string
	Email    string
}

func GetUserInformation(ctx context.Context, userID string) (User, error) {
	return User{
		ID:       "123",
		MobileID: "456",
		Name:     "John Doe",
		Email:    "john.doe@email.com",
	}, nil
}
