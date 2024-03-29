package controllers

import (
	"beegoTest/models/class"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type PublicController struct {
	BaseController
}

func (this *PublicController) Login() {
	this.TplName = "register/login.html"
}

type RET struct {
	Ok      bool        `json:"success"`
	Content interface{} `json:"content"`
}

func (this *PublicController) HeadImg() {
	this.TplName = "register/headimg.html"
}

func (this *PublicController) ChangeHeadImg() {
	ret := RET{
		Ok:      true,
		Content: "success",
	}

	defer func() {
		this.Data["json"] = ret
		this.ServeJSON()
	}()

	webUser := this.GetSession("webUser").(class.User)

	header := this.GetString("head")
	path := fmt.Sprintf("%s%s%s", "/static/img/head/", strconv.Itoa(webUser.Id), ".png")
	err2 := writeFile("."+path, header) //写入文件(字节数组)

	if err2 {
		webUser.Headimg = path
		webUser.Update()
		this.DoLogin(webUser)// 更新session
	}else{
		ret.Ok = false
		ret.Content = "上传失败"
	}

}

func writeFile(path string, base64_img_content string) bool {
	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64_img_content)
	if !b {
		return false
	}
	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)

	base64str := re.ReplaceAllString(base64_img_content, "")

	byte, _ := base64.StdEncoding.DecodeString(base64str)

	err2 := ioutil.WriteFile(path, byte, 0666)

	if err2 != nil {
		panic(err2)
	}
	return true
}

func (this *PublicController) ToLogin() {
	// 登录注册操作
	ret := RET{
		Ok:      true,
		Content: "success",
	}

	defer func() {
		if ret.Ok {
			flash := beego.NewFlash()
			flash.Notice("成功")
			flash.Store(&this.Controller)
			fmt.Println("flash ok")
		}
		this.Data["json"] = ret
		this.ServeJSON()
	}()

	account := this.GetString("account")
	password := this.GetString("password")

	// 验证账号是手机还是邮箱
	valid := validation.Validation{}
	valid.Email(account, "Email")
	phone := ""
	email := account
	if valid.HasErrors() && valid.Errors[0].Key == "Email" {
		valid = validation.Validation{}
		valid.Phone(account, "Phone")
		email = ""
		phone = account
	}

	valid.Required(password, "Password")
	switch {

	case valid.HasErrors():

	default:

		// 数据写入
		u := &class.User{
			Id:       0,
			Email:    email,
			Phone:    phone,
			Nick:     GetNick(),
			Status:   1,
			Password: PwGen(password),
			Created:  time.Now(),
			Headimg:  "/static/img/head/default.jpg",
		}
		isRegister := false
		condition := map[string]string{}
		switch {
		case u.ExistFiled("Email"):
			condition["email"] = u.Email
			isRegister = true
			//valid.Error("用户邮箱被占用")
		case u.ExistFiled("Phone"):
			condition["phone"] = u.Phone
			isRegister = true

			//valid.Error("用户手机被占用")
		default:
			err := u.Create()
			if err != nil {
				valid.Error(fmt.Sprintf("%v", err))
			}
		}

		if isRegister {
			err := u.Read(condition) // 获取用户数据
			if err == orm.ErrNoRows {
				valid.Error("账号存在，或密码不正确")
			}
			if PwCheck(password, u.Password) {
				this.DoLogin(*u)
				return
			} else {
				valid.Error("账号存在，或密码不正确")
			}
		} else {
			fmt.Println("new login", *u)
			this.DoLogin(*u)
			return
		}
	}

	ret.Ok = false
	ret.Content = valid.Errors[0].Key + " " + valid.Errors[0].Message

	return
}

func GetNick() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	no := r.Intn(99999)
	return fmt.Sprintf("呐喊者%d", no)
}

func PwGen(pass string) string {
	salt := strconv.FormatInt(time.Now().UnixNano()%9000+1000, 10)
	return Base64Encode(Sha1(Md5(pass)+salt) + salt)
}

func PwCheck(pwd, saved string) bool {
	saved = Base64Decode(saved)
	if len(saved) < 4 {
		return false
	}
	salt := saved[len(saved)-4:]
	return Sha1(Md5(pwd)+salt)+salt == saved
}

func Sha1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) string {
	res, _ := base64.StdEncoding.DecodeString(s)
	return string(res)
}
