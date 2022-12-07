package models

import (
	"time"

	. "github.com/ilies-a/artom-api/app/services"
	"github.com/ilies-a/artom-api/app/utils"
)

type EmailRequestStructure struct {
	Email string `json:"email" binding:"required,email"`
}

type Email struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

func CreateEmailsTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS emails (
		id serial PRIMARY KEY,
		value VARCHAR ( 255 ) UNIQUE NOT NULL,
		created_on TIMESTAMP NOT NULL
	)`
	_, err := DB.Exec(createTableQuery)
	utils.CheckError(err)
}

func DropEmailsTable() {
	dropTableQuery := `DROP TABLE IF EXISTS emails`
	_, err := DB.Exec(dropTableQuery)
	utils.CheckError(err)
}

func GetAllEmails() ([]Email, error) {
	query := `select * from "emails"`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	var (
		id         int
		value      string
		created_at string
		emails     []Email = []Email{}
	)

	for rows.Next() {
		err := rows.Scan(&id, &value, &created_at)
		if err != nil {
			panic(err)
		}
		emails = append(emails, Email{id, value})
	}
	return emails, nil
}

func SaveEmail(email string) error {
	insertStmt := `insert into "emails"("value", "created_on") values($1, $2)`
	_, err := DB.Exec(insertStmt, email, time.Now().UTC())
	return err
}

func DeleteEmail(emailId string) error {
	deleteStmt := `delete from "emails" where id = ` + emailId
	_, err := DB.Exec(deleteStmt)
	return err
}
