package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-learning/configs"
	"go-learning/databases/connection"
	"go-learning/databases/migration"
	"go-learning/modules/car"
	"go-learning/modules/employee"
	"go-learning/modules/user"
	. "go-learning/utils/email"
	"go-learning/utils/logger"
	"go-learning/utils/redis"
	"go-learning/utils/swagger"
	"time"
)

// @title Swagger Documentation
// @version 1.1.2
// @description This is documentation go_learning.
// @host localhost:8080
func main() {
	configs.Initiator()

	logger.Initiator()

	redis.Initiator()

	connection.Initiator()
	defer connection.DBConnections.Close()

	migration.Initiator(connection.DBConnections)

	template := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>Notification Email</title></head><body style="font-family:Arial,sans-serif;background-color:#f4f4f4;margin:0;padding:0"><div style="max-width:600px;margin:20px auto;background-color:#fff;border:1px solid #e0e0e0;border-radius:5px;box-shadow:0 4px 8px rgba(0,0,0,.1)"><div style="background-color:#007bff;color:#fff;padding:20px;text-align:center;border-top-left-radius:5px;border-top-right-radius:5px"><h1 style="margin:0;font-size:24px">Notification</h1></div><div style="padding:20px;font-size:16px;line-height:1.6;color:#333"><h2 style="color:#007bff">Hello, {{.Name}}</h2><p>We wanted to let you know about a recent activity on your account.</p><h3 style="color:#007bff">Login Activity Details</h3><table style="width:100%;border-collapse:collapse;margin-top:10px"><tr><th style="text-align:left;padding:8px;background-color:#f4f4f4;border:1px solid #ddd;color:#007bff">Device</th><td style="padding:8px;border:1px solid #ddd">{{.Device}}</td></tr><tr><th style="text-align:left;padding:8px;background-color:#f4f4f4;border:1px solid #ddd;color:#007bff">IP Address</th><td style="padding:8px;border:1px solid #ddd">{{.IPAddress}}<br><span style="font-size:14px;color:#666">Note: The IP address represents the location from which the login occurred.</span></td></tr><tr><th style="text-align:left;padding:8px;background-color:#f4f4f4;border:1px solid #ddd;color:#007bff">Location</th><td style="padding:8px;border:1px solid #ddd">{{.Location}}</td></tr><tr><th style="text-align:left;padding:8px;background-color:#f4f4f4;border:1px solid #ddd;color:#007bff">Login Time</th><td style="padding:8px;border:1px solid #ddd">{{.LoginTime}}</td></tr></table><p>If you have any questions, feel free to<a href="mailto:support@example.com" style="color:#007bff;text-decoration:none"> contact our support team</a>.</p></div><div style="background-color:#f4f4f4;padding:10px;text-align:center;font-size:12px;color:#666;border-bottom-left-radius:5px;border-bottom-right-radius:5px"><p>&copy; 2024 Your Company. All rights reserved.</p><p><a href="https://www.example.com/unsubscribe" style="color:#007bff;text-decoration:none">Unsubscribe</a></p></div></div></body></html>`
	userNotification := Email{
		Sender:  viper.GetString("notification.email.sender"),
		Subject: "Testing Email Notification",
		Receiver: []string{
			"arifyuniarto88@gmail.com",
		},
		Message: template,
		Data: map[string]string{
			"Device":    "Samsung",
			"Name":      "Arif",
			"IPAddress": "127.0.0.1",
			"Location":  "San Francisco",
			"LoginTime": time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	userNotification.SendEmail(&gin.Context{})

	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()

	swagger.Initiator(router)

	car.Initiator(router)
	user.Initiator(router)
	employee.Initiator(router)

	router.Run(":8080")
}
