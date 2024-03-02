package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	token, err := generateToken(userFromDb.Id)
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

func generateToken(userId int) (string, error) {
	//temporarily
	signingKey := []byte("super secret key")

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(43200, 0)),
		Issuer:    "test",
		//Audience: ,
		Subject: strconv.Itoa(userId),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	return ss, err
}
