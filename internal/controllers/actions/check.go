package actions

import (
	"AirAccountGateway/conf"
	"AirAccountGateway/internal/models/webapi/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TransferCheck
// @Tags Instructions
// @Description 转账结果查询指令
// @Accept json
// @Product json
// @Router /api/instructions/transfer/check [get]
// @Param   op     query     string     true  "transfer op"
// @Success 200
func TransferCheck(ctx *gin.Context) {

	op := ctx.Query("op")
	api := func() (*http.Response, error) {

		return http.Get(
			conf.GetNodeHost() + "/wallet/checkOp?op=" + op,
		)
	}
	onSuccess := func(v *BalanceResponse) {
		response.Success(ctx, gin.H{
			"status": v.Status,
		})
	}
	walletApiInvoker(ctx, api, onSuccess)
}
