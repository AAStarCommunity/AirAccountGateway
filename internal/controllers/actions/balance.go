package actions

import (
	"AirAccountGateway/conf"
	"AirAccountGateway/internal/models/webapi/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BalanceResponse struct {
	BaseResponse
	Data string
	Unit string
}

// Balance
// @Tags Instructions
// @Description 查询余额指令
// @Accept json
// @Product json
// @Router /api/instructions/balance [get]
// @Param   id     query     string     true  "identity"
// @Success 200
func Balance(ctx *gin.Context) {

	id := ctx.Query("id")
	id = idToHash(id, SrcSms)
	api := func() (*http.Response, error) {
		return http.Get(conf.GetNodeHost() + "/wallet/getBalance?certificate=" + id)
	}
	onSuccess := func(v *BalanceResponse) {
		v.Unit = "USDT"
		response.Success(ctx, gin.H{
			"status":  v.Status,
			"balance": v.Data,
			"unit":    v.Unit,
		})
	}
	walletApiInvoker(ctx, api, onSuccess)
}
