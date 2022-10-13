package cmd

import (
	"context"
	"fmt"
	"generate-deveui-cli/internal/cli"
	config "generate-deveui-cli/internal/configuration"

	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "generate-deveui-cli",
	Short: "An application to generate devEui",
	Long: `An application to generate devEui. It generates a batch of 100 unique DevEUIs and registers them with the
	LoRaWAN API`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		config.Load()
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for _ = range c {
				logrus.Warn("received process interrupt, stopping server gracefully, pls wait")
				cancel()
			}
		}()
		cli.Start(ctx)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
