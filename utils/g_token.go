package utils

type Token struct {
	Token string
}

func GenerateToken(username string, password string) Token {
	return Token{Token: username + password}
}
