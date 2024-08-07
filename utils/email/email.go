package email

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-learning/utils/logger"
	"gopkg.in/gomail.v2"
	"html/template"
	"net/smtp"
)

type Notification interface {
	SetAuth() smtp.Auth
	SendEmail(ctx *gin.Context)
}

type Email struct {
	Sender   string
	Receiver []string
	Subject  string
	Message  string
	Data     map[string]string
}

type Authentication struct {
	Identity string
	Username string
	Password string
	Host     string
}

func (e *Email) SetAuth() Authentication {
	return Authentication{
		Identity: "",
		Username: viper.GetString("notification.email.sender"),
		Password: viper.GetString("notification.email.password"),
		Host:     viper.GetString("notification.email.host"),
	}
}

func (e *Email) SendEmail(ctx *gin.Context) {
	go func() {
		auth := e.SetAuth()

		// set data to body
		t, err := template.New("webpage").Parse(e.Message)
		if err != nil {
			err = errors.New("Error parsing template: " + err.Error())
			logger.ErrorWithCtx(ctx, nil, err)
		}

		var msg bytes.Buffer
		err = t.Execute(&msg, e.Data)
		if err != nil {
			err = errors.New("Error set data template: " + err.Error())
			logger.ErrorWithCtx(ctx, nil, err)
		}

		// send email
		m := gomail.NewMessage()
		m.SetHeader("From", viper.GetString("notification.email.sender"))
		m.SetHeader("To", e.Receiver...)
		m.SetHeader("Subject", e.Subject)
		m.SetBody("text/html", msg.String())

		d := gomail.NewDialer(auth.Host, 587, auth.Username, auth.Password)

		if err = d.DialAndSend(m); err != nil {
			logger.ErrorWithCtx(ctx, nil, err)
		}
	}()
}
