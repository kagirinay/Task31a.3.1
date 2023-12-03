# Task31a.3.1
Практика по работе с базами данных из Go 31а.3.1

Задание

Для решения задачи от вас требуется следующее:

1. Разработать схему БД PostgresSQL в форме SQL-запроса. Запрос должен быть помещён в файл schema.sql в корневой каталог проекта.
2. По аналогии с пакетом "memdb" разработать пакет "postgres" для поддержки базы данных под управлением СУБД PostgresSQL.
3. По аналогии с пакетом "memdb" разработать пакет "mongo" для поддержки базы данных под управлением СУБД MongoDB.

Контейнер для PostgresSQL на основе dockerfile.
docker build -t my-postgres .
docker run -d --name my-postgres-container -p 5432:5432 my-postgres

Контейнер для MongoDB.
docker run -d --name mongo -p 27017:52017 mongo:4.6.6