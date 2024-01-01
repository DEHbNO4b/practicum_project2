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

func textToDomain(t TextData) models.TextData {

	td := models.TextData{}

	td.SetUserID(t.UserID)
	td.SetText(t.Text)
	td.SetMeta(t.Meta)

	return td
}

func domainTextToLocal(td models.TextData) TextData {

	t := TextData{}

	t.UserID = td.UserID()
	t.Text = td.Text()
	t.Meta = td.Meta()

	return t
}

func binaryToDomain(b BinaryData) models.BinaryData {

	bd := models.BinaryData{}

	bd.SetUserID(b.UserID)
	bd.SetData(b.Data)
	bd.SetMeta(b.Meta)

	return bd
}

func domainBinaryToLocal(bd models.BinaryData) BinaryData {

	b := BinaryData{}

	b.UserID = bd.UserID()
	b.Data = bd.Data()
	b.Meta = bd.Meta()

	return b
}

func cardToDomain(c Card) (models.Card, error) {

	cd := models.Card{}

	cd.SetUserID(c.UserID)
	err := cd.SetCardID([]rune(c.CardID))
	if err != nil {
		return cd, err
	}
	err = cd.SetPass(c.Pass)
	if err != nil {
		return cd, err
	}
	cd.SetDate(c.Date)
	cd.SetMeta(c.Meta)

	return cd, nil
}

func domainCardToLocal(cd models.Card) Card {
	c := Card{}

	c.UserID = cd.UserID()
	c.CardID = string(cd.CardID())
	c.Date = cd.Date()
	c.Pass = cd.Pass()
	c.Meta = cd.Meta()

	return c
}
