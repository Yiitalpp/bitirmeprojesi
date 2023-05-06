package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"project/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"project/models"
)

type UserRepo struct {
	Db *gorm.DB
}

type StatusUser struct {
	Status string
}

func NewUserController() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}

func (repository *UserRepo) Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	// Hash password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}
	user.Password = string(hashedPassword)

	user.ActivationCode = generateActivationCode()
	err = models.Register(repository.Db, &user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	/*
		if err := repository.Db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	*/
	if err := sendActivationEmail(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send activation email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func generateActivationCode() string {
	token := make([]byte, 32)
	rand.Read(token)
	return hex.EncodeToString(token)
}

func sendActivationEmail(user models.User) error {
	// Kullanıcının e-posta adresini alın
	recipientEmail := user.Email

	// E-posta konusu ve içeriği oluşturun
	subject := "Hesap Aktivasyonu Tamamlandı"
	body := "Merhaba " + user.Username + ",\n\nHesabınız başarıyla aktive edildi. Artık giriş yapabilir ve sitemizi kullanabilirsiniz.\n\nTeşekkürler,\nSitemiz Ekibi"

	// Get Sender Name and Sender Email Address from environment variables
	senderName := os.Getenv("SENDER_NAME")
	senderEmailVisible := os.Getenv("SENDER_EMAIL_VISIBLE")

	// E-postanın gönderici adresi ve bilgileri
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Printf("SMTP_PORT is not a valid integer: %s\n", err)
		return err
	}

	// E-posta gövdesini ayarlayın
	//message := []byte("From: " + senderName + "<" + senderEmailVisible + ">" + "\r\n" + "To: " + recipientEmail + "\r\n" + "Subject: " + subject + "\r\n" + "\r\n" + body + "\r\n")
	message := []byte("To: " + recipientEmail + "\r\n" + "From: \"" + senderName + "\" <" + senderEmailVisible + ">\r\n" + "Subject: " + subject + "\r\n" + "\r\n" + body + "\r\n")

	// SMTP sunucusuna bağlanın ve e-postayı gönderin
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)
	err = smtp.SendMail(smtpServer+":"+strconv.Itoa(smtpPort), auth, senderEmail, []string{recipientEmail}, message)
	if err != nil {
		log.Printf("Error sending activation email to %s: %s\n", recipientEmail, err)
		return err
	}

	return nil
}

func (repository *UserRepo) Login(c *gin.Context) {
	var user models.User
	err := models.Login(repository.Db, &user, user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	/*
		// Retrieve user record from database
		var dbUser models.User
		if err := repository.Db.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
	*/

	// Compare hashed passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !user.Active {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account is not activated. Please activate your account."})
		return
	}

	//token, err := models.CreateToken(repository.Db, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	/*
		// Set session cookie
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	*/

	//c.JSON(http.StatusOK, gin.H{"token": token.Token, "start": token.StartingDate, "expiry": token.EndingDate})
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}

func (repository *UserRepo) Logout(c *gin.Context) {
	/*
		// Clear session cookie
		session := sessions.Default(c)
		session.Clear()
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
	*/
}
