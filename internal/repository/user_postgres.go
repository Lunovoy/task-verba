package repository

import (
	"fmt"

	"github.com/Lunovoy/test-project-verba/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) AddUser(user models.User) ([]string, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	var email []string
	addUserQuery := fmt.Sprintf("INSERT INTO %s (gender, title, first_name, last_name, street_number, street_name, city, state, country, postcode, latitude, longitude, timezone_offset, timezone_description, email, username, password, salt, md5, sha1, sha256, dob_date, dob_age, registered_date, registered_age, phone, cell, id_name, id_value, picture_large, picture_medium, picture_thumbnail, nat) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33) RETURNING email", userTable)

	for _, r := range user.Results {
		row := tx.QueryRow(addUserQuery, r.Gender, r.Name.Title, r.Name.First, r.Name.Last, r.Location.Street.Number, r.Location.Street.Name, r.Location.City, r.Location.State, r.Location.Country, r.Location.Postcode, r.Location.Coordinates.Latitude, r.Location.Coordinates.Longitude, r.Location.Timezone.Offset, r.Location.Timezone.Description, r.Email, r.Login.Username, r.Login.Password, r.Login.Salt, r.Login.Md5, r.Login.Sha1, r.Login.Sha256, r.Dob.Date, r.Dob.Age, r.Registered.Date, r.Registered.Age, r.Phone, r.Cell, r.ID.Name, r.ID.Value, r.Picture.Large, r.Picture.Medium, r.Picture.Thumbnail, r.Nat)
		var tempEmail string
		if err := row.Scan(&tempEmail); err != nil {
			tx.Rollback()
			return nil, err
		}
		email = append(email, tempEmail)
	}

	return email, tx.Commit()
}
