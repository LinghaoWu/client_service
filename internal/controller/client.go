package controller

import (
	"client_service/consts"
	"client_service/internal/application/clientsvc"
	"client_service/internal/domain/client/entity"
	"client_service/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClientHandler(c *gin.Context) {
	var req *types.ClientReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, RespError(consts.BadRequest, err, "client parameter binding failed"))
		return
	}

	var clientEntity *entity.ClientDO = types.ClientReqToClientDO(req)
	if !clientEntity.IsValidClient() {
		c.JSON(http.StatusOK, RespError(consts.BadRequest, nil, "invalid client data"))
		return
	}

	if err := clientsvc.ServiceImplIns.CreateClientSvc(c, clientEntity); err != nil {
		c.JSON(http.StatusOK, RespError(consts.InternalServerError, err, "failed to create client"))
		return
	}

	resp := types.ClientDOToClientResp(clientEntity)
	c.JSON(http.StatusOK, RespSuccessWithData(resp))
}
