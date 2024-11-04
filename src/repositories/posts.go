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
		SELECT p.*, u.nick FROM posts p INNER JOIN users u on p.authorid = u.id WHERE p.id = $1
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

func (postsRepository PostsRepository) Fetch(userId uint64) ([]models.Post, error) {
	query := `
		SELECT DISTINCT 
			p.*,
			u.nick
		FROM posts p
		INNER JOIN users u ON p.authorid = u.id
		LEFT JOIN followers f ON p.authorid = f.userId
		WHERE u.id = $1 OR f.followerId = $2
		ORDER BY p.id DESC
	`

	rows, err := postsRepository.db.Query(query, userId, userId)
	if err != nil {
		return []models.Post{}, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return []models.Post{}, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (postsRepository PostsRepository) Update(postId uint64, post models.Post) error {
	query := `
		UPDATE posts
			SET
				title = $1,
				content = $2
			WHERE id = $3
	`

	statement, err := postsRepository.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postId); err != nil {
		return err
	}

	return nil
}

func (postsRepository PostsRepository) Delete(postId uint64) error {
	query := `
		DELETE 
		FROM posts
		WHERE id = $1
	`

	statement, err := postsRepository.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

func (postsRepository PostsRepository) FetchByUser(userId uint64) ([]models.Post, error) {
	query := `
		SELECT DISTINCT 
			p.*,
			u.nick
		FROM posts p
		INNER JOIN users u ON p.authorid = u.id
		WHERE p.authorid = $1
		ORDER BY p.id DESC
	`

	rows, err := postsRepository.db.Query(query, userId)
	if err != nil {
		return []models.Post{}, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return []models.Post{}, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (postsRepository PostsRepository) Like(postId uint64) error {
	query := `
		UPDATE 
			posts
		SET
			likes = likes + 1
		WHERE id = $1
	`

	statement, err := postsRepository.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

func (postsRepository PostsRepository) Deslike(postId uint64) error {
	query := `
		UPDATE 
			posts
		SET
			likes = CASE WHEN likes > 0 THEN likes - 1 ELSE likes END
		WHERE id = $1
	`

	statement, err := postsRepository.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}
