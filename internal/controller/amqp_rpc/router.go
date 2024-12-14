package amqprpc

import (
	"github.com/dariuszdroba/go-from-template/internal/usecase"
	"github.com/dariuszdroba/go-from-template/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
