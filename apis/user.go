package apis

import (
	"fmt"
	"net/http"
	"time"

	"task-system/log"
	"task-system/models/user"
	"task-system/src/common"

	"github.com/gin-gonic/gin"
)

//AddUserAPI 添加用户
func AddUserAPI(c *gin.Context) {
	var u user.User
	err := c.Bind(&u)
	if err != nil {
		msg := fmt.Sprintln("shoul bind err")
		log.Logger.Errorln(err)
		common.Rmsg(c, http.StatusOK, msg, false)
		return
	}
	t := time.Now()
	u.Created = t
	u.Updated = t
	id, err := u.AddUser()
	if err != nil {
		msg := fmt.Sprintln("add user err")
		log.Logger.Errorln(err)
		common.Rmsg(c, http.StatusOK, msg, false)
		return
	}

	msg := fmt.Sprintf("add user id")
	common.Rmsg(c, http.StatusOK, msg, id)
}
