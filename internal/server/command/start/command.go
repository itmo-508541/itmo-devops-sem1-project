package start

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

const startServerUse = "start-server"

func New(srv *http.Server, rootCtx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   startServerUse,
		Short: "Start web-server",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("Starting web-server...")
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Println(fmt.Errorf("srv.ListenAndServe: %w", err))
				}
			}()
			log.Printf("Listening on %s", srv.Addr)
			<-rootCtx.Done()

			log.Println("Stopping Web-server...")
			err := srv.Shutdown(context.Background())
			if err != nil {
				log.Println(fmt.Errorf("srv.Shutdown: %w", err))
			}
			log.Println("Web-server stopped")
			return err
		},
	}
}
