package mail

import (
	"gopkg.in/gomail.v2"
)

type Mail struct {
	User string
	Pass string
	Host string
	Port int
}

func NewMail(data Mail) *Mail {
	return &Mail{User: data.User, Pass: data.Pass, Host: data.Host, Port: data.Port}
}

func (mail *Mail) Send(mailTo []string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "<"+mail.User+">")
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mail.Host, mail.Port, mail.User, mail.Pass)

	err := d.DialAndSend(m)
	return err
}
