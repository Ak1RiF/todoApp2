package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/todoApp/internal/models"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(databaseUrl string) *UserRepository {
	db, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, passwordhash, avatarurl, sumexperience, amountexperiencetolvl, lvl FROM users WHERE username = $1`
	var user models.User

	row := r.db.QueryRow(context.Background(), query, username)
	err := row.Scan(&user.Id, &user.Username, &user.PasswordHash, &user.AvatarUrl, &user.TotalExperience, &user.AmountExperienceToLvl, &user.Lvl)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetById(id int) (*models.User, error) {
	query := `SELECT id, username, passwordhash, avatarurl, sumexperience, amountexperiencetolvl, lvl FROM users WHERE id = $1`
	var user models.User

	row := r.db.QueryRow(context.Background(), query, id)
	err := row.Scan(&user.Id, &user.Username, &user.PasswordHash, &user.AvatarUrl, &user.TotalExperience, &user.AmountExperienceToLvl, &user.Lvl)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(user models.User) (int, error) {
	var userId int

	err := r.db.QueryRow(context.Background(),
		"INSERT INTO users (username, passwordhash, avatarurl, sumexperience, amountexperiencetolvl, lvl) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		user.Username, user.PasswordHash, user.AvatarUrl, user.TotalExperience, user.AmountExperienceToLvl, user.Lvl).Scan(&userId)
	if err != nil {
		return -1, err
	}

	return userId, nil
}
