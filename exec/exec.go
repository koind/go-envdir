package exec

import (
	"fmt"
	"os"
	"os/exec"
)

func Command(commandName, envDir string) {
	cmd := exec.Command(commandName, envDir)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err := cmd.CombinedOutput()
	os.Stdout.Write(out)
	os.Stderr.Write([]byte(err.Error()))
}
