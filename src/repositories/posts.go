package repositories

import (
	"api/src/models"
	"database/sql"
)

type PostsRepository struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *PostsRepository {
	return &PostsRepository{db}
}

func (postsRepository PostsRepository) Create(post models.Post) (uint64, error) {
	query := `
		INSERT INTO posts (title, content, authorid) VALUES ($1, $2, $3) RETURNING ID
	`

	statement, err := postsRepository.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var ID int
	err = statement.QueryRow(post.Title, post.Content, post.AuthorID).Scan(&ID)
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}
