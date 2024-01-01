package tests

import (
	"crypto/rand"
	"fmt"

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
func getRandomBinaryData(n int) *pb.BinaryData {
	ans := &pb.BinaryData{} // Инициализируем структуру

	// Инициализируем срез ans.Data с нужной длиной
	ans.Data = make([]byte, n)

	// Генерируем n случайных байт и копируем их в ans.Data
	if _, err := rand.Read(ans.Data); err != nil {
		fmt.Println(err)
		return nil // Если есть ошибка, возвращаем nil
	}

	// Заполняем Info фальшивым предложением из 10 слов
	ans.Info = gofakeit.Sentence(10)

	return ans
}
