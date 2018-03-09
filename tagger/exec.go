package tagger

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

var (
	mut      = &sync.Mutex{}
	colorInd = 0
	colorMap = map[int]func(string, ...interface{}){
		0: color.Blue,
		1: color.Cyan,
		2: color.Green,
		3: color.Magenta,
		4: color.Red,
		5: color.White,
		6: color.Yellow,
	}
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

	colorFunc := getColorFunc()

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

func getColorFunc() (colorFunc func(string, ...interface{})) {
	mut.Lock()
	colorFunc = colorMap[colorInd%len(colorMap)]
	colorInd++
	mut.Unlock()
	return
}
