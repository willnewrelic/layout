package nrecho

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/willnewrelic/layout/v3/newrelic"
)

func init() {
	fmt.Println("echo error", echo.ErrUnsupportedMediaType.Error())
	fmt.Println("agent version", newrelic.Version)
}
