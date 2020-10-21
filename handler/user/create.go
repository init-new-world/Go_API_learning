package user

import (
	"fmt"
	"net/http"

	"github.com/lexkong/log/lager"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
	"github.com/init-new-world/Go_API_learning/model"
	"github.com/init-new-world/Go_API_learning/pkg/errno"
	"github.com/init-new-world/Go_API_learning/util"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

func Create(ctx *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(ctx)})
	var req CreateRequest
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

	if err := newUser.Validate(); err != nil {
		err = errno.New(errno.ErrValidation, fmt.Errorf("Validation error."))
		log.Errorf(err, "User_Create_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	if err := newUser.Encrypt(); err != nil {
		err = errno.New(errno.ErrEncrypt, fmt.Errorf("Encrypt error."))
		log.Errorf(err, "User_Create_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	if test := newUser.NewRecord(); !test {
		err := errno.New(errno.ErrUserAlreadyExist, fmt.Errorf("Username exist."))
		log.Errorf(err, "User_Create_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	if err := newUser.Create(); err != nil {
		err = errno.New(errno.ErrDatabase, fmt.Errorf("Database create error."))
		log.Errorf(err, "User_Create_Error")
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	}

	resp := &CreateResponse{
		Username: newUser.Username,
	}
	ctx.JSON(status, resp)
}
