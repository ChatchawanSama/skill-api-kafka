package skill

import (
	"database/sql"
	"fmt"
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

func (r *skillRepo) GetSkillByKey(key string) (Skill, error) {
	// skill := Skill{}

	fmt.Println("Entering getSkillByKey handler")

	q := "SELECT key, name, description, logo, tags FROM skill where key=$1"
	row := r.db.QueryRow(q, key)
	var name, description, logo string
	var tags pq.StringArray

	err := row.Scan(&key, &name, &description, &logo, &tags)
	if err != nil {
		fmt.Println("Error")
	}
	return Skill{key, name, description, logo, tags}, nil

	// skills := []Skill{}
	// query := "SELECT key, name, description, logo, tags FROM skill where key=$1"
	// records, err := r.db.Query(query)
	// if err != nil {
	// 	return []Skill{}, errors.NewError(http.StatusInternalServerError, err.Error())
	// }
	// for records.Next() {
	// 	skill := Skill{}
	// 	err := records.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
	// 	if err != nil {
	// 		return []Skill{}, errors.NewError(http.StatusInternalServerError, err.Error())
	// 	}
	// 	skills = append(skills, skill)
	// }

	// return skills, nil
}
