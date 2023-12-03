package postgres

import (
	"Task31a.3.1/pkg/storage"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Store Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

// New Конструктор, принимает строку подключения к БД.
func New(constr string) (*Store, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Store{db: db}
	return &s, err
}

// Posts возвращает список записей из БД.
func (s *Store) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT
			posts.id,
			posts.author_id,
			author.name,
			posts.title,
			posts.content,
			posts.created_at
		FROM posts
		JOIN author ON posts.author_id = author.id
		ORDER BY posts.id;
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []storage.Post
	// Итерирование по результату выполнения запроса и сканирование каждой строки в переменную.
	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.AuthorID,
			&t.AuthorName,
			&t.Title,
			&t.Content,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		// Добавление переменной в массив результатов.
		posts = append(posts, t)
	}
	//Важно не забыть проверить rows.Err()
	return posts, rows.Err()
}

// AddPost создаёт новую запись и возвращает её ID.
func (s *Store) AddPost(t storage.Post) error {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO posts (author_id, title, content)
		VALUES ($1, $2, $3) RETURNING id;
		`,
		t.AuthorID,
		t.Title,
		t.Content,
	).Scan(&id)
	return err
}

// UpdatePost Изменение данных по номеру ID.
func (s *Store) UpdatePost(t storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
		UPDATE posts
		SET
		    title = $2,
		    content = $3
		WHERE id = $1;
	`,
		t.ID,
		t.Title,
		t.Content,
	)
	if err != nil {
		return err
	}

	row := s.db.QueryRow(context.Background(), `
		SELECT
			id,
			author_id,
			title,
			content,
			created_at
		FROM posts
		WHERE id = $1;
		`,
		t.ID,
	)

	var post storage.Post

	err = row.Scan(
		&post.ID,
		&post.AuthorID,
		&post.Title,
		&post.Content,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeletePost Удаляет запись в таблице tasks.
func (s *Store) DeletePost(t storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
		DELETE FROM posts
		WHERE id = $1;
		`,
		t.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

var posts = []storage.Post{
	{
		ID:      1,
		Title:   "Effective Go",
		Content: "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.",
	},
	{
		ID:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
	},
}
