package controllers

import (
	"net/http"

	db "root/database"
	"root/model"

	"github.com/gofiber/fiber/v2"
)

// @Summary Создание поста
// @Description Создаёт новый пост с заголовком, содержимым и изображением
// @Tags v2
// @Accept json
// @Produce json
// @Param post body model.Post true "Данные поста"
// @Success 201 {object} model.Post
// @Failure 400 {object} map[string]string "Неверный формат данных"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /v2/posts [post]
func CreatePost(c *fiber.Ctx) error {
	var post model.Post

	// Парсинг тела запроса
	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
			"message": "Не удалось прочитать данные",
		})
	}

	// Сохранение поста в БД
	result := db.DB.DB.Create(&post)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Не удалось создать пост",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Пост успешно создан",
		"post":    post,
	})
}

// @Summary Получение всех постов
// @Description Возвращает список всех постов
// @Tags v2
// @Produce json
// @Success 200 {array} model.Post "Список постов"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /v2/posts [get]
func GetPosts(c *fiber.Ctx) error {
	var posts []model.Post

	// Получение всех постов из БД
	result := db.DB.DB.Find(&posts)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Не удалось получить посты",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
		"posts":  posts,
	})
}

// @Summary Получение поста по ID
// @Description Возвращает пост на основе его уникального идентификатора
// @Tags v2
// @Produce json
// @Param id path int true "ID поста"
// @Success 200 {object} model.Post "Информация о посте"
// @Failure 404 {object} map[string]string "Пост не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /v2/posts/{id} [get]
func GetPostByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var post model.Post

	// Поиск поста по ID
	result := db.DB.DB.First(&post, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "fail",
			"message": "Пост не найден",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
		"post":   post,
	})
}

// @Summary Обновление поста
// @Description Обновляет существующий пост по его уникальному ID
// @Tags v2
// @Accept json
// @Produce json
// @Param id path int true "ID поста"
// @Param post body model.Post true "Обновлённые данные поста"
// @Success 200 {object} model.Post "Обновлённый пост"
// @Failure 400 {object} map[string]string "Неверный формат данных"
// @Failure 404 {object} map[string]string "Пост не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /v2/posts/{id} [put]
func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post model.Post

	// Поиск поста по ID
	result := db.DB.DB.First(&post, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "fail",
			"message": "Пост не найден",
		})
	}

	// Обновление данных поста
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
			"message": "Некорректные данные",
		})
	}

	db.DB.DB.Model(&post).Updates(data)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Пост обновлён",
		"post":    post,
	})
}

// @Summary Удаление поста
// @Description Удаляет существующий пост по его уникальному ID
// @Tags v2
// @Param id path int true "ID поста"
// @Success 200 {object} map[string]string "Пост успешно удалён"
// @Failure 404 {object} map[string]string "Пост не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /v2/posts/{id} [delete]
func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post model.Post

	// Поиск поста по ID
	result := db.DB.DB.First(&post, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "fail",
			"message": "Пост не найден",
		})
	}

	// Удаление поста
	db.DB.DB.Delete(&post)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Пост удалён",
	})
}
