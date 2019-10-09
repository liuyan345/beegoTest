package class

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)


type User struct {
	Id int `orm:"pk;auto"`
	Nick string `orm:"size(60)"`
	Email string `orm:"size(100)"`
	Phone string `orm:"size(15)"`
	Password string `orm:"size(64)"`
	Status int8 `orm:"default(1)" description:"状态,1 正常 2 禁用"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

// 多字段唯一索引
func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"Email", "Phone"},
	}
}

func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	if err := o.Read(u);err != nil {
		panic(err)
	}
	return
}

func (u *User) Read(condition map[string]string) error {
	cond := orm.NewCondition()

	for field,value := range condition {
		cond = cond.And(field, value)
	}

	o := orm.NewOrm()

	return  o.QueryTable(beego.AppConfig.String("PGY::prefix") + "user").SetCond(cond).One(u)
}

func (u *User) Create() (err error){
	o := orm.NewOrm()
	_, err = o.Insert(u)
	return
}

func (u User) Update() (err error){
	o := orm.NewOrm()
	_, err = o.Update(&u)
	return
}

// 验证字段是否存在
func (u User) ExistFiled(filed string) bool {
	o := orm.NewOrm()
	value := ""
	if filed == "Email" {
		value = u.Email
	}else if filed == "Phone" {
		value = u.Phone
	}
	if value == ""{
		return false
	}else{
		return o.QueryTable(beego.AppConfig.String("PGY::prefix") + "user").Filter(filed,value).Exist()
	}

}



