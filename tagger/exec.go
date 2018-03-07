package tagger

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ExecDirectory(commandStr string, directory string) (err error) {
	if err := os.Chdir(directory); err != nil {
		log.Errorf("Failed to cd %s: %s", directory, err)
		return err
	}

	return Exec(commandStr)
}

func Exec(commandStr string) (err error) {
	if strings.TrimSpace(commandStr) == "" {
		return errors.New("No command provided")
	}

	var name string
	var args []string

	cmdArr := strings.Split(commandStr, " ")
	name = cmdArr[0]

	if len(cmdArr) > 1 {
		args = cmdArr[1:]
	}

	command := exec.Command(name, args...)
	out, err := command.CombinedOutput()
	printOutput := log.GetLevel() == log.DebugLevel
	if printOutput {
		fmt.Println(string(out))
	}
	if err != nil {
		fmt.Println("Error on combined output:", err)
	}

	return
}
