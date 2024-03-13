package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/todoApp/internal/models"
)

type QuestRepository struct {
	db *pgx.Conn
}

func NewQuestRepository(databaseUrl string) *QuestRepository {
	db, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	return &QuestRepository{db: db}
}

// methods
func (r *QuestRepository) Get(userId int) ([]models.Quest, error) {
	query := `SELECT id, title, description, dificulty, completed FROM quests WHERE user_id = $1`

	rows, err := r.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quests []models.Quest
	for rows.Next() {
		var q models.Quest

		err = rows.Scan(&q.Id, &q.Title, &q.Description, &q.Dificulty, &q.Completed)
		if err != nil {
			return nil, err
		}

		quests = append(quests, q)
	}
	return quests, nil
}

func (r *QuestRepository) GetById(id, userId int) (*models.Quest, error) {
	query := `SELECT id, title, description, dificulty, completed FROM quests WHERE id = $1 AND user_id = $2`
	var quest models.Quest

	row := r.db.QueryRow(context.Background(), query, id, userId)
	err := row.Scan(&quest.Id, &quest.Title, &quest.Description, &quest.Dificulty, &quest.Completed)

	if err != nil {
		return nil, err
	}

	return &quest, nil
}

func (r *QuestRepository) Create(quest models.Quest, userId int) (int, error) {
	var id int

	err := r.db.QueryRow(context.Background(), "INSERT INTO quests (title, description, dificulty, completed, user_id) VALUES($1, $2, $3, $4, $5) RETURNING id", quest.Title, quest.Description, quest.Dificulty, quest.Completed, userId).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *QuestRepository) Update(id, userId int, quest models.Quest) error {
	_, err := r.db.Exec(context.Background(), `UPDATE quests SET title = $1, description = $2, dificulty = $3, completed = $4 WHERE id = $5 AND user_id = $6`,
		quest.Title, quest.Description, quest.Dificulty, quest.Completed, id, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *QuestRepository) Delete(id, userId int) error {
	_, err := r.db.Exec(context.Background(), `DELETE FROM quests WHERE id = $1 AND user_id = $2`, id, userId)
	if err != nil {
		return err
	}
	return nil
}
