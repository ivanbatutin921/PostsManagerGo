package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/swaggo/fiber-swagger"

	"root/controllers"
	"root/model"

	db "root/database"
	_"root/docs"
	//middleware "root/middleware"
	//model "root/model"
)

// @title API для авторизации и постов
// @version 1.0
// @description Swagger-документация для API регистрации, авторизации и управления постами.
// @host localhost:3000
// @BasePath /

type Server struct {
	app  *fiber.App
	port string
}

func (s *Server) allRoutes() {

	v1 := s.app.Group("v1")
	v1.Post("/singup", controllers.SignUp)
	v1.Post("/login", controllers.Login)

	v2 := s.app.Group("v2")                         // Маршруты для CRUD постов
	v2.Post("/posts", controllers.CreatePost)       // Создать пост
	v2.Get("/posts", controllers.GetPosts)          // Получить все посты
	v2.Get("/posts/:id", controllers.GetPostByID)   // Получить пост по ID
	v2.Put("/posts/:id", controllers.UpdatePost)    // Обновить пост по ID
	v2.Delete("/posts/:id", controllers.DeletePost) // Удалить пост по ID

	s.app.Get("/swagger/*", fiberSwagger.WrapHandler)

}

func NewServer(port string) *Server {
	s := &Server{
		app:  fiber.New(),
		port: port,
	}

	//s.app.Use(logger.New())
	//s.app.Static("/image", "./image")
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	return s
}

func (s *Server) Run() {
	s.allRoutes()

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	log.Fatal(s.app.Listen(":" + s.port))

	//log.Fatal(s.app.ListenTLS(":"+s.port, "./certs/minica.pem", "./certs/minica-key.pem"))
}

func main() {

	db.ConnectToDB()
	db.DB.DB.AutoMigrate(&model.User{}, &model.Post{})

	s := NewServer("3000")
	s.Run()

}
