package handlers

import (
	"github.com/FarhadHabibie/dock_go/database"
	"github.com/FarhadHabibie/dock_go/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Welcome to go api App!")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func DeleteFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Delete(&facts, "question LIKE ?", "%siapa%")

	return c.Status(200).JSON(facts)
}

func DropFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Migrator().DropTable(&facts)

	return c.Status(200).JSON(facts)
}
