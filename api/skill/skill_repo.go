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

func (r *skillRepo) PostSkillByKey(skill Skill) error {
	fmt.Println("Entering Post Skill By Key Repo")

	query := "INSERT INTO skill (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5) RETURNING key"
	err := r.db.QueryRow(query, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags)).Scan(&skill.Key)
	if err != nil {
		fmt.Println("Error inserting new skill:", err)
		return errors.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil

	// fmt.Println("Skill created with Key:", skill.Key)
	// // ctx.JSON(http.StatusCreated, skill)
	// ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": skill})

	// q := "SELECT key, name, description, logo, tags FROM skill where key=$1"
	// row := r.db.QueryRow(q, key)
	// var name, description, logo string
	// var tags pq.StringArray

	// err := row.Scan(&key, &name, &description, &logo, &tags)
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	// return Skill{key, name, description, logo, tags}, nil

}
