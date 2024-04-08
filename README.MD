# Catalog car

Установка зависимостей:

    go mod tidy

Создайте файл .env в корневой директории проекта и настройте его, включая настройки и конфигурационные данные, необходимые для работы вашего сервиса. Пример .env файла:

* DB_USER=your_db_user
* DB_PASSWORD=your_db_password
* DB_NAME=your_db_name
* DB_HOST=your_db_host
* DB_PORT=5432
* PORT=8080

## Установите migrate

Для установки перейдите по ссылке [migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md "сюда") и следуйте инструкциям на странице проекта


### Создайте схему базы данных
    make createDB
### Удалить схему базы данных
    make dropDB

Запустите сервер:

    go run main.go

**Каталог будет доступен по адресу** http://localhost:8080.
REST API Методы

## Документация через Swagger

API документирован с использованием Swagger. Для доступа к документации необходимо запустить API..
