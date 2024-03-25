# XSIS ASSIGNMENT TEST

- [XSIS ASSIGNMENT TEST](#xsis-assignment-test)
  - [HOW TO USE](#how-to-use)
  - [API DOCUMENTATION](#api-documentation)
  - [ADVANCED CONFIGURATION](#advanced-configuration)

## HOW TO USE

- Run PostgreSQL with docker: `docker compose up -d db`
- Run main program: `go run .` and check out the [API Documentation](#api-documentation)
- Run main program with docker: `docker compose up -d api`
- Run unit tests: `go test -cover ./...`

## API DOCUMENTATION

- Add new Movie
  ```bash
  curl --request POST --url 'http://localhost:3000/Movies' --header 'Content-Type: application/json' --data '{"title":"judul1","description":"deskripsi","image":"gambar","rating":3}'
  ```
- Get list of Movies
  ```bash
  # default page is 1 and limit is 10
  curl --request GET --url 'http://localhost:3000/Movies'

  # custom page and limit
  curl --request GET --url 'http://localhost:3000/Movies?page=1&limit=1'
  ```
- Get Movie details
  ```bash
  curl --request GET --url 'http://localhost:3000/Movies/:movie_id'

  # sample
  curl --request GET --url 'http://localhost:3000/Movies/b91cdd9e-ea96-11ee-a760-b6a8b9f1dd47'
  ```
- Update Movie
  ```bash
  curl --request PATCH --url 'http://localhost:3000/Movies/b91cdd9e-ea96-11ee-a760-b6a8b9f1dd47' --header 'Content-Type: application/json' --data '{"title":"judul baru","description":"deskripsi baru","image":"gambar baru","rating":5}'
  ```
- Delete Movie
  ```bash
  curl --request DELETE --url 'http://localhost:3000/Movies/:movie_id'

  # sample
  curl --request DELETE --url 'http://localhost:3000/Movies/b91cdd9e-ea96-11ee-a760-b6a8b9f1dd47'
  ```

## ADVANCED CONFIGURATION

Set environment variables on your operating system (or docker-compose.yml service api if run main program on docker) with the following options:
  - DATABASE_URL: the database connection, default is `postgres://postgres:postgres@localhost:5432/xsis?sslmode=disable`.
  - ENVIRONMENT: the application environment, default is `development`.
  - PORT: application running port number, default is `3000`.
  - APP_TZ: application time zone, default is `Asia/Jakarta`.