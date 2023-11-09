package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	"test-application/v1.0.0/v2/cmd/test-application/config"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

var db *sql.DB

func main() {
	// Инициализация базы данных PostgreSQL
	var err error
	connStr := "postgres://" + config.Config.DB_USER + ":" + config.Config.DB_PASSWD + "@" + config.Config.DB_HOST + ":" + config.Config.DB_PORT + "/" + config.Config.DB + "?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	// Создание экземпляра Gin-приложения
	r := gin.Default()

	// Маршрут для отображения главной страницы
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "templates/index.html", nil)
	})

	// Маршрут для обработки данных пользователя
	r.POST("/greet", func(c *gin.Context) {
		name := c.PostForm("name")

		// Записываем пользователя в базу данных
		_, err := db.Exec("INSERT INTO users (name) VALUES ($1)", name)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "Ошибка при записи в базу данных")
			return
		}

		// Получаем последних 5 пользователей из базы данных
		users, err := getUsers()
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "Ошибка при чтении из базы данных")
			return
		}

		// Отправляем список пользователей на страницу
		c.HTML(http.StatusOK, "templates/greet.html", gin.H{
			"users": users,
		})
	})

	// Запуск веб-сервера на порту 8080
	r.Run(":8080")
}

func getUsers() ([]User, error) {
	// Запрос к базе данных для получения последних 5 пользователей
	rows, err := db.Query("SELECT id, name, created_at FROM users ORDER BY created_at DESC LIMIT 5")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Итерация по результатам запроса и создание списка пользователей
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}