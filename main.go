package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Movie struct {
	Title     string
	Year      int
	Director string
	Actors    string
	PosterURL string
}

func getMovies() ([]Movie, error) {
	// Получаем переменные окружения для подключения к базе данных
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Формируем строку подключения
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Открываем подключение к базе данных
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Выполняем запрос к базе данных
	rows, err := db.Query("SELECT title, year, director, actors, poster_url FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var m Movie
		if err := rows.Scan(&m.Title, &m.Year, &m.Director, &m.Actors, &m.PosterURL); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	movies, err := getMovies()
	if err != nil {
		http.Error(w, "Ошибка при подключении к базе данных", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	tmpl, err := template.New("movies").Parse(`
		<html>
		<head><title>Фильмы</title></head>
		<body>
			<h1>Информация о фильмах</h1>
			{{range .}}
				<div>
					<h2>{{.Title}} ({{.Year}})</h2>
					<p><strong>Режиссер:</strong> {{.Director}}</p>
					<p><strong>Актеры:</strong> {{.Actors}}</p>
					<img src="{{.PosterURL}}" alt="{{.Title}}" width="200"/>
				</div>
			{{end}}
		</body>
		</html>
	`)

	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, movies)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Запуск сервера на порту 8080...")
	http.ListenAndServe(":8080", nil)
}

