package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sheng/air"
)

func main() {
	air.GET("/", func(req *air.Request, res *air.Response) error {
		time.Sleep(5 * time.Second)
		return res.String("Finished.")
	})

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := air.Serve(); err != nil {
			air.ERROR("server error", map[string]interface{}{
				"error": err.Error(),
			})
		}
	}()

	<-shutdownChan
	air.INFO("shutting down the server")
	air.Shutdown(0)
	air.INFO("server gracefully stopped")
}
