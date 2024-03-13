package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/todoApp/internal/models"
	"github.com/todoApp/internal/repository"
	"github.com/todoApp/internal/service"
	"github.com/todoApp/internal/service/dtos"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler() *UserHandler {
	repo := repository.NewRepository()
	deps := service.Deps{Repositories: repo}
	return &UserHandler{
		userService: *service.NewUserService(deps.Repositories.Users),
	}
}

// handlers
func (u *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var userInput dtos.InputUserDto

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	userFromDb := u.userService.GetUserByUsername(userInput.Username)

	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.PasswordHash), []byte(userInput.Password))
	if err != nil {
		http.Error(w, "Invalid Username or Password", http.StatusBadRequest)
	}

	token, err := generateToken(*userFromDb)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func (u *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var userInput dtos.InputUserDto

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	u.userService.CreateUser(userInput)
	w.WriteHeader(http.StatusCreated)
}

// jwt token logic
var signingKey = []byte("super secret key")

func generateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Username,
			"user_id":  user.Id,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	signingKey := []byte("super secret key")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
