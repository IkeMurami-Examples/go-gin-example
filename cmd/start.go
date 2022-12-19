/*
Copyright © 2022 Ike Murami murami.ike@gmail.com
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	ginserver "github.com/IkeMurami-Examples/go-gin-example/pkg/cmd"
	"github.com/IkeMurami-Examples/go-gin-example/pkg/utils"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		debugMode := viper.GetBool("debug")
		ctx := context.Background()
		logger, _ := zap.NewProduction()
		logger.Info("Start asman", zap.Bool("debug_mode", debugMode))
		if debugMode {
			logger, _ = zap.NewDevelopment()

			for _, key := range viper.GetViper().AllKeys() {
				logger.Debug("Config", zap.Any(key, viper.Get(key)))
			}
		}

		// Add the Zap logger to context
		ctx = utils.ContextWithLogger(ctx, logger)

		if err := ginserver.StartServer(ctx); err != nil {
			logger.Warn("Couldn't start the asman server", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
