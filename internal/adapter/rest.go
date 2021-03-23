package adapter

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mkfsn/rimokon/internal/domain"
)

type handler struct {
	http.Handler

	rimokonUsecase domain.RimokonUsecase
}

func NewHandler(rimokonUsecase domain.RimokonUsecase) http.Handler {
	router := mux.NewRouter()

	handler := handler{
		Handler:        router,
		rimokonUsecase: rimokonUsecase,
	}

	router.HandleFunc("/health", handler.healthChecker)
	router.HandleFunc("/device/{device}/command/{command}", handler.handleDeviceCommand)

	return handler.Handler
}

func (h *handler) handleDeviceCommand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		device  = vars["device"]
		command = vars["command"]
	)

	now := time.Now()
	defer func() {
		log.Printf("Handle request device=%s command=%s latency=%s\n", device, command, time.Now().Sub(now))
	}()

	err := h.rimokonUsecase.ExecuteDeviceCommand(r.Context(), device, command)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) healthChecker(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
