package class

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Msg struct {
	Id int `orm:"pk;auto"`
	Content string `orm:"size(225)"`
	Status int8 `orm:"default(1)" description:"状态,1 正常 2 禁用"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	User *User `orm:"rel(fk)"`
}

func (m *Msg) Create() error{
	o := orm.NewOrm()
	_, err := o.Insert(m)
	return err
}

func (m *Msg) ReadDB(){
	o := orm.NewOrm()
	o.Read(m)
}

func (msgs Msg) ReadList(limit,offset int) (ret []Msg,tn int64) {

 	o := orm.NewOrm()

 	if msgs.User != nil {
		tn, _ = o.QueryTable(beego.AppConfig.String("PGY::prefix") + "msg").Filter("User", msgs.User).RelatedSel().Count()
		o.QueryTable(beego.AppConfig.String("PGY::prefix") + "msg").Filter("User",msgs.User).RelatedSel().Limit(limit).Offset(offset).OrderBy("-created").All(&ret)
	}else{
		tn, _ = o.QueryTable(beego.AppConfig.String("PGY::prefix") + "msg").RelatedSel().Count()
		o.QueryTable(beego.AppConfig.String("PGY::prefix") + "msg").RelatedSel().Limit(limit).Offset(offset).OrderBy("-created").All(&ret)
	}

	return
}

func (m Msg) Delete() (err error){
	o := orm.NewOrm()
	_, err = o.Delete(&m)
	return
}
