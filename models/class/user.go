package class

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id int `orm:"pk"`
	Nick string `orm:"size(60);default()"`
	Email string `orm:"size(100);default()"`
	Phone string `orm:"size(15);default()"`
	Password string `orm:"size(32);default()"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func TestORM(){
	o := orm.NewOrm()
	u := User{
		Nick:"liuyan",
	}
	o.Insert(&u)

}

