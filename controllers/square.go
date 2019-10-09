package controllers

import (
	"beegoTest/models/class"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type SquareController struct {
	BaseController
	ret RET
}

type MsgList struct {
	Msg class.Msg
	Comment []class.Comment
}

func (this *SquareController) SquareData() {
	//this.CheckLogin()
	// 获取用户id
	msg := class.Msg{}

	if this.Ctx.Input.Param(":id") != "" {
		uid,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
		user := &class.User{Id:uid}
		user.ReadDB()
		msg.User = user
		this.Data["pageUrl"] = "/" + this.Ctx.Input.Param(":id")
	}else{
		this.Data["pageUrl"] = "/"
	}
	var offset int
	if this.GetString("pageNo") == "" {
		offset = 0
	}else{
		offset,_ = strconv.Atoi(this.GetString("pageNo"))
		offset = offset - 1
	}

	limit := 10

	data, tn := msg.ReadList(limit, offset)
	msgList := map[int] MsgList{}

	for i,v := range(data) {
		comment := class.Comment{}
		cmsg := &class.Msg{Id:v.Id}
		cmsg.ReadDB()
		comment.Msg = cmsg
		commentList := comment.ReadList()
		msgList[i] = MsgList{v,commentList}
	}

	// 分页信息
	formatInt := strconv.FormatInt(tn, 10)
	total,_ := strconv.Atoi(formatInt)
	this.Data["pageInfo"] = PageInfo(limit,offset,total)

	beego.ReadFromRequest(&this.Controller)

	this.Data["msgList"] = msgList
	this.TplName = "square/index.html"
}

type pageNum struct {
	No int
	Status int
}

func PageInfo(limit int, offset int, total int) interface{} {
	// 显示5个页码 是单数
	var pageNums int = 5
	pageInfo := map[string] interface{}{}

	pageInfo["firstPage"] = pageNum{1,1}

	endPage := total / limit

	i := total % limit

	if i != 0 {
		endPage += 1
	}
	pageInfo["endPage"] = pageNum{endPage,1}
	var pageList = map[int] pageNum{}

	if endPage < pageNums {
		for i = 0; i< endPage; i++ {
			if offset == i {
				pageList[i] = pageNum{i+1,1}
			}else {
				pageList[i] = pageNum{i+1,0}
			}
		}
	}else{
		if offset <= 2 {
			for i = 0;i < pageNums; i++ {
				if offset == i {
					pageList[i] = pageNum{i+1,1}
				}else {
					pageList[i] = pageNum{i+1,0}
				}
			}
		}else if offset > pageNums/2 && offset < endPage - pageNums/2 {
			for i = offset - pageNums/2; i <= offset + pageNums/2 ; i++ {
				if offset == i {
					pageList[i] = pageNum{i+1,1}
				}else {
					pageList[i] = pageNum{i+1,0}
				}
			}
		}else {
			for i = endPage - pageNums; i < endPage ; i++ {
				if offset == i {
					pageList[i] = pageNum{i+1,1}
				}else {
					pageList[i] = pageNum{i+1,0}
				}
			}
		}
	}

	pageInfo["pageList"] = pageList

	return pageInfo
}

func (this *SquareController) AddPage() {
	this.CheckLogin()
	this.TplName = "square/add.html"
}

func (this *SquareController) ToAdd() {
	this.CheckLogin()
	if this.GetString("content") == "" {
		this.ret.Ok = false
		this.ret.Content = "内容为空，不能发布"
		this.Data["json"] = this.ret
		this.ServeJSON()
		return
	}
	user := this.GetSession("webUser").(class.User)
	msg := &class.Msg{
		Content:strings.TrimSpace(this.GetString("content")),
		Status:1,
		User:&user,
	}
	err := msg.Create()
	if err != nil {
		this.ret.Ok = false
		this.ret.Content = "发布失败，请重新发布"
	}

	this.ret.Ok = true
	this.Data["json"] = this.ret
	this.ServeJSON()
	return
}

func (this *SquareController) Del() {
	this.CheckLogin()
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	m := class.Msg{Id:id}
	err := m.Delete()

	if err != nil {
		this.ret.Ok = false
		this.ret.Content = "撤销失败"
	}else {
		this.ret.Ok = true
	}

	this.Data["json"] = this.ret
	this.ServeJSON()
}

func (this *SquareController) CommentAdd(){
	this.CheckLogin()
	if this.GetString("content") == "" {
		this.ret.Ok = false
		this.ret.Content = "内容为空，不能评论"
		this.Data["json"] = this.ret
		this.ServeJSON()
		return
	}
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	m := &class.Msg{Id:id}
	m.ReadDB()

	user := this.GetSession("webUser").(class.User)

	comment := &class.Comment{
		Content:strings.TrimSpace(this.GetString("content")),
		Status:1,
		User:&user,
		Msg:m,
	}
	err := comment.Create()
	if err != nil {
		this.ret.Ok = false
		this.ret.Content = "评论失败，请重新评论"
	}

	this.ret.Ok = true
	this.Data["json"] = this.ret
	this.ServeJSON()
	return
}

func (this *SquareController) CommentDel(){

}