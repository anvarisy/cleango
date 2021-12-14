package interfaces

import (
	"cleango/app/domains"
	"database/sql"
	"log"
	"time"
)

type UserRepository struct {
	ConnMySQL *sql.DB
}

func (ur *UserRepository) Save(req domains.User) (string, error) {
	tx := ur.ConnMySQL
	sqlStatement := `INSERT INTO users (id, full_name, place_of_birth, date_of_birth, avatar) VALUES ($1, $2, $3, $4, $5)`
	_, err := tx.Exec(sqlStatement,
		req.ID, req.FullName, req.PlaceOfBirth, req.DateOfBirth, req.Avatar)
	if err != nil {
		panic(err)
	}
	return req.ID, nil
}

func (ur *UserRepository) FindAll() (domains.Users, error) {
	var users domains.Users
	rows, err := ur.ConnMySQL.Query(`
		SELECT
			*
		FROM
			users
	`)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			ID           string
			FullName     string
			PlaceOfBirth string
			DateOfBirth  string
			Avatar       string
		)
		if err = rows.Scan(
			&ID,
			&FullName,
			&PlaceOfBirth,
			&DateOfBirth,
			&Avatar,
		); err != nil {
			return nil, err
		}
		layout := "2006-01-02T15:04:05.999999999Z07:00"
		date, _ := time.Parse(layout, DateOfBirth)
		log.Println(DateOfBirth)
		log.Println(date)
		user := domains.User{
			ID:           ID,
			FullName:     FullName,
			PlaceOfBirth: PlaceOfBirth,
			DateOfBirth:  date,
			Avatar:       Avatar,
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
