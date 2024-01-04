package tui

import "github.com/DEHbNO4b/practicum_project2/internal/domain/models"

func lpToDomain(lp LogPass) models.LogPassData {

	data := models.LogPassData{}
	data.SetLogin(lp.Login)
	data.SetPass(lp.Pass)
	data.SetMeta(lp.Info)

	return data
}
