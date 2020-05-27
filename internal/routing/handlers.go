package routing

import (
	"errors"
	"fmt"

	"github.com/logrusorgru/aurora"

	"github.com/gin-gonic/gin"
)

func getRequestHandler(handlerID string) (gin.HandlerFunc, error) {

	// switch handlerID {

	// case handlersSet.HandlerBillingDebtInquiry:
	// 	return billing.GetHandler().GetBillInfo, nil

	// case handlersSet.HandlerGetToken:
	// 	return auth.TokenGeneratorhandler, nil

	// case handlersSet.HandlerBillingFacade:
	// 	return billing.GetHandler().GetServicesFacade, nil
	// }

	for hID, handler := range handlersMap {
		if hID == handlerID {
			return handler, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("There are no url handlers for the presented handler id: {%s} . you should check "+
		"your identity file and check if there are any handlers defined in the internal/routing/init.go\nCurrently, "+
		"the following handlers are available %v",
		aurora.Yellow(handlerID), aurora.Red(handlersSet)))
}

func RegisterHandler(handler gin.HandlerFunc, handlerID string) {

	if _, ok := handlersMap[handlerID]; !ok {
		handlersMap[handlerID] = handler

	}

}
