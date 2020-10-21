package user

import (
	"fmt"
	"net/http"

	"github.com/init-new-world/Go_API_learning/model"

	"github.com/gin-gonic/gin"
	"github.com/init-new-world/Go_API_learning/pkg/errno"
	"github.com/init-new-world/Go_API_learning/util"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Get(ctx *gin.Context) {
	log.Info("User Get function called.", lager.Data{"X-Request-Id": util.GetReqID(ctx)})
	status := http.StatusOK
	username := ctx.Param("username")
	log.Infof("Use HTTP Method: %s", ctx.Request.Method)
	log.Debugf("Username: %s", username)

	newUser := &model.User{
		Username: username,
	}

	if test := newUser.NewRecord(); test {
		err := errno.New(errno.ErrUserNotFound, fmt.Errorf("User cannot found."))
		log.Errorf(err, "User_Get_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	if err := newUser.Get(); err != nil {
		err := errno.New(errno.ErrDatabase, fmt.Errorf("Database get error."))
		log.Errorf(err, "User_Get_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	ctx.JSON(status, newUser)
}
