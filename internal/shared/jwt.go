package shared

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWT struct{ Secret []byte }

func NewJWT(cfg *Config) *JWT {
	return &JWT{Secret: []byte(cfg.JWT.Secret)}
}

type Claims struct {
	UserID string `json:"uid"`
	jwt.RegisteredClaims
}

func (r *JWT) IssueToken(uid string) (token, jti string, err error) {
	jti = uuid.NewString()
	claims := Claims{
		UserID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString(r.Secret)
	return
}

func (r *JWT) ParseToken(tok string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(tok, &Claims{}, func(*jwt.Token) (any, error) {
		return r.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return t.Claims.(*Claims), nil
}
