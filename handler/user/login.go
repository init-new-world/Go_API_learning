package user

import "C"
import (
	"fmt"
	"net/http"

	"github.com/init-new-world/Go_API_learning/pkg/token"

	"github.com/init-new-world/Go_API_learning/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/init-new-world/Go_API_learning/model"
	"github.com/init-new-world/Go_API_learning/pkg/errno"
	"github.com/init-new-world/Go_API_learning/util"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

type LoginRequest struct {
	Username string
	Password string
}

func Login(ctx *gin.Context) {
	log.Info("User Login function called.", lager.Data{"X-Request-Id": util.GetReqID(ctx)})
	var req LoginRequest
	status := http.StatusOK
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(status, errno.ErrorJSON(errno.ErrBind))
		return
	}
	log.Infof("Use HTTP Method: %s", ctx.Request.Method)
	log.Debugf("Username: %s,Password: %s", req.Username, req.Password)

	newUser := &model.User{
		Username: req.Username,
		Password: req.Password,
	}

	compUser := &model.User{Username: req.Username}

	if err := compUser.Get(); err != nil {
		err := errno.New(errno.ErrUserNotFound, fmt.Errorf("User cannot found."))
		log.Errorf(err, "User_Login_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	if err := auth.Compare(compUser.Password, newUser.Password); err != nil {
		err := errno.New(errno.ErrPasswordIncorrect, fmt.Errorf("User password incorrect."))
		log.Errorf(err, "User_Login_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	tkn, err := token.Sign(ctx, token.Context{ID: compUser.ID, Username: compUser.Username}, "")
	if err != nil {
		err := errno.New(errno.ErrToken, fmt.Errorf("Error token signed."))
		log.Errorf(err, "User_Login_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	ctx.Header("Authorization", fmt.Sprintf("Initial %s", tkn))
	ctx.JSON(http.StatusOK, errno.OK)
}
