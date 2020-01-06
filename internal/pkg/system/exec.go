package system

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)


func CheckError(output[]byte, err error) {
	if err  != nil {
		fmt.Print(string(output))
		fmt.Println(err)
		os.Exit(1)
	}
}


func ShellExec(dir string, cmd string, args ...string) (error) {

	ctx := context.Background()
	shell := exec.CommandContext(ctx,cmd,args...)
	shell.Dir = dir
	shell.Stdout = os.Stdout
	shell.Stderr = os.Stderr

	err := shell.Run()
	ctx.Done()

	return err
}