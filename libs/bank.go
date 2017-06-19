package libs

import (
	"bank/libs/mongodb"
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

func (this *Bank) FindAll(where map[string]interface{}) map[string]interface{}{
	iter := this.mongoDb.FindAllData(where)
	row := BankInfo{}
	result := make(map[string]interface{})
	for iter.Next(row) {
		result = map[string]interface{}{
			"id": row.Id,
			"name": row.Name,
			"data": row.Data,
			"time": row.Time,
		}
	}
	return result
}

func (this *Bank) GetCount() int {
	return this.mongoDb.GetCount()
}