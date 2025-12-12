package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "bwanews",
	Short: "Bwa News is a news application",
	Long:  `Bwa News is a news application built with Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Run(startCmd, args)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .env)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		// Load config from cfgFile
		viper.SetConfigFile(cfgFile)
	} else {
		// Load default config
		viper.SetConfigName(`.env`)
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		// Config file found and successfully read
		fmt.Fprintln(os.Stderr,"Using config file:", viper.ConfigFileUsed())
	}
}