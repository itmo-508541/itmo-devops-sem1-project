package start

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

const startServerUse = "start-server"

func New(srv *http.Server) *cobra.Command {
	return &cobra.Command{
		Use:   startServerUse,
		Short: "Start web-server",
		// @see https://github.com/sarulabs/di-example/blob/master/main.go
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Listening on %s", srv.Addr)

			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Println(err.Error())
				}
			}()
			stop := make(chan os.Signal, 1)
			signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
			<-stop

			log.Println("Stopping the http server")

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			return srv.Shutdown(ctx)
		},
	}
}
