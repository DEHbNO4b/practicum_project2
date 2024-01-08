package tui

import "github.com/DEHbNO4b/practicum_project2/internal/domain/models"

func lpToDomain(lp LogPass) models.LogPassData {

	data := models.LogPassData{}
	data.SetLogin(lp.Login)
	data.SetPass(lp.Pass)
	data.SetMeta(lp.Info)

	return data
}

func cdToDomain(cd Card) models.Card {
	card := models.Card{}

	card.SetCardID([]rune(cd.CardID))
	card.SetPass(cd.Pass)
	card.SetDate(cd.Date)
	card.SetMeta(cd.Info)

	return card
}

func tdToDomain(td TextData) models.TextData {
	text := models.TextData{}

	text.SetText(td.Text)
	text.SetMeta(td.Meta)

	return text
}

func bdToDomain(bd BinaryData) models.BinaryData {
	data := models.BinaryData{}

	data.SetData([]byte(bd.Data))
	data.SetMeta(bd.Info)

	return data
}
