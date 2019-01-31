package tifig

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func getBinPath() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if runtime.GOOS == "darwin" {
		return fmt.Sprintf("%s/bins/tifig-darwin", pwd), nil
	} else if runtime.GOOS == "linux" {
		return fmt.Sprintf("%s/bins/tifig-linux", pwd), nil
	}

	return "", fmt.Errorf("%s not support tifig", runtime.GOOS)
}

func ConvertReader(input io.Reader, output string) (string, error) {
	f, err := ioutil.TempFile("", "tifig-*.heic")
	if err != nil {
		return "", err
	}
	defer os.Remove(f.Name())

	if _, err := io.Copy(f, input); err != nil {
		return "", err
	}

	return Convert(f.Name(), output)
}

func Convert(input, output string) (string, error) {
	binPath, err := getBinPath()
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(input)
	if output == "" {
		output = strings.Replace(input, ext, ".png", -1)
	}
	stderr := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	cmd := exec.Command(binPath, "-i", input, "-o", output, "-q=100")
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err = cmd.Run(); err != nil {
		return "", errors.New(strings.TrimSpace(stderr.String()))
	}

	return output, nil
}
