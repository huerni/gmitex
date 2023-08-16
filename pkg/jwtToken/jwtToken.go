package jwtToken

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"net/http"
)

const signingKey string = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC62BKDitin7jnx\n5hLrdQUcNbEsdbsBbI2eb3E8aSFr1GxRfNQVX/Dni0m0kLe1p0PMARDb9MNcmCBc\nU+1mQHaRsntXemxdV4wEQ3/a3nm6Z2DN+wQ5dfW7GtPk0ElBxkymDtQddZrFZBJP\nbWniDCeizm7YbHffSvMXfbt5ndjxCUO/iWitgZIYq+6v0N8FPO0+m+7dp4V+CTVA\nbFrvXd0hGz4c4zz3vEIQU3s0o3AJDmNdXpoX4MEHTEMx33/wha/ilcLrYSBw0xwI\nozuMN+JPj4eCKxfCBk3B7E6aIFLVMsgSQVGgDSq6pwt4IFwb8FFHSQzYNTwxWK8k\nO91+RfZVAgMBAAECggEAfqg+VKFooN3itdIq/SYEYs0a33KnZB28Gqyc7ECwATKs\ngsjF0/+HhM5tFlQL1L4gPUhzr0dKr5gIR8403d3RAo8lAXXhw5y3M6S2JR4vEmdF\nvhvtDy6hd0aGYVO4dTgBeYgPzjCMzEY8C0+2OR/YNosNpPRShjF+fGwlDBoul6Vv\nN3ygL6bOowm0KMG1wMgMnSkYkflbr8AUtawHxaPfRHKJrSQuGP3aKjONqZ44W7N/\nUGGTSD16RgV+hYRpM7WyHppeD9T82wb+9bYbVLSq6eUQmB1VA5SjaEya7iI/NO0O\nlw3Ayga+X/LIDJ6ttLV5OoXLI56z1DqgLf9vX4/8gQKBgQD4WiTghS1toLaIHVdA\nq8y/xaRilgiU0K+GOtJV6zsY+GviyoQ1n//6awDLzXq5148sd1DJVQQhesskCrJ7\n541d9z8Gg9XnDr2iUFZ2JZT7HxEAUWgljoZPupNcUVxZTKXgkHa9v0rEIEcOsnpx\nd3fSzoSeHjEgG3hkY3tgaKm1xQKBgQDAmQlB6ZINSwtn9smb8BdfArZ9+52GSRYH\n0l6lqXjui5SMo15CytVxffHGKFx1RlKzK8v85ACxXnnJ7tZbVaaWVCRzEjVTc72f\nQ3qBx8RoZu0JUk+B4wwmqOD0b2PK7pPqB8Df63UCctyAS9yNvqXvxZvSNkEOUcp2\nMqFo7u/XUQKBgFa91Wd93HIP0fEUnmb+GlNYyqOMV47ynHu7i79qm4eLLNNXfHnm\nWleyi+Ki8Bx8x4r2WYcYZIr1AoKiIdjY9S0+sAqsfUdohJ9ug+RcF/7lyOBdjoyf\njRXHyrRRznl6Je2bR33alFiQFYFyoQWEfptoejVnwiy+q3wUqwDvTWcpAoGBAJjx\n7tQU9BLyYWByLrBS/XxJ3zo0smeNap1Thi3wY0SsO49jvNs10EKMTY+bRbEr40i9\nowR028f+yqB9tmRZpC0FLNzkvMxEwXTUVVjylxqBggNBBjqTX3bj7aCvRIRG6deT\nyKsJhKYpKMoJdGBr4cKDHrbUttz0Pt+WXW/DL1vBAoGAbJQaI1IFQx0eG2Q6zPhc\naq2KANRktb0hyrgLkE6vu6poMSIsKKnWHikv5lrhtbkLmyWxvhRUqxPEmGcNDqih\nCR482hFFMLw2GJ2LlwNLrQuE1TyzVCGxSkz35RDE4fiSqFtS6tkUbZ+PBvbTjJJc\nseKMw6qqyPmzspwOvdP/d3w=\n"

func GenerateToken(claims jwt.MapClaims) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(signingKey))
	return token, err
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("conversion claims error")
	}
	return claims, nil
}

func ParseTokenRequest(r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (any, error) {
			return []byte(signingKey), nil
		}, request.WithParser(newParser()))

	if err != nil {
		return nil, err
	}

	return token, nil
}

func doParseToken(r *http.Request, secret string) (*jwt.Token, error) {
	return request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		}, request.WithParser(newParser()))
}

func newParser() *jwt.Parser {
	return jwt.NewParser(jwt.WithJSONNumber())
}
