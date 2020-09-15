package httpserver

import (
	"net/http"
	"strconv"

	"github.com/dapperkop/blank/apiserver/httpserver/handler/test"
	"github.com/dapperkop/blank/consts"
	"github.com/dapperkop/blank/logger"
	"github.com/dapperkop/blank/types"
	"github.com/gorilla/mux"
)

var (
	// Router var ...
	Router = newRouter()
	server = &http.Server{
		Handler:      Router,
		ReadTimeout:  consts.DefaultAPIHTTPReadTimeout,
		WriteTimeout: consts.DefaultAPIHTTPWriteTimeout,
	}
)

func newRouter() *mux.Router {
	var router = mux.NewRouter()

	router.HandleFunc("/ping", test.Ping).Methods("GET")

	return router
}

// Run func ...
func Run() {
	logger.Logger.Fatalln(server.ListenAndServe())
}

// Setup func ...
func Setup(config types.HTTP) {
	server.Addr = config.Host + ":" + strconv.FormatInt(int64(config.Port), 10)
}
