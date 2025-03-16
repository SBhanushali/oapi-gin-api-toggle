# OAPI-Gin-API-Toggle

## Overview
OAPI-Gin-API-Toggle is a Go package that provides middleware for managing feature flags in a Gin-based web application. It utilizes OpenAPI specifications extensions to determine which routes are enabled based on the provided feature flags.

## Installation

To install the package, use the following command:
```bash
go get github.com/SBhanushali/oapi-gin-api-toggle
```

## Usage

1. **Import the package:**

```go
import "github.com/SBhanushali/oapi-gin-api-toggle"
```

2. **Example Usage:**

```go

package main

import (
	"github.com/SBhanushali/oapi-gin-api-toggle"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  swagger, _ := spec.GetSwagger() // Get the swagger spec from oapi-codegen generated code

  config := toggle.Config{
    ExtensionName: "x-feature-flag",
    FeatureFlags: map[string]bool{
      "feature1": true,
      "feature2": false,
    },
  }

  router.Use(toggle.New(swagger, config))

  router.Run(":8080")
}
```
Full example can be found [here](https://github.com/SBhanushali/oapi-gin-api-toggle-example)




