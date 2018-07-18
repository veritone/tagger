// Copyright Veritone Corporation 2018. All rights reserved.
// See LICENSE for more information.

package tagger

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func ExecDirectory(commandStr string, colorFunc func(string, ...interface{}), directory string) (err error) {
	if err := os.Chdir(directory); err != nil {
		log.Errorf("Failed to cd %s: %s", directory, err)
		return err
	}

	return Exec(commandStr, colorFunc)
}

func Exec(commandStr string, colorFunc func(string, ...interface{})) (err error) {
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

	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Error("Failed creating command stdoutpipe: ", err)
		return err
	}
	defer stdout.Close()
	stdoutReader := bufio.NewReader(stdout)

	stderr, err := command.StderrPipe()
	if err != nil {
		log.Error("Failed creating command stderrpipe: ", err)
		return err
	}
	defer stderr.Close()
	stderrReader := bufio.NewReader(stderr)

	if err := command.Start(); err != nil {
		log.Error("Failed starting command: ", err)
		return err
	}

	go handleReader(stdoutReader, false, colorFunc)
	go handleReader(stderrReader, true, colorFunc)

	if err := command.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Debug("Exit Status: ", status.ExitStatus())
				return err
			}
		}
		log.Debug("Failed to wait for command: ", err)
		return err
	}

	return
}

func handleReader(reader *bufio.Reader, isStderr bool, colorFunc func(string, ...interface{})) {
	printOutput := log.GetLevel() == log.DebugLevel
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if printOutput {
			colorFunc(str)
		}
	}
}
