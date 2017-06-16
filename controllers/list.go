package controllers

import (
	"github.com/astaxie/beego"
	"bank/libs"
	"fmt"
)

type ListController struct {
	beego.Controller
	bank libs.Bank
}

type BankInfos struct {
	Id string "bson:`id`"
	Name string "bson:`name`"
	Data string "bson:`data`"
	Time string "bson:`time`"
}

func (list *ListController) Get() {
	list.Data["Website"] = "Eric Jin"
	list.Data["Email"] = "352512482@qq.com"
	list.TplName = "index.html"
}

func (list *ListController) ShowList() {

	list.EnableRender = false

	bankId := list.Ctx.Input.Param(":bankId")

	list.bank.Init(bankId)

	fmt.Println(list.bank.GetCount())
}