package common_tools

import (
	"bufio"
	"fmt"
	"os/exec"
)

/*
ExecCommand
 is used in linux, need test!!
*/
func ExecCommand(command string) (err error, data *[]string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		err = fmt.Errorf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}
	if err = cmd.Start(); err != nil {
		return
	}
	outputBuf := bufio.NewReader(stdout)
	var redData []string
	for {
		output, _, errs := outputBuf.ReadLine()
		if errs != nil {
			if errs.Error() != "EOF" {
				err = errs
				return
			}
			break
		}
		redData = append(redData, string(output))
	}
	if err = cmd.Wait(); err != nil {
		return
	}
	return err, &redData
}
