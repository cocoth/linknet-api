package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/cocoth/linknet-api/src/utils"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var RootCmd = &cobra.Command{
	Use:   "linknet-api",
	Short: "Linknet API",
	Long:  "Linknet API is a Restfull API for managing users and roles etc...",
}

func Exec() {
	if err := RootCmd.Execute(); err != nil {
		utils.Error(err.Error(), "RootCmd")
		os.Exit(1)
	}
}

func PromptInput(prompt string) string {
	r := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + ": ")
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

func PromptInputCredentials(prompt string) string {
	fmt.Print(prompt + ": ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		os.Exit(1)
	}
	fmt.Println()
	return strings.TrimSpace(string(bytePassword))
}

// SaveEnv saves environment variables to a .env file
func SaveEnv(key, value string) error {
	file, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s=%s\n", key, value))
	return err
}

// UpdateEnv updates or adds a key-value pair in the .env file
func UpdateEnv(key, value string) error {
	// Read the existing .env file
	file, err := os.Open(".env")
	if err != nil {
		return SaveEnv(key, value)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			lines = append(lines, fmt.Sprintf("%s=%s", key, value))
			found = true
		} else {
			lines = append(lines, line)
		}
	}

	if !found {
		lines = append(lines, fmt.Sprintf("%s=%s", key, value))
	}

	// Write the updated .env file
	file, err = os.OpenFile(".env", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
