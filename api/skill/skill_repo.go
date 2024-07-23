package skill

import (
	"api/errors"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lib/pq"
)

type skillRepo struct {
	db *sql.DB
}

func NewSkillRepo(db *sql.DB) SkillRepo {
	return &skillRepo{db: db}
}

func (r *skillRepo) GetSkillsRepo() ([]Skill, error) {

	skills := []Skill{}
	query := "SELECT key, name, description, logo, tags FROM skill"
	records, err := r.db.Query(query)
	if err != nil {
		return []Skill{}, errors.NewError(http.StatusInternalServerError, err.Error())
	}
	for records.Next() {
		skill := Skill{}
		err := records.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
		if err != nil {
			return []Skill{}, errors.NewError(http.StatusInternalServerError, err.Error())
		}
		skills = append(skills, skill)
	}

	return skills, nil
}

func (r *skillRepo) GetSkillByKeyRepo(key string) (Skill, error) {
	fmt.Println("Entering Get Skill By Key Repo")

	q := "SELECT key, name, description, logo, tags FROM skill where key=$1"
	row := r.db.QueryRow(q, key)
	var name, description, logo string
	var tags pq.StringArray

	err := row.Scan(&key, &name, &description, &logo, &tags)
	if err != nil {
		fmt.Println("Error")
	}
	return Skill{key, name, description, logo, tags}, nil
}