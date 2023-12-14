package postgres

import "github.com/DEHbNO4b/practicum_project2/internal/domain/models"

func userToDomain(u User) models.User {

	user := models.User{}

	user.SetID(u.Id)
	user.SetLogin(u.Login)
	user.SetPassHash(u.PassHash)

	return user
}

func domainUserToLocal(u models.User) User {

	user := User{}

	user.Id = u.ID()
	user.Login = u.Login()
	user.PassHash = u.PassHash()

	return user
}
