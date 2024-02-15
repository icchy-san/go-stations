package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	svc := service.NewTODOService(todoDB)
	panicHandler := handler.NewPanicHandler()

	mux.HandleFunc("/todos", handler.NewTODOHandler(svc).ServeHTTP)
	mux.Handle("/healthz", middleware.BasicAuthentication(middleware.UserAudit(middleware.AccessLog(handler.NewHealthzHandler()))))
	mux.Handle("/do-panic", middleware.Recovery(panicHandler))
	return mux
}
