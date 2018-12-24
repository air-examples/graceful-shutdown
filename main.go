package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aofei/air"
)

func main() {
	a := air.Default
	a.DebugMode = true

	a.GET("/", func(req *air.Request, res *air.Response) error {
		time.Sleep(5 * time.Second)
		return res.WriteString("Finished.")
	})

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := a.Serve(); err != nil {
			a.ERROR("server error", map[string]interface{}{
				"error": err.Error(),
			})
		}
	}()

	<-shutdownChan
	a.INFO("shutting down the server")
	a.Shutdown(0)
	a.INFO("server gracefully stopped")
}
