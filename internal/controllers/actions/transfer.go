package actions

import (
	"AirAccountGateway/conf"
	"AirAccountGateway/internal/models/webapi/request"
	"AirAccountGateway/internal/models/webapi/response"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Transfer
// @Tags Instructions
// @Description 转账指令
// @Accept json
// @Product json
// @Router /api/instructions/transfer [post]
// @Param   id     			query     string     			true  "identity"
// @Param   ContentBody     body      request.Transfer     	true  "Request"
// @Success 202
func Transfer(ctx *gin.Context) {

	id := ctx.Query("id")
	id = idToHash(id, SrcSms)
	var req request.Transfer
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	req.Receiver = idToHash(req.Receiver, SrcSms)
	api := func() (*http.Response, error) {
		body, _ := json.Marshal(struct {
			From  string `json:"from"`
			To    string `json:"to"`
			Value string `json:"value"`
		}{
			From:  id,
			To:    req.Receiver,
			Value: req.Value,
		})

		return http.Post(
			conf.GetNodeHost()+"/wallet/transfer",
			"application/json",
			bytes.NewBuffer(body),
		)
	}
	onSuccess := func(v *BaseResponse) {
		response.Success(ctx, gin.H{
			"success": true,
			"op":      v.Data,
		})
	}
	walletApiInvoker(ctx, api, onSuccess)
}
