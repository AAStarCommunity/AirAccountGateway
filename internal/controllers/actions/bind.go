package actions

import (
	"AirAccountGateway/conf"
	"AirAccountGateway/internal/models/webapi/response"
	"bytes"
	"encoding/json"
	"fmt"
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

	fmt.Println("[GET] /wallet/check: " + conf.GetNodeHost() + "/wallet/check?certificate=" + id)
	if resp, err := http.Get(
		conf.GetNodeHost() + "/wallet/check?certificate=" + id,
	); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusUnprocessableEntity, &msg)
		return
	} else {
		if resp.StatusCode == 201 {
			msg := "airaccount already exists"
			response.Fail(ctx, http.StatusNotAcceptable, &msg)
			return
		}
	}

	api := func() (*http.Response, error) {
		body, _ := json.Marshal(struct {
			Certificate string `json:"certificate"`
		}{
			Certificate: id,
		})

		fmt.Println("[POST] /wallet/bind: " + conf.GetNodeHost() + "/wallet/bind" + " {" + string(body) + "}")
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
