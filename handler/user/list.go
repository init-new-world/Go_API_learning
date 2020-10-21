package user

import (
	"net/http"

	"github.com/init-new-world/Go_API_learning/service"

	"github.com/init-new-world/Go_API_learning/model"

	"github.com/gin-gonic/gin"
	"github.com/init-new-world/Go_API_learning/pkg/errno"
	"github.com/init-new-world/Go_API_learning/util"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint          `json:"totalCount"`
	UserList   []*model.User `json:"userList"`
}

func List(ctx *gin.Context) {
	log.Info("User List function called.", lager.Data{"X-Request-Id": util.GetReqID(ctx)})
	var req ListRequest
	status := http.StatusOK
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(status, errno.ErrorJSON(errno.OK))
		return
	}
	log.Infof("Use HTTP Method: %s", ctx.Request.Method)

	if infos, count, err := service.ListUser(req.Username, req.Offset, req.Limit); err != nil {
		ctx.JSON(status, errno.ErrorJSON(err))
		return
	} else {
		ctx.JSON(status, ListResponse{
			TotalCount: count,
			UserList:   infos,
		})
	}
}
