package main
import (
	"fmt"
	"github.com/jegacs/simpleapi/handlers"
)

func main() {
	handlers.SetHelloWorldHandler()
	handlers.SetShortenUrlHandler()
	fmt.Println("Server running in :8080")
	handlers.Run(":8080")
}