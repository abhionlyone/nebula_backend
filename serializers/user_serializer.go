package serializers

import "nebula_backend/models"

type UserSerializer struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SerializeUser(user models.User) UserSerializer {
	return UserSerializer{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func SerializeUsers(users []models.User) []UserSerializer {
	serializedUsers := make([]UserSerializer, len(users))
	for i, user := range users {
		serializedUsers[i] = SerializeUser(user)
	}
	return serializedUsers
}
