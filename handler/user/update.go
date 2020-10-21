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

type UpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Update(ctx *gin.Context) {
	log.Info("User Update function called.", lager.Data{"X-Request-Id": util.GetReqID(ctx)})
	var req UpdateRequest
	status := http.StatusOK
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(status, errno.ErrorJSON(errno.OK))
		return
	}
	log.Infof("Use HTTP Method: %s", ctx.Request.Method)
	log.Debugf("Username: %s,Password: %s", req.Username, req.Password)

	newUser := &model.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := newUser.Validate(); err != nil {
		err = errno.New(errno.ErrValidation, fmt.Errorf("Validation error."))
		log.Errorf(err, "User_Update_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	if test := newUser.NewRecord(); test {
		err := errno.New(errno.ErrUserNotFound, fmt.Errorf("User cannot found."))
		log.Errorf(err, "User_Update_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	if err := newUser.Update(); err != nil {
		err = errno.New(errno.ErrDatabase, fmt.Errorf("Database update error."))
		log.Errorf(err, "User_Update_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	resp := &UpdateResponse{
		Username: newUser.Username,
		Password: newUser.Password,
	}
	ctx.JSON(status, resp)
}
