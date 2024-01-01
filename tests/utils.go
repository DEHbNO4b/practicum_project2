package tests

import (
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
	"github.com/brianvoe/gofakeit/v6"
)

func getRandomLogPassData() *pb.LogPassData {
	return &pb.LogPassData{
		Login:    gofakeit.Name(),
		Password: gofakeit.Password(true, true, true, true, false, 10),
		Info:     gofakeit.Sentence(10),
	}
}
func getRandomTextData() *pb.TextData {
	return &pb.TextData{
		Text: gofakeit.Paragraph(1, 3, 40, "."),
		Info: gofakeit.Sentence(10),
	}
}
