package cmd

import (
	"fmt"

	"log"
	"os"
	"redistr/common"
	server "redistr/server"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "redistr",
	Short: "Redis Traffic replicator",
	Long:  `Redis Traffic replicator`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(viper.GetViper().AllSettings())
		log.Println(common.MapToYamlString(viper.GetViper().AllSettings()))

		primaryopt := redis.Options{
			Addr:     viper.GetViper().GetString("redisserver.primary.address"),
			Password: viper.GetViper().GetString("redisserver.primary.password"),
			DB:       0}

		anabranchopt := redis.Options{
			Addr:     viper.GetViper().GetString("redisserver.anabranch.address"),
			Password: viper.GetViper().GetString("redisserver.anabranch.password"),
			DB:       0}

		server.SetPrimaryRedisClient(primaryopt)
		server.SetAnabranchRedisClient(anabranchopt)

		if viper.GetString("server.port") != "" {
			server.SetServerPort(viper.GetString("server.port"))
		}
		server.Server()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if common.Exists("./config.yml") {
	// 	// Search config in home directory with name ".redistr" (without extension).
	// 	log.Println(common.Exists("./config.yml"))
	// 	viper.AddConfigPath(".")
	// 	viper.SetConfigName("config")
	// 	// viper.SetConfigFile("./config.yml")
	// }

	if cfgFile != "" && common.Exists(cfgFile) {
		// Use config file from the flag.
		log.Println(cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Println(err)
		os.Exit(1)
	}
}
