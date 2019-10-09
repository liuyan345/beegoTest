package class

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	Id int `orm:"pk;auto"`
	Content string `orm:"size(225)"`
	Status int8 `orm:"default(1)" description:"状态,1 正常 2 禁用"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	User *User `orm:"rel(fk)"`
	Msg *Msg `orm:"rel(fk)"`
}

func (c *Comment) Create() error{
	o := orm.NewOrm()
	_, err := o.Insert(c)
	return err
}

func (c Comment) ReadList() (ret []Comment) {
 	o := orm.NewOrm()
 	o.QueryTable(beego.AppConfig.String("PGY::prefix") + "comment").Filter("Msg", c.Msg).RelatedSel().OrderBy("created").All(&ret)
	return
}

func (c Comment) Delete() (err error){
	o := orm.NewOrm()
	_, err = o.Delete(&c)
	return
}
