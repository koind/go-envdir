package exec

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Command(commandName, envDir string) {
	envList, err := getEnv(envDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = setEnv(envList)
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd := exec.Command(commandName)
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

func getEnv(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	envList := make(map[string]string, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		texts := strings.Split(scanner.Text(), "=")
		if len(texts) >= 2 {
			envList[texts[0]] = texts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envList, nil
}

func setEnv(evnList map[string]string) error {
	for key, val := range evnList {
		err := os.Setenv(key, val)
		if err != nil {
			return err
		}
	}

	return nil
}
