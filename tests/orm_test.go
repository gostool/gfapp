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

