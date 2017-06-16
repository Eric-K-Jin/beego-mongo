package mongodb

import (
	"gopkg.in/mgo.v2-unstable"
	"github.com/astaxie/beego"
	"sync"
)
type MongoDB struct {
	collection *mgo.Collection
}

var (
	mgoDb *mgo.Session
	host = beego.AppConfig.String("host")
	database = beego.AppConfig.String("database")
	lock *sync.Mutex = &sync.Mutex{}
)

//连接db生成session会话
func ConnectDb() *mgo.Session {
	lock.Lock()
	defer lock.Unlock()
	var err error
	if mgoDb == nil {
		mgoDb, err = mgo.Dial(host)
		if err != nil {
			panic(err)
		}
	}

	return mgoDb
}

//获取集合资源
func (this *MongoDB) SetCollection(collectionName string) {
	dbSession := ConnectDb()
	this.collection = dbSession.DB(database).C(collectionName)
}

//获取集合内数据总数
func (this *MongoDB) GetCount() int {
	count, err := this.collection.Count()
	if err != nil {
		panic(err)
	}

	return count
}

//新增数据
func (this *MongoDB) InsertData(data ... interface{}) error {
	return this.collection.Insert(data)
}

//查找所有数据
func (this *MongoDB) FindAllData(where interface{}) *mgo.Iter {
	return this.collection.Find(where).Iter()
}

//查找一条数据
func (this *MongoDB) FindData(where map[string]interface{}, result interface{}) {
	err := this.collection.Find(where).One(&result)
	if err != nil {
		panic(err)
	}
}

//更新数据
func (this *MongoDB) UpdateData(where map[string]interface{}, data map[string]interface{}) error {
	return this.collection.Update(where, data)
}

//删除符合where的一条数据
func (this *MongoDB) RemoveData(where map[string]interface{}) error {
	return this.collection.Remove(where)
}

//删除符合where的所有数据
func (this *MongoDB) removeAllData(where map[string]interface{}) error {
	_, err := this.collection.RemoveAll(where)
	return err
}