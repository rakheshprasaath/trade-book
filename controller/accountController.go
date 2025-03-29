package controller

import (
	"fmt"
	"strings"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rakheshprasaath/trade-book.git/database"
	"github.com/rakheshprasaath/trade-book.git/models"
)

func AddAccount(c *fiber.Ctx)error{
	var data map[string]interface{}
	var accountData models.Account

	// Parse the body of the request
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}

	// Check if email already exists in the database
	database.DB.Where("account_number = ?", strings.TrimSpace(data["accountNumber"].(string))).First(&accountData)
	if accountData.AccountNumber != "" {
        c.Status(400)
		return c.JSON(fiber.Map{
			"message": "account already exists",
		})
    }

	// Concatenate the variables into the string "accountNumber:server:password"
	combinedString := fmt.Sprintf("%s:%s:%s", data["accountNumber"], data["server"], data["password"])

	// Encryption key
	key := "mySuperSecureKey"

	encryptedValue, err := models.EncryptAccountId(key, combinedString)
	if err != nil {
		fmt.Println("Error encrypting:", err)
	}
	

	account := models.Account{
		UserId: data["userId"].(string),
		AccountId: encryptedValue,
		AccountNumber: data["accountNumber"].(string),
		Password: data["password"].(string),
		Server: data["server"].(string),
	}

	result := database.DB.Create(&account)
	if result != nil {
		log.Println(result)
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"user": account,
		"message": "Account added succesfully",
	})
	

	
	
}