package libs

import (
	"bank/libs/mongodb"
	"gopkg.in/mgo.v2-unstable"
)

type Bank struct {
	mongoDb mongodb.MongoDB
}

type BankInfo struct {
	Id string "bson:`id`"
	Name string "bson:`name`"
	Data string "bson:`data`"
	Time string "bson:`time`"
}

//初始化设置collection
func (this *Bank) Init(bankId string) {
	this.mongoDb.SetCollection(bankId)
}

func (this *Bank) Find(where map[string]interface{}) []BankInfo {
	result := []BankInfo{}
	this.mongoDb.FindData(where, &result)
	return result
}

func (this *Bank) FindAll(where map[string]interface{}) *mgo.Iter{
	iter := this.mongoDb.FindAllData(where)
	return iter
}

func (this *Bank) GetCount() int {
	return this.mongoDb.GetCount()
}