package service

import (
	"fmt"

	"github.com/saltbo/gopkg/mailutil"
)

func SignupNotify(email, link string) error {
	template := `
       <h3>账户激活链接</h3>
       <p><a href="%s">点击此处账户激活</a></p>
		<p>如果您没有进行账号注册请忽略！</p>
       `
	body := fmt.Sprintf(template, link)
	return mailutil.Send("账号注册成功，请激活您的账户", email, body)
}

func PasswordResetNotify(email, link string) error {
	template := `
       <h3>密码重置链接</h3>
       <p><a href="%s">点击此处重置密码</a></p>
		<p>如果您没有申请重置密码请忽略！</p>
       `
	body := fmt.Sprintf(template, link)
	return mailutil.Send("密码重置申请", email, body)
}
