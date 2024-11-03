package repositories

import (
	"api/src/models"
	"database/sql"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (usersRepository UsersRepository) Create(user models.User) (uint64, error) {
	query := `
		INSERT INTO 
			users 
				(name, nick, email, password)
			VALUES 
				($1, $2, $3, $4)
			RETURNING id	
		`

	statement, err := usersRepository.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var ID int
	err = statement.QueryRow(user.Name, user.Nick, user.Email, user.Password).Scan(&ID)
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}
