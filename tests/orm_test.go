package tests

import (
	"github.com/gogf/gf/frame/g"
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
	count, err := g.DB().Table("user").Where("id", 3).Count()
	if err != nil {
		panic(err)
	}
	g.Log().Infof("user id=3, count:%v", count)
}