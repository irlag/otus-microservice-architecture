package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"otus-microservice-architecture/app/api/responses"
	"otus-microservice-architecture/app/processors"
)

type Healthcheck struct {
	processors *processors.Processors
}

func NewHealthcheckApi(processors *processors.Processors) *Healthcheck {
	return &Healthcheck{
		processors: processors,
	}
}

func (h *Healthcheck) HandleMethods(router *mux.Router) {
	router.HandleFunc(routes.Healthcheck.Path, h.Check()).Methods(routes.Healthcheck.Method)
}

func (h *Healthcheck) Check() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		status, err := h.processors.HealthcheckProcessor.Check()
		if err != nil {
			response := responses.NewErrorResponse(http.StatusBadRequest, err)
			response.WriteErrorResponse(writer)

			return
		}

		healtcheckOk := responses.NewHealthcheckOkResponse(status)
		healtcheckOk.WriteResponse(writer)
	}
}
