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

func (usersRepository UsersRepository) Get(id uint64) (models.User, error) {
	query := `
		SELECT 
			id,
			name,
			nick,
			email,
			"createdAt"
		FROM users
		WHERE id = $1
	`

	row, err := usersRepository.db.Query(query, id)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (usersRepository UsersRepository) Update(id uint64, user models.User) error {
	query := `
		UPDATE users
			SET 
				name = $1,
				nick = $2,
				email = $3
			WHERE id = $4
	`

	statement, err := usersRepository.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (usersRepository UsersRepository) Delete(id uint64) error {
	query := `DELETE FROM users WHERE id = $1`

	statement, err := usersRepository.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (usersRepository UsersRepository) GetByEmail(email string) (models.User, error) {
	query := `
		SELECT 
			id,
			password
		FROM users
		WHERE email = $1
	`

	row, err := usersRepository.db.Query(query, email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, nil
		}
	}

	return user, nil
}

func (usersRepository UsersRepository) Follow(id, followerId uint64) error {
	query := `
		INSERT INTO followers (userId, followerId) values ($1, $2) ON CONFLICT (userId, followerId) DO NOTHING
	`

	statement, err := usersRepository.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id, followerId); err != nil {
		return err
	}

	return nil
}

func (usersRepository UsersRepository) Unfollow(id, followerId uint64) error {
	query := `
		DELETE FROM followers WHERE userId = $1 AND followerId = $2
	`

	statement, err := usersRepository.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id, followerId); err != nil {
		return err
	}

	return nil
}

func (usersRepository UsersRepository) FetchFollowers(id uint64) ([]models.User, error) {
	query := `
		SELECT
			u.id,
			u.name,
			u.nick,
			u.email,
			u."createdAt"
		FROM users u 
		inner join followers f on u.id = f.followerId where f.userId = $1
	`

	rows, err := usersRepository.db.Query(query, id)
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

func (usersRepository UsersRepository) FetchFollowing(id uint64) ([]models.User, error) {
	query := `
		SELECT
			u.id,
			u.name,
			u.nick,
			u.email,
			u."createdAt"
		FROM users u 
		inner join followers f on u.id = f.userId where f.followerId = $1
	`

	rows, err := usersRepository.db.Query(query, id)
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
