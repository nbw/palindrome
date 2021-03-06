/*
Note: Almost all of the code here was generated by Cobra's CLI
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/nbw/palindrome/palindrome"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "palindrome",
	Short: "Determines if a phrase is a paldindrome",
	Long:  "Determines if a phrase is a paldindrome",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("[Error]: Input required.\n\nExample: palindrome \"racecar\"")
			return
		}

		if len(args) > 1 {
			fmt.Println("[Error]: Don't forget to surround your input in quotations.\n\nExample: palindrome \"A Toyota’s a Toyota\"")
			return
		}

		result := palindrome.IsPalindrome(args[0])
		fmt.Println(result)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.palindrome.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".palindrome")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
