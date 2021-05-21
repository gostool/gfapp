package tests

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"testing"
)

// Insert
func TestInsert(t *testing.T) {
	// INSERT INTO `user`(`name`) VALUES('john')
	_, err := g.DB().Table("user").Data(g.Map{"id": 4, "name": "john"}).Insert()
	if err != nil {
		panic(err)
	}
}

// Update
func TestUpdate(t *testing.T) {
	// UPDATE `user` SET `name`='john guo' WHERE name='john'
	_, err := g.DB().Table("user").Data("name", "john guo").
		Where("name", "john").Update()
	if err != nil {
		panic(err)
	}
}

// Delete
func TestDelete(t *testing.T) {
	// DELETE FROM `user` WHERE uid=10
	_, err := g.DB().Table("user").Where("id", 4).Delete()
	if err != nil {
		panic(err)
	}
}

// Select Where
func TestWhere(t *testing.T) {
	// 数量
	//count, err := g.DB().Table("user").Count()
	count, err := g.DB().Table("user").Where("id", 4).Count()
	if err != nil {
		panic(err)
	}
	g.Log().Infof("user id=4, count:%v", count)
}

// ALL: 获取查询结果 Result
func TestAll(t *testing.T) {
	// 也可以简写为 select * from user as t where t.id > 0
	logger := g.Log()
	result, err := g.DB().Table("user").As("t").All("t.id > ?", 0)
	if err != nil {
		logger.Error(err)
		t.Fail()
	}
	for index, value := range result {
		logger.Infof("index:%v row id:%v name:%v\n", index, value["id"], value["name"])
	}
}

func TestRedisOperation(t *testing.T) {
	// redis字符串操作
	g.Redis().Do("SET", "k", "v")
	v, _ := g.Redis().Do("GET", "k")
	g.Log().Info(gconv.String(v))

	// DoVar转换
	v3, _ := g.Redis().DoVar("GET", "k")
	g.Log().Info(v3.String())

	// list
	g.Redis().Do("RPUSH", "keyList", "v5")
	v5, _ := g.Redis().DoVar("LPOP", "keyList")
	g.Log().Info(v5.String())

	// hash
	g.Redis().Do("HSET", "keyHash", "v1", "v6")
	v6, _ := g.Redis().DoVar("HGET", "keyHash", "v1")
	g.Log().Info(v6.String())

	// set
	g.Redis().Do("SADD", "keySet", "v7")
	v7, _ := g.Redis().DoVar("SPOP", "keySet")
	g.Log().Info(v7.String())

	// sort set
	g.Redis().Do("ZADD", "keySortSet", 1, "v8")
	v8, _ := g.Redis().DoVar("ZREM", "keySortSet", "v8")
	g.Log().Info(v8.Int())
}
