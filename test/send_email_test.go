package test

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendEmail(t *testing.T) {
	os.Setenv("MAILPASSWORD", "填写自己的授权码")
	MAILPASSWORD := os.Getenv("MAILPASSWORD")
	defer os.Unsetenv("MAILPASSWORD") // 清理环境变量，避免影响其他测试

	toUserEmail, code := "10240362@qq.com", "666666"

	fmt.Println("MAILPASSWORD => ", MAILPASSWORD)
	e := email.NewEmail()
	e.From = "Get <light_blue_dre@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")

	e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "light_blue_dre@163.com", MAILPASSWORD, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}
