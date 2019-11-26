package nrecho

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/willnewrelic/layout/v3/newrelic"
)

func init() {
	fmt.Println("echo version", echo.Version)
	fmt.Println("agent version", newrelic.Version)
}