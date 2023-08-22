package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var file, dir, format *string
	rootCmd := &cobra.Command{
		Use:   "Convertica",
		Short: "Convertica is a file format converter",
	}

	convCmd := &cobra.Command{
		Use:   "conv",
		Short: "Convert file format",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("flags:\n\t%v\n\t%v\n\t%v\n", *file, *dir, *format)
		},
	}

	file = convCmd.Flags().StringP("file", "c", "", "The directory of the file to be compressed")
	dir = convCmd.Flags().StringP("dir", "o", "", "The directory of the converted file")
	format = convCmd.Flags().StringP("format", "f", "", "Format of the new file")

	rootCmd.AddCommand(convCmd)

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

func readContent(cmd *cobra.Command, args []string) (string, error) {
	filePath, _ := cmd.Flags().GetString("file")
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func saveContentToDirectory(cmd *cobra.Command, content string, newName string) {
	outDir, _ := cmd.Flags().GetString("dir")
	if strings.HasSuffix(outDir, "/") {
		newFilePath := filepath.Join(outDir, newName)
		file, err := os.Create(newFilePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing content to file:", err)
			return
		}
		fmt.Println("File saved at:", newFilePath)
	} else {
		fmt.Println("Invalid directory")
	}
}
