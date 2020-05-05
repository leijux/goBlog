package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DbPing(t *testing.T) {
	err := Db.Ping()
	if err != nil {
		assert.Error(t, err, "发生错误")
	}
}
