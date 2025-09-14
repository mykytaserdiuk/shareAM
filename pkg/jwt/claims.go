package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

const (
	AuthHeader string = "Authorization"
)

var (
	ErrBadJwt    = errors.New("JWT token invalid")
	ErrJwtEncode = errors.New("error while JWT encoding")
)

func ExtractClaims(jwtToken string) (*jwt.StandardClaims, error) {
	parts := strings.Split(jwtToken, ".")

	if len(parts) < 2 {
		return nil, ErrBadJwt
	}

	decoded, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, ErrJwtEncode
	}

	claims := &jwt.StandardClaims{}
	err = json.Unmarshal(decoded, claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func ExtractClaimsFromRequest(r *http.Request) (*jwt.StandardClaims, error) {
	return ExtractClaims(r.Header.Get(AuthHeader))
}
