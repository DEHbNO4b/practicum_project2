package client

import (
	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
)

func pbLpdToDomain(lpd *pb.LogPassData) models.LogPassData {
	l := models.LogPassData{}

	l.SetLogin(lpd.Login)
	l.SetPass(lpd.Password)
	l.SetMeta(lpd.Info)

	return l
}

func pbCardToDomain(cd *pb.CardData) models.Card {
	c := models.Card{}

	c.SetCardID([]rune(cd.CardID))
	c.SetPass(cd.Pass)
	c.SetDate(cd.Date)
	c.SetMeta(cd.Info)

	return c
}

func pbTextToDomain(td *pb.TextData) models.TextData {

	t := models.TextData{}

	t.SetText(td.Text)
	t.SetMeta(td.Info)

	return t
}

func pbBinaryToDomain(bd *pb.BinaryData) models.BinaryData {

	b := models.BinaryData{}

	b.SetData(bd.Data)
	b.SetMeta(bd.Info)

	return b
}

func pbDataToDomain(data *pb.Data) *models.Data {

	d := models.NewData()

	for _, el := range data.Lpd {
		d.AddLpd(pbLpdToDomain(el))
	}

	for _, el := range data.Cd {
		d.AddCd(pbCardToDomain(el))
	}

	for _, el := range data.Td {
		d.AddTd(pbTextToDomain(el))
	}

	for _, el := range data.Bd {
		d.AddBd(pbBinaryToDomain(el))
	}
	return d
}
