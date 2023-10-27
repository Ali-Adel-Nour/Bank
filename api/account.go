package api

import "github.com/gin-gonic/gin"

type createAccountRequest struct {
	Owner    string `json:"owner"binding:"required"`
	Currency string `json:"currency"binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
var req createAccountRequest
if err := ctx.ShouldBindJSON(&req); err != nil {
	ctx.JSON(http.StatusBadRequest,errResponse(err) )
	return
}
arg:= db.CreateAccountParams{
	Owner: req.Owner,
	Currency: req.Currency,
	Balance : 0 ,
}
account,err := server.store.CreateAccount(ctx,arg){
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errResponse(err))
		return
	}
	func(server *Server) Start(addres string) error {
		server.Router.Run(address)
	}
	ctx.JSON(http.StatusOK,account)
}
}
