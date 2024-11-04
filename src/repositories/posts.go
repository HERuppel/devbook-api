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

func (postsRepository PostsRepository) Get(id uint64) (models.Post, error) {
	query := `
		SELECT p.*, u.nick FROM posts p INNER JOIN users u on p.authorid = u.id WHERE id = $1
	`

	row, err := postsRepository.db.Query(query, id)
	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()

	var post models.Post
	if row.Next() {
		if err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}
