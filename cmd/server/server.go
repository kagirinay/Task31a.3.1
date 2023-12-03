package main

import (
	"Task31a.3.1/pkg/api"
	"Task31a.3.1/pkg/storage"
	"Task31a.3.1/pkg/storage/memdb"
	"Task31a.3.1/pkg/storage/mongo"
	"Task31a.3.1/pkg/storage/postgres"
	"log"
	"net/http"
)

// Сервер
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	// Создаём объекты баз данных.

	// БД заглушка
	db1 := memdb.New()
	_ = db1

	// Реляционная БД PostgresSQL.
	db2, err2 := postgres.New("postgres://test_user:qwerty123@localhost:5432/testdb")
	if err2 != nil {
		log.Fatal(err2)
	}
	_ = db2

	// Документная БД MongoDB.
	db3, err3 := mongo.New("mongodb://localhost:27017/")
	if err3 != nil {
		log.Fatal(err3)
	}
	_ = db3

	// Инициализируем хранилище выбранного сервера БД.
	srv.db = db2

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем web-сервер на порту 8080 на всех интерфейсах.
	// Передаём серверу маршрутизатор запросов, поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	http.ListenAndServe(":8080", srv.api.Router())
}
