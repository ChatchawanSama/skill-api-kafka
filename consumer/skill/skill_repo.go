package skill

import (
	"database/sql"
	"fmt"
	"net/http"
	"consumer/errors"

	"github.com/lib/pq"
)

type skillRepo struct {
	db *sql.DB
}

func NewSkillRepo(db *sql.DB) SkillRepo {
	return &skillRepo{db: db}
}

func (r *skillRepo) PostSkillByKey(skill Skill) error {
	fmt.Println("Entering Post Skill By Key Repo")

	query := "INSERT INTO skill (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5) RETURNING key"
	err := r.db.QueryRow(query, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags)).Scan(&skill.Key)
	if err != nil {
		fmt.Println("Error inserting new skill:", err)
		return errors.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
