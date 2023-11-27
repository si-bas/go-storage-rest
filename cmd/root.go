/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/si-bas/go-storage-rest/config"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-storage-rest",
	Short: "Storage Service REST API",
	Long:  `The Storage Service REST API provides a set of endpoints to manage and interact with a storage system. It enables users to perform operations such as uploading, downloading, and deleting files.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is in current directory .env)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = "./.env"
	}

	var k = koanf.New(".")

	// Load dotenv config from file
	if err := k.Load(file.Provider(cfgFile), dotenv.ParserEnv("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	})); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", cfgFile)
	}

	// Load environment variables and merge into the loaded config
	k.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil)

	if err := k.Unmarshal("", &config.Config); err != nil {
		fmt.Fprintln(os.Stderr, "failed to unmarshal config to struct variable", err.Error())
	}
}
