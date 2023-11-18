package actions

import (
	"AirAccountGateway/conf"
	"AirAccountGateway/internal/models/webapi/response"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Bind
// @Tags Instructions
// @Description 绑定钱包指令
// @Accept json
// @Product json
// @Router /api/instructions/bind [post]
// @Param   id     query     string     true  "identity"
// @Success 201
func Bind(ctx *gin.Context) {
	id := strings.TrimSpace(ctx.Query("id"))

	id = idToHash(id, SrcSms)
	api := func() (*http.Response, error) {
		body, _ := json.Marshal(struct {
			Certificate string `json:"certificate"`
		}{
			Certificate: id,
		})

		return http.Post(
			conf.GetNodeHost()+"/wallet/bind",
			"application/json",
			bytes.NewBuffer(body),
		)
	}
	onSuccess := func(v *BaseResponse) {
		response.Success(ctx)
	}
	walletApiInvoker(ctx, api, onSuccess)
}
