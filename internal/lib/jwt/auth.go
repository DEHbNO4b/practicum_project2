package jwt

import "context"

type JwtCredentials struct {
	Token string
}

func (c JwtCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	// Вставляем токен JWT в метаданные запроса для каждого вызова
	return map[string]string{
		"authorization": "Bearer " + c.Token,
	}, nil
}

func (JwtCredentials) RequireTransportSecurity() bool {
	// Здесь возвращаем true, чтобы использовать только зашифрованное соединение (через TLS)
	return true
}
