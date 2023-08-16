package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "Convertica",
		Short: "Convertica is a file format converter",
	}

	var convCmd = &cobra.Command{
		Use:   "conv",
		Short: "Convert file format",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fileDir, _ := cmd.Flags().GetString("file")
			if fileDir == "" {
				fmt.Println("Please provide a file directory using the --file flag")
				return

			}
		},
	}

	var outCmd = &cobra.Command{
		Use:   "outDir",
		Short: "The directory of the converted file",
		Run: func(cmd *cobra.Command, args []string) {
			//çıkış dizinini aldığın fonksiyonu ekle
		},
	}

	var formatCmd = &cobra.Command{
		Use:   "Format",
		Short: "Format of the new file",
		Run: func(cmd *cobra.Command, args []string) {
			//dosya formatını aldığın func
		},
	}

	convCmd.Flags().StringP("file", "c", "", "The directory of the file to be compressed")
	outCmd.Flags().StringP("dir", "o", "", "The directory of the converted file")
	formatCmd.Flags().StringP("format", "f", "", "Format of the new file")

	rootCmd.AddCommand(convCmd)
	rootCmd.AddCommand(outCmd)
	rootCmd.AddCommand(formatCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
