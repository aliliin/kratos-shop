package server

import (
	"github.com/google/wire"
)

// ProviderSet is s providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer)
