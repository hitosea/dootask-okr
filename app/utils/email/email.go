package email

import (
	"context"
	"dootask-okr/app/constant"
	e "dootask-okr/app/utils/error"
	"fmt"
	"net/http"
	"net/mail"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"gopkg.in/gomail.v2"
)

// EmailService 用于发送邮件的服务
type EmailService struct {
	smtpServer string
	port       int
	account    string
	password   string
}

// NewEmailService 用于创建一个新的EmailService对象
func NewEmailService(smtpServer string, port int, account string, password string) *EmailService {
	return &EmailService{
		smtpServer: smtpServer,
		port:       port,
		account:    account,
		password:   password,
	}
}

// MailGunSend 用于mailgun包发送邮件
func (s *EmailService) MailGunSend(to string, subject string, body string) error {
	// 验证收件人地址是否合法
	if _, err := mail.ParseAddress(to); err != nil {
		return e.New(constant.ErrMailToInvalid)
	}

	// 创建Mailgun客户端
	mg := mailgun.NewMailgun(s.smtpServer, s.password)

	// 设置HTTP客户端
	mg.SetClient(&http.Client{
		Timeout: 10 * time.Second,
	})

	// 设置邮件消息
	message := mg.NewMessage(
		fmt.Sprintf("%s <%s>", "Dootask", s.account),
		subject,
		"",
		to,
	)

	// 设置邮件正文
	message.SetHtml(body)

	// 发送邮件
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, _, err := mg.Send(ctx, message)
	if err != nil {
		if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() {
			return e.New(constant.ErrRequestTimeout)
		} else if mgErr, ok := err.(*mailgun.UnexpectedResponseError); ok && mgErr.Actual == 550 {
			return e.New(constant.ErrMailContentReject)
		} else {
			return err
		}
	}

	return nil
}

// Send 用于发送邮件（默认）
func (s *EmailService) Send(to string, subject string, body string) error {
	if s.smtpServer == "" || strconv.Itoa(s.port) == "" || s.account == "" || s.password == "" {
		return e.New(constant.ErrMailNotConfig)
	}
	m := gomail.NewMessage()
	// 添加别名，也可以直接用<code>m.SetHeader("From", MAIL_USER)</code>
	m.SetHeader("From", " dootask "+"<"+s.account+">")
	// m.SetHeader("To", to...) 	// 发送给多个用户 to []string
	m.SetHeader("To", to)           // 发送给单个用户
	m.SetHeader("Subject", subject) // 设置邮件主题
	m.SetBody("text/html", body)    // 设置邮件正文
	d := gomail.NewDialer(strings.TrimSpace(s.smtpServer), s.port, strings.TrimSpace(s.account), strings.TrimSpace(s.password))
	// 发送邮件
	return d.DialAndSend(m)
}
