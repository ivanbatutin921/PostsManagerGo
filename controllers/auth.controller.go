package controllers

import (
	"net/http"

	db "root/database"
	"root/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Регистрация пользователя
// @Description Создание нового пользователя с ролью "user"
// @Tags v1
// @Accept json
// @Produce json
// @Param login body string true "Логин пользователя"
// @Param password body string true "Пароль пользователя"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]string
// @Router /v1/signup [post]
func SignUp(c *fiber.Ctx) error {
	var body struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
			"message": "Не удалось прочитать данные",
		})
	}

	var user model.User
	db.DB.DB.First(&user, "login = ?", body.Login)

	if user.ID != 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
			"message": "Пользователь уже существует",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Не удалось хэшировать пароль",
		})
	}

	user = model.User{
		Login:    body.Login,
		Password: string(hash),
		Role:     "user", // user ставиим по умолчанмю
	}

	result := db.DB.DB.Create(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Не удалось создать пользователя",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Пользователь успешно создан",
		"user": fiber.Map{
			"id":    user.ID,
			"login": user.Login,
			"role":  user.Role,
		},
	})
}

// @Summary Авторизация пользователя
// @Description Авторизация с логином и паролем
// @Tags v1
// @Accept json
// @Produce json
// @Param login body string true "Логин пользователя"
// @Param password body string true "Пароль пользователя"
// @Success 200 {object} model.User
// @Failure 401 {object} map[string]string
// @Router /v1/login [post]
func Login(c *fiber.Ctx) error {
	var body struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
			"message": "Не удалось прочитать данные",
		})
	}

	var user model.User
	db.DB.DB.First(&user, "login = ?", body.Login)

	if user.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
			"message": "Пользователь не найден",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"status":  "fail",
			"message": "Неверный пароль",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Успешный вход",
		"user": fiber.Map{
			"id":    user.ID,
			"login": user.Login,
			"role":  user.Role,
		},
	})
}
