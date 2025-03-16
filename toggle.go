package toggle

import (
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gin-gonic/gin"
)

type Config struct {
	// ExtensionName is the name of the extension starting with x-*. This is the name of the extension that will be used to check for feature flags.
	ExtensionName string
	// FeatureFlags is a map of feature flag assigned to extension name and their enabled status
	FeatureFlags map[string]bool
}

func New(swagger *openapi3.T, config Config) gin.HandlerFunc {
	router, err := gorillamux.NewRouter(swagger)
	if err != nil {
		panic(fmt.Sprintf("failed to create router: %v", err))
	}
	return func(c *gin.Context) {
		route, _, _ := router.FindRoute(c.Request)
		if route == nil {
			// Request doesn't match any route from the swagger therefore out of scope for feature flags
			c.Next()
			return
		}
		if featureFlag, ok := route.Operation.Extensions[config.ExtensionName]; ok {
			if !config.FeatureFlags[featureFlag.(string)] {
				c.AbortWithStatus(http.StatusNotFound)
			}
		}
		c.Next()
	}
}
