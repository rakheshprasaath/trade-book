package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rakheshprasaath/trade-book.git/database"
	"github.com/rakheshprasaath/trade-book.git/models"
	"github.com/rakheshprasaath/trade-book.git/util"
)

func validateEmail(email string) bool {
	re := regexp.MustCompile(`[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}`)
	return re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	// Parse the body of the request
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}

	// Check if password is less than or equal to 6 characters
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 characters",
		})
	}

	// Validate email format
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid email format",
		})
	}

	// Check if email already exists in the database
	database.DB.Where("email = ?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	// Creating a new user object and mapping the input fields
	user := models.User{
		UserName: data["user_name"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
		Phone: data["phone"].(string),
	}

	// Set the password after processing it (assuming SetPassword hashes it)
	user.SetPassword(data["password"].(string))

	// Save user to the database
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"user": user,
		"message": "Account created succesfully",
	})
}


func Login(c *fiber.Ctx) error {
    var data map[string]string

    if err := c.BodyParser(&data); err != nil {
        fmt.Println("Unable to parse body")
    }

    var user models.User
    database.DB.Where("email=?", data["email"]).First(&user)
    if user.Id == 0 {
        c.Status(404)
        return c.JSON(fiber.Map{
            "message": "Email Address doesn't exist, kindly create an account",
        })
    }

    if err := user.ComparePassword(data["password"]); err != nil {
        c.Status(400)
        return c.JSON(fiber.Map{
            "message": "incorrect password",
        })
    }

    token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
    if err != nil {
        c.Status(fiber.StatusInternalServerError)
		return nil
    }

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour *2),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "login successfully",
		"user": user,
	})

    
}
