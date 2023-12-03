package storage

import "errors"

// Post - публикация.
type Post struct {
	ID          int
	Title       string
	Content     string
	AuthorID    int
	AuthorName  string
	CreatedAt   int64
	PublishedAt int64
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Post() ([]Post, error) // Получение всех публикаций.
	AddPost(Post) error    // Создание новой публикации.
	UpdatePost(Post) error // Обновление публикации.
	DeletePost(Post) error // Удаление публикации по ID.
}

var ErrPostNotFound = errors.New("пост не найден")
