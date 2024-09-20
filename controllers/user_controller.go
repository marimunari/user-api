package controllers

import (
	"user-api/auth"
	"user-api/database"
	"user-api/models"
	"user-api/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Method for registering a user
func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(utils.ErrFailedToParse.Code).JSON(utils.ErrFailedToParse)
	}

	var existingUser models.User
	if err := database.DB.Where("email = ?", data["email"]).First(&existingUser).Error; err == nil {
		return c.Status(utils.ErrEmailExists.Code).JSON(utils.ErrEmailExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(utils.ErrFailedToHashPassword.Code).JSON(utils.ErrFailedToHashPassword)
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: hashedPassword,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(utils.ErrFailedToCreate.Code).JSON(utils.ErrFailedToCreate)
	}

	return c.Status(utils.SuccessUserRegistered.Code).JSON(utils.SuccessUserRegistered)
}

// Method for user to log in
func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(utils.ErrFailedToParse.Code).JSON(utils.ErrFailedToParse)
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		return c.Status(utils.ErrUserNotFound.Code).JSON(utils.ErrUserNotFound)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil {
		return c.Status(utils.ErrInvalidCredentials.Code).JSON(utils.ErrInvalidCredentials)
	}

	accessToken, refreshToken, err := auth.GenerateToken(data["email"])
	if err != nil {
		return c.Status(utils.ErrFailedToGenerateToken.Code).JSON(utils.ErrFailedToGenerateToken)
	}

	responseData := fiber.Map{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return c.Status(utils.SuccessLogin.Code).JSON(models.SuccessResponse{
		APIResponse: models.APIResponse{
			Code:    utils.SuccessLogin.Code,
			Message: utils.SuccessLogin.Message,
			Data:    responseData,
		},
	})
}

// Method for user to log out
func Logout(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(utils.ErrAuthorizationHeaderMissing.Code).JSON(utils.ErrAuthorizationHeaderMissing)
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	auth.AddTokenToBlacklist(tokenString)

	return c.Status(utils.SuccessLogout.Code).JSON(utils.SuccessLogout)
}

// Protected | Method to get user details
func GetUserDetail(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*models.Claims)

	if claims == nil {
		return c.Status(utils.ErrUnauthorized.Code).JSON(utils.ErrUnauthorized)
	}

	var user models.User
	if err := database.DB.Where("email = ?", claims.Email).First(&user).Error; err != nil {
		return c.Status(utils.ErrUserNotFound.Code).JSON(utils.ErrUserNotFound)
	}

	responseData := fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}

	return c.Status(utils.SuccessGetUserDetail.Code).JSON(models.SuccessResponse{
		APIResponse: models.APIResponse{
			Code:    utils.SuccessGetUserDetail.Code,
			Message: utils.SuccessGetUserDetail.Message,
			Data:    responseData,
		},
	})
}

// Method to refresh tokens
func RefreshToken(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(utils.ErrFailedToParse.Code).JSON(utils.ErrFailedToParse)
	}

	refreshToken := data["refresh_token"]
	claims, err := auth.ValidateToken(refreshToken, true)
	if err != nil {
		return c.Status(utils.ErrInvalidToken.Code).JSON(utils.ErrInvalidToken)
	}

	accessToken, newRefreshToken, err := auth.GenerateToken(claims.Email)
	if err != nil {
		return c.Status(utils.ErrFailedToGenerateToken.Code).JSON(utils.ErrFailedToGenerateToken)
	}

	responseData := fiber.Map{
		"access_token":  accessToken,
		"refresh_token": newRefreshToken,
	}

	return c.Status(utils.SuccessLogin.Code).JSON(models.SuccessResponse{
		APIResponse: models.APIResponse{
			Code:    utils.SuccessLogin.Code,
			Message: "Tokens atualizados com sucesso",
			Data:    responseData,
		},
	})
}
