package cmd

import (
	logadapter "github.com/rafaelpissolatto/synthetic-monitor-exporter/internal/adapters/log/zerolog"
	"github.com/rafaelpissolatto/synthetic-monitor-exporter/internal/config"
	"github.com/rafaelpissolatto/synthetic-monitor-exporter/internal/core/services"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the synthetic-monitor-exporter server",
	Long:  `Start the synthetic-monitor-exporter server. This server will expose a /metrics endpoint that can be scraped by Prometheus.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("hello synthetic-monitor-exporter")
		runServer()

	},
}

// init initializes the command
func init() {
	RootCmd.AddCommand(ServerCmd)
}

// runServer starts the synthetic-monitor-exporter server
func runServer() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Load the configuration file
	cfgFile := viper.ConfigFileUsed()
	log.Info().Msgf("Loading configuration file: %s", cfgFile)

	config, err := config.NewConfig(cfgFile)
	if err != nil {
		log.Fatal().Msgf("Failed to load configuration file: %v", err)
	}

	// Create a new LogAdapter
	logger := logadapter.NewLogAdapter(config.ServerConfig.LogLevel)

	//TODO: remove this
	// logger.Log(config.HTTPConfig.Endpoints[0].URL)

	// Create a new Server
	logger.Log("Starting http checker...")
	checkerHTTPService := services.NewHTTPService()
	checkerHTTPService.StartCheckerHTTPServer(*config)

}
