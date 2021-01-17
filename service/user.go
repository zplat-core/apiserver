package service

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/saltbo/gopkg/gormutil"
	"github.com/saltbo/gopkg/regexputil"
	"github.com/saltbo/gopkg/strutil"

	"github.com/zplat-core/apiserver/model"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) FindAll(email string, offset, limit int) (list []model.UserFormats, total int64, err error) {
	ut := model.User{}.TableName()
	pt := model.UserProfile{}.TableName()
	query := fmt.Sprintf("left join %s on %s.ux = %s.ux", pt, pt, ut)
	sn := gormutil.DB().Table(ut)
	if email != "" {
		sn = sn.Where("email like ?", fmt.Sprintf("%%%s%%", email))
	}
	sn.Count(&total)
	//sn = sn.Order("id desc")
	err = sn.Offset(offset).Limit(limit).Select("*").Joins(query).Find(&list).Error
	for idx, item := range list {
		list[idx] = *item.Format()
	}
	return
}

func UserEmailExist(email string) (*model.User, bool) {
	return userExist("email", email)
}

func UsernameExist(username string) (*model.User, bool) {
	return userExist("username", username)
}

func UserTicketExist(ticket string) (*model.User, bool) {
	return userExist("ticket", ticket)
}

func userExist(k, v string) (*model.User, bool) {
	user := new(model.User)
	if !gormutil.DB().Where(k+"=?", v).First(user).RecordNotFound() {
		return user, true
	}

	return nil, false
}

type UserCreateOption struct {
	ux        string
	Roles     string
	Ticket    string
	Origin    string
	Activated bool
}

func NewUserCreateOption() UserCreateOption {
	return UserCreateOption{ux: strutil.RandomText(32)}
}

func UserSignup(email, password string, opt UserCreateOption) error {
	if err := UserCreate(email, password, opt); err != nil {
		return err
	} else if opt.Origin == "" {
		return nil
	}

	token, err := TokenCreate(opt.ux, 6*3600, opt.Roles)
	if err != nil {
		return err
	}

	return SignupNotify(email, ActivateLink(opt.Origin, email, token))
}

func UserCreate(email, password string, opt UserCreateOption) error {
	_, exist := UserEmailExist(email)
	if exist {
		return fmt.Errorf("email already exist")
	}

	var parentUser *model.User
	if opt.Ticket != "" {
		pu, exist := UserTicketExist(opt.Ticket)
		if !exist {
			return fmt.Errorf("ticket not exist")
		}

		parentUser = pu
	}

	return gormutil.DB().Transaction(func(tx *gorm.DB) error {
		// 创建基本信息
		user := &model.User{
			Ux:       opt.ux,
			Email:    email,
			Username: fmt.Sprintf("mu%s", strutil.RandomText(18)),
			Password: strutil.Md5Hex(password),
			Roles:    opt.Roles,
			Ticket:   strutil.RandomText(6),
		}
		if opt.Activated {
			user.Status = model.StatusActivated
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// 创建个人信息
		userProfile := &model.UserProfile{
			Ux:       user.Ux,
			Nickname: email[:strings.Index(email, "@")],
		}
		if err := tx.Create(userProfile).Error; err != nil {
			return err
		}

		if parentUser == nil {
			return nil
		}

		// 记录邀请来源
		userInvitation := &model.UserInvitation{
			Ux:    parentUser.Ux,
			SubUx: user.Ux,
		}
		return tx.Create(userInvitation).Error
	})
}

func UserGet(ux string) (*model.User, error) {
	user := new(model.User)
	if gormutil.DB().Where("ux=?", ux).First(user).RecordNotFound() {
		return nil, fmt.Errorf("user not exist")
	}

	return user, nil
}

func UserSignIn(usernameOrEmail, password string) (*model.User, error) {
	userFinder := UsernameExist
	if regexputil.EmailRegex.MatchString(usernameOrEmail) {
		userFinder = UserEmailExist
	}

	user, exist := userFinder(usernameOrEmail)
	if !exist {
		return nil, fmt.Errorf("user not exist")
	}

	if user.Password != strutil.Md5Hex(password) {
		return nil, fmt.Errorf("invalid password")
	}

	return user, nil
}

func UserActivate(ux string) error {
	user, err := UserGet(ux)
	if err != nil {
		return err
	}

	if err := gormutil.DB().Model(user).Update("status", model.StatusActivated).Error; err != nil {
		return err
	}

	return nil
}

// ResetPassword update the new password
func UserPasswordReset(ux, newPwd string) error {
	user, err := UserGet(ux)
	if err != nil {
		return err
	}

	if err := gormutil.DB().Model(user).Update("password", strutil.Md5Hex(newPwd)).Error; err != nil {
		return err
	}
	// record the old password

	return nil
}
