package email

import (
	"bytes"
	"errors"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"html/template"
)

type Notification interface {
	SetAuth() Authentication
	SendEmail() (err error)
}

type EmailNotif struct {
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

func (e *EmailNotif) SetAuth() Authentication {
	return Authentication{
		Identity: "",
		Username: viper.GetString("notification.email.sender"),
		Password: viper.GetString("notification.email.password"),
		Host:     viper.GetString("notification.email.host"),
	}
}

func (e *EmailNotif) SendEmail() (err error) {
	errChan := make(chan error, 1)

	go func() {
		auth := e.SetAuth()

		// set data to body
		t, err := template.New("webpage").Parse(e.Message)
		if err != nil {
			errChan <- errors.New("Error parsing template: " + err.Error())
			return
		}

		var msg bytes.Buffer
		err = t.Execute(&msg, e.Data)
		if err != nil {
			errChan <- errors.New("Error set data template: " + err.Error())
			return
		}

		// send email
		m := gomail.NewMessage()
		m.SetHeader("From", viper.GetString("notification.email.sender"))
		m.SetHeader("To", e.Receiver...)
		m.SetHeader("Subject", e.Subject)
		m.SetBody("text/html", msg.String())

		d := gomail.NewDialer(auth.Host, 587, auth.Username, auth.Password)

		if err = d.DialAndSend(m); err != nil {
			errChan <- err
			return
		}

		errChan <- nil
	}()

	err = <-errChan
	return err
}
