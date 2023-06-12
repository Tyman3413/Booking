package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Ключ для генерации токена
const SecretKey = "secret"

// Функция генерации токена
func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	return claims.SignedString([]byte(SecretKey))
}

/*
	- На вход задаем идентификатор, который будет включен в JWT
	- Был выбран алгоритм подписи jwt.SigningMethodHS256
	- Структура jwt.StandardClaims предоставляет стандартные поля для JWT: издатель и время истечения
	- Метод SignedString подписывает токен, генерируя при этом байтовый массив
*/

// Функция проверки JWT, переданного в виде cookie
func Parsejwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", nil
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}

/*
	- На вход подается строка JWT
	- Метод jwt.ParseWithClaims разбирает и проверяет JWT
	- Если в процессе разбора или проверки токена происходит ошибка или токен не является действительным,
	функция возвращает пустую строку и значение nil для ошибки
	- Если разбор и проверка токена прошли успешно, извлекаются утверждения (claims) из токена
*/
