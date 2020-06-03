
package services

// ServiceHandler handling the base functionality
type ServiceHandler struct {
	Name string
}

var (
	handler *ServiceHandler
)

func GetHandler() *ServiceHandler {
	return handler
}

