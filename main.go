package main

import (
	"fmt"
	"os"
	"strings"

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
		},
	}

	var outCmd = &cobra.Command{
		Use:   "outDir",
		Short: "The directory of the converted file",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	var formatCmd = &cobra.Command{
		Use:   "Format",
		Short: "Format of the new file",
		Run: func(cmd *cobra.Command, args []string) {
			formatType, _ := cmd.Flags().GetString("format")
			if formatType == "" {
				fmt.Println("Please provide a format using the -f flag")
				return
			}
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

func converter(cmd *cobra.Command, args []string) string {
	lastFileName := getFileName(cmd, args)
	fileFormat, _ := cmd.Flags().GetString("format")

	newName := lastFileName + fileFormat
	return newName
}

func isValidDirectory(fileDir string) bool {

	fileInfo, err := os.Stat(fileDir)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func getFileName(cmd *cobra.Command, args []string) string {

	fileName, _ := cmd.Flags().GetString("file")

	if isValidDirectory(fileName) == true {

		parts := strings.Split(fileName, ".")
		if len(parts) > 1 {
			baseName := strings.Join(parts[:len(parts)-1], ".")
			lastSlashIndex := strings.LastIndex(baseName, "/")
			if lastSlashIndex != -1 {
				return baseName[lastSlashIndex+1:]
			}
			return baseName
		}
	}
	return ""
}
