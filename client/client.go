package client

import (
	"fmt"
	"github.com/edvargas/go-tests/logging"
)

func PrintSomething() {
	logger := logging.NewLogger("test")

	logger.Info("something")
	fmt.Println("another thing")
}
