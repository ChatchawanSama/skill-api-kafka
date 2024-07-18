package skill

import (
	"database/sql"
	"net/http"
	"skill-api-kafka/errors"

	"github.com/lib/pq"
)

type skillRepo struct {
	db *sql.DB
}

func NewSkillRepo(db *sql.DB) SkillRepo {
	return &skillRepo{db: db}
}

func (r *skillRepo) GetSkills() ([]Skill, error) {

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
