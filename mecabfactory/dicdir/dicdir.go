package dicdir

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type Func func() (string, error)

func Fallback(f1, f2 Func) Func {
	return func() (string, error) {
		d, err := f1()
		if err == nil {
			return d, nil
		}

		return f2()
	}
}

func ByConstant(dicDir string, err error) Func {
	return func() (string, error) {
		return dicDir, err
	}
}

func Ipa(searchBaseDir Func) Func {
	return func() (string, error) {
		dirDict, err := searchBaseDir()
		if err != nil {
			return "", err
		}

		return filepath.Join(dirDict, "ipadic"), nil
	}
}

func Neologd(searchBaseDir Func) Func {
	return func() (string, error) {
		dirDict, err := searchBaseDir()
		if err != nil {
			return "", err
		}

		return filepath.Join(dirDict, "mecab-ipadic-neologd"), nil
	}
}

func ByMecabConfig() Func {
	return func() (string, error) {
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}
		cmd := exec.Command("mecab-config", "--dicdir")
		cmd.Stdout = stdout
		cmd.Stderr = stderr

		if err := cmd.Start(); err != nil {
			return "", fmt.Errorf("starting mecab-config failed: %w\n%s", err, stderr.String())
		}

		if err := cmd.Wait(); err != nil {
			return "", fmt.Errorf("waiting mecab-config failed: %w\n%s", err, stderr.String())
		}

		return strings.TrimSpace(stdout.String()), nil
	}
}
