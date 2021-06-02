package testsdb

import (
	"github.com/gogf/gf/frame/g"
	"testing"
)

// Delete
func TestDelete(t *testing.T) {
	// DELETE FROM `user` WHERE uid=10
	_, err := g.DB().Table("user").Where("id", 4).Delete()
	if err != nil {
		panic(err)
	}
}

// Insert
func TestInsert(t *testing.T) {
	// INSERT INTO `user`(`passport`) VALUES('john')
	_, err := g.DB().Table("user").
		Data(g.Map{"id": 4, "passport": "john", "password": "john", "nickname": "littleJohn"}).Insert()
	if err != nil {
		panic(err)
	}
}

// Update
func TestUpdate(t *testing.T) {
	// UPDATE `user` SET `passport`='john guo' WHERE passport='john'
	_, err := g.DB().Table("user").Data("passport", "john guo").
		Where("passport", "john").Update()
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
		logger.Infof("index:%v row id:%v passport:%v\n", index, value["id"], value["passport"])
	}
}
