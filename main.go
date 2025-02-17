package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Movie struct {
	Title       string
	Year        int
	Director    string
	Actors      string
	PosterURL   string
}

func handler(w http.ResponseWriter, r *http.Request) {
	movies := []Movie{
		{
			Title:     "Бандитский Петербург. Адвокат.",
			Year:      2001,
			Director:  "Петр Буслов",
			Actors:    "Александр Баширов, Сергей Гармаш",
			PosterURL: "https://upload.wikimedia.org/wikipedia/ru/thumb/7/75/%D0%91%D0%B0%D0%BD%D0%B4%D0%B8%D1%82%D1%81%D0%BA%D0%B8%D0%B9_%D0%9F%D0%B5%D1%82%D0%B5%D1%80%D0%B1%D1%83%D1%80%D0%B3._%D0%90%D0%B4%D0%B2%D0%BE%D0%BA%D0%B0%D1%82._%D0%9F%D0%BB%D0%B0%D0%BA%D0%B0%D1%82.jpg/500px-%D0%91%D0%B0%D0%BD%D0%B4%D0%B8%D1%82%D1%81%D0%BA%D0%B8%D0%B9_%D0%9F%D0%B5%D1%82%D0%B5%D1%80%D0%B1%D1%83%D1%80%D0%B3._%D0%90%D0%B4%D0%B2%D0%BE%D0%BA%D0%B0%D1%82._%D0%9F%D0%BB%D0%B0%D0%BA%D0%B0%D1%82.jpg",
		},
		{
			Title:     "Зеленая миля",
			Year:      1999,
			Director:  "Фрэнк Дарабонт",
			Actors:    "Том Хэнкс, Майкл Кларк Дункан",
			PosterURL: "https://upload.wikimedia.org/wikipedia/ru/thumb/0/09/The_Green_Mile_%28poster%29.jpg/500px-The_Green_Mile_%28poster%29.jpg",
		},
		{
			Title:     "Мы Миллеры",
			Year:      2013,
			Director:  "Роберт Швентке",
			Actors:    "Дженнифер Энистон, Джейсон Судейкис",
			PosterURL: "https://upload.wikimedia.org/wikipedia/ru/thumb/0/06/Were_the_Millers_poster.jpg/500px-Were_the_Millers_poster.jpg",
		},
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

