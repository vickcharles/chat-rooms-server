package user

import (
	"chat-rooms-server/util"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	secretKey = "secret"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
  ctx, cancel := context.WithTimeout(c, s.timeout)

  defer cancel()

  hashedPassword, err := util.HashPassword(req.Password)

  if err != nil {
	return nil, err
  }

  u := &User{
	Username: req.Username,
	Email: req.Email,
	Password: hashedPassword,
  }

  r, err := s.Repository.CreateUser(ctx, u)	

  res := &CreateUserRes{
	ID: r.ID,
	Username: r.Username,
	Email: r.Email,
  }

  return res, err

}

type MyJWTClaims struct {
	ID string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{AccessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}

func (s *service) ParseToken(ctx context.Context, token string) (*User, error) {
    parsedToken, err := jwt.ParseWithClaims(token, &MyJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })

    if err != nil || !parsedToken.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    claims, ok := parsedToken.Claims.(*MyJWTClaims)

    if !ok {
        return nil, fmt.Errorf("invalid token")
    }
	id, err := strconv.ParseInt(claims.ID, 10, 64)
    if err != nil {
        return nil, fmt.Errorf("invalid user ID")
    }

    user := &User{
        ID:  id,
		Username: claims.Username,
    }

    return user, nil
}

