/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/richbai90/img_to_pix/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "img_to_pix path/to/img path/to/output",
	Short: "convert image file to pixel buffer",
	Long:  `Convert an image file to a pixel buffer`,
	Args:  cobra.MinimumNArgs(2),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		img := utils.DecodeImg(args[0])
		bytes := utils.GetBytes(img)
		f, err := os.Create(args[1])
		if err != nil {
			log.Fatal("Failed to create outfile ", args[2], "\nError: ", err.Error())
		}

		for i, b := range bytes {
			if (i+1)%4 == 0 {
				fmt.Fprintln(f)
				continue
			}
			fmt.Fprintf(f, "%02X", b)
		}
		if rpt, _ := cmd.Flags().GetString("report"); rpt != "" {
			r, err := os.Create(rpt)
			if err != nil {
				log.Fatal("Could not create report ", rpt, "\nError: ", err.Error())
			}

			fmt.Fprintf(r, `Filename: %s
Dimensions: %dx%d
Total Pixels: %d
Pixels Per Row: %d
bytes: %d
Bytes Per Row: %d`,
				args[0], img.Bounds().Dx(), img.Bounds().Dy(), len(bytes)/4, len(bytes)/4/img.Bounds().Dy(), len(bytes)/4*3, (len(bytes)/4*3)/img.Bounds().Dy())
		}

		defer f.Close()
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.img_to_pix.yaml)")
	rootCmd.Flags().StringP("report", "r", "", "Generate a report with information about the image file and format including dimensions")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
