package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

import (
	"bytes"
	"errors"
	"os/exec"
)

type Svn struct{}

func runCommand(commandArgs []string) error {
	var arguments []string
	var stderr bytes.Buffer
	executable := commandArgs[0]
	if len(commandArgs) > 1 {
		arguments = commandArgs[1:]
	} else {
		arguments = []string{}
	}
	cmd := exec.Command(executable, arguments...)
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		if stderr.Len() > 0 {
			return errors.New(stderr.String())
		}
		return err
	}
	return nil
}

func commandResult(commandArgs []string) ([]byte, error) {
	var arguments []string
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	executable := commandArgs[0]
	if len(commandArgs) > 1 {
		arguments = commandArgs[1:]
	} else {
		arguments = []string{}
	}
	cmd := exec.Command(executable, arguments...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		if stderr.Len() > 0 {
			return nil, errors.New(stderr.String())
		}
		return nil, err
	}
	return stdout.Bytes(), nil
}
