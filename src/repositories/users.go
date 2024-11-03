package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (usersRepository UsersRepository) Fetch(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	query := `SELECT 
		id, 
		name, 
		nick, 
		email, 
		"createdAt" 
	FROM users 
	WHERE name ILIKE $1 OR nick ILIKE $2`

	rows, err := usersRepository.db.Query(
		query,
		nameOrNick,
		nameOrNick,
	)

	fmt.Println("CHEUGFIE AQUI O")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
