package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"skillmaster/internal/repository"
	"skillmaster/pkg/entities/models"
	"skillmaster/pkg/entities/responces"
	"strconv"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, username, email, password string) (*models.User, error)
	Login(ctx context.Context, username, password string) (*responces.AuthResponse, error)
	ValidateToken(tokenString string) (*responces.Claims, error)
}

type authService struct {
	repo      repository.AuthRepository
	jwtSecret string
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	return &authService{repo: repo, jwtSecret: jwtSecret}
}

func (s *authService) Register(ctx context.Context, username, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	if err = s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(ctx context.Context, username, password string) (*responces.AuthResponse, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	//todo проверка на протухший и создание нового
	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, expiresAt, err := s.generateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	if err = s.repo.SaveRefreshToken(ctx, strconv.Itoa(int(user.ID)), refreshToken, expiresAt); err != nil {
		return nil, fmt.Errorf("failed to save refresh token: %w", err)
	}

	return &responces.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) ValidateToken(tokenString string) (*responces.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &responces.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*responces.Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	if claims.Exp < time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}

func (s *authService) generateAccessToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	})
	return token.SignedString(s.jwtSecret)
}

func (s *authService) generateRefreshToken() (string, time.Time, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", time.Time{}, err
	}
	token := base64.URLEncoding.EncodeToString(b)
	return token, time.Now().AddDate(0, 0, 7), nil
}
