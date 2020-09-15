package apiserver

import (
	"github.com/dapperkop/blank/apiserver/httpserver"
	"github.com/dapperkop/blank/types"
)

// Setup func ...
func Setup(config types.APIServer) {
	httpserver.Setup(config.HTTP)
}
