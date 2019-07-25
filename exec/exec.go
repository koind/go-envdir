package exec

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func Command(commandName, envDir string) {
	envList, err := getEnv(envDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd := exec.Command(commandName)
	cmd.Env = append(cmd.Env, envList...)
	err = cmd.Start()
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

func getEnv(filePath string) ([]string, error) {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	envList := make([]string, 0, 10)
	texts := trimSpaces(string(body))
	slice := strings.Split(texts, " ")

	for _, text := range slice {
		envList = append(envList, strings.Split(text, "=")...)
	}

	return envList, nil
}

func trimSpaces(text string) string {
	space := regexp.MustCompile(`\s+`)

	return space.ReplaceAllString(text, " ")
}
