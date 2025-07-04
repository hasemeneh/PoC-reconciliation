package public

import (
	"net/http"

	"github.com/hasemeneh/PoC-OnlineStore/helper/httphandler"
	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/service"
	"github.com/julienschmidt/httprouter"
)

type Public struct {
	Service *service.Service
}

func NewHandler(Service *service.Service) *Public {
	return &Public{
		Service: Service,
	}
}
func (p *Public) Register(r *httprouter.Router) {
	apiHttpHandlers := httphandler.New("/api", r)
	apiHttpHandlers.GET("/ping", p.PING)
	apiHttpHandlers.GET("/reconcile", p.HandleGetReconcileReport)
	apiHttpHandlers.POST("/reconcile", p.HandleReconcile)
}
func (p *Public) PING(r *http.Request) response.HttpResponse {
	return response.NewJsonResponse().SetMessage("Pong").SetData("Pung")
}
