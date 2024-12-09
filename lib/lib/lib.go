package lib

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func Bash(cmd string) error {
	c := exec.Command("/bin/bash", "-c", cmd)
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	return c.Run()
}
func BashV(cmd string) error {
	log.Println("cmd:", cmd)
	return Bash(cmd)
}

func BashOutput(cmd string) (string, error) {
	out, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput()
	return string(out), err
}

func ReadFile(path string) (string, error) {
	ret, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}
func WriteFile(path, data string) error {
	return os.WriteFile(path, []byte(data), 0644)
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	err = out.Sync()
	if err != nil {
		return err
	}

	err = out.Close()

	return err
}

func Plus(in bool) string {
	if in {
		return "+"
	}
	return "-"
}
func BoolToInt(in bool) int {
	if in {
		return 1
	}
	return 0
}
