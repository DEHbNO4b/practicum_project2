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

func lpToDomain(lp LogPassData) models.LogPassData {

	dlp := models.LogPassData{}

	dlp.SetLogin(lp.Login)
	dlp.SetUserID(lp.UserID)
	dlp.SetPass(lp.Pass)
	dlp.SetMeta(lp.Meta)

	return dlp
}

func domainLpToLocal(lp models.LogPassData) LogPassData {

	local := LogPassData{}

	local.UserID = lp.UserID()
	local.Login = lp.Login()
	local.Pass = lp.Pass()
	local.Meta = lp.Meta()

	return local
}
