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
	fmt.Println("checking.....")

	account := models.Account{
		UserId: data["userId"].(string),
		AccountKey: encryptedValue,
		AccountNumber: data["accountNumber"].(string),
		Password: data["password"].(string),
		Server: data["server"].(string),
	}
	fmt.Println("checking.....")

	result := database.DB.Create(&account)
	if result.Error != nil {
		log.Println(result.Error)
		c.Status(500)
	return c.JSON(fiber.Map{
		"error": result.Error,
	})
		

	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"user": account,
		"message": "Account added succesfully",
	})
	

	
	
}


func GetAccounts(c *fiber.Ctx) error {
	var data map[string]interface{}

	// Parse the body of the request
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}

	var accounts []models.Account

	// Enable query logging with Debug() to print SQL in the terminal
	if err := database.DB.Debug().Where("user_id = ?", data["userId"]).Find(&accounts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error fetching accounts",
			"error":   err.Error(),
		})
	}


	// Check if any accounts are found
	if len(accounts) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No accounts found for this user",
		})
	}

	// Return the accounts
	return c.Status(200).JSON(fiber.Map{
		"accounts": accounts,
	})
}


func GetCurrentPositionsByAccountKey(c *fiber.Ctx) error {
	accountKey := c.Params("accountKey") // Get account_key from the route parameter

	var positions []models.CurrentPosition

	// Query database for positions with the given accountKey in descending order
	if err := database.DB.Debug().
		Where("account_key = ?", accountKey).
		Order("time DESC").
		Find(&positions).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error fetching current positions",
			"error":   err.Error(),
		})
	}

	// Check if any positions are found
	if len(positions) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No positions found for this account key",
		})
	}

	// Return the positions
	return c.Status(200).JSON(fiber.Map{
		"positions": positions,
	})
}
