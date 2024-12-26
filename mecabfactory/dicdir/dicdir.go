package dicdir

import (
	"bytes"
	"fmt"
	"github.com/Kuniwak/name/sliceutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type Func func() (string, error)
type DictPathFunc func() []string
type SearchPathFunc func() []string
type DictNameFunc func() []string

func FirstAvailable(dictPathFunc DictPathFunc) Func {
	return func() (string, error) {
		dictPaths := dictPathFunc()

		for _, dictPath := range dictPaths {
			stat, err := os.Stat(dictPath)
			if err != nil {
				continue
			}
			if !stat.IsDir() {
				continue
			}
			return dictPath, nil
		}

		return "", fmt.Errorf("no available dictionary path:\n%s", strings.Join(dictPaths, "\n"))
	}
}

func ByDictPath(dictPath string) DictPathFunc {
	return func() []string {
		return []string{dictPath}
	}
}

func ByDictNamesWithSearchPaths(searchPathFunc SearchPathFunc, dicNameFunc DictNameFunc) DictPathFunc {
	return func() []string {
		searchPaths := searchPathFunc()
		dicNames := dicNameFunc()
		dicPaths := make([]string, 0, len(searchPaths)*len(dicNames))
		for _, dicPath := range sliceutil.Cartesian(searchPaths, dicNames) {
			dicPaths = append(dicPaths, filepath.Join(dicPath...))
		}
		return dicPaths
	}
}

func DictNameConcat(fs ...DictNameFunc) DictNameFunc {
	return func() []string {
		result := make([]string, 0)
		for _, f := range fs {
			result = append(result, f()...)
		}
		return result
	}
}

func NeologdOrIPADicUTF8() DictNameFunc {
	return DictNameConcat(Neologd(), IPADicUTF8())
}

func Neologd() DictNameFunc {
	return func() []string {
		return []string{
			"mecab-ipadic-neologd",
		}
	}
}

func IPADicUTF8() DictNameFunc {
	return func() []string {
		return []string{
			"ipadic-utf8",
			"ipadic",
		}
	}
}

func SearchPathEmpty() SearchPathFunc {
	return func() []string {
		return []string{}
	}
}

func SearchPathConcat(fs ...SearchPathFunc) SearchPathFunc {
	return func() []string {
		result := make([]string, 0)
		for _, f := range fs {
			result = append(result, f()...)
		}
		return result
	}
}

func searchPathByMecabConfig() (string, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := exec.Command("mecab-config", "--dicdir")
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start mecab-config: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("failed to wait mecab-config: %w\n%s", err, stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}

func SearchPathByMecabConfig() SearchPathFunc {
	return func() []string {
		s, err := searchPathByMecabConfig()
		if err != nil {
			return []string{}
		}
		return []string{s}
	}
}

func SearchPathByPath(dir string) SearchPathFunc {
	return func() []string {
		stat, err := os.Stat(dir)
		if err != nil {
			return []string{}
		}
		if !stat.IsDir() {
			return []string{}
		}
		return []string{dir}
	}
}

func SearchPathByOS() SearchPathFunc {
	switch runtime.GOOS {
	case "linux":
		return SearchPathConcat(SearchPathForLinux(), SearchPathForHomebrew())
	case "darwin":
		return SearchPathForHomebrew()
	case "windows":
		return SearchPathForWindows()
	default:
		return SearchPathEmpty()
	}
}

func SearchPathForLinux() SearchPathFunc {
	return SearchPathConcat(
		SearchPathByMecabConfig(),
		SearchPathByPath(`/var/lib64/mecab/dic`),
		SearchPathByPath(`/var/lib/mecab/dic`),
	)
}

func SearchPathForHomebrew() SearchPathFunc {
	return SearchPathConcat(
		SearchPathByMecabConfig(),
		SearchPathByPath(`/opt/homebrew/lib/mecab/dic`),
	)
}

func SearchPathForWindows() SearchPathFunc {
	programFiles := os.Getenv("programfiles")
	if programFiles == "" {
		programFiles = filepath.Join(`C:`, ``, `Program Files`)
	}
	programFilesX86 := os.Getenv("programfiles(x86)")
	if programFilesX86 == "" {
		programFilesX86 = filepath.Join(`C:`, ``, `Program Files (x86)`)
	}
	programData := os.Getenv("programdata")
	if programData == "" {
		programData = filepath.Join(`C:`, ``, `ProgramData`)
	}
	return SearchPathConcat(
		SearchPathByMecabConfig(),
		SearchPathByPath(filepath.Join(`C:`, ``, `MeCab`, `dic`)),
		SearchPathByPath(filepath.Join(programData, `MeCab`, `dic`)),
		SearchPathByPath(filepath.Join(programFiles, `MeCab`, `dic`)),
		SearchPathByPath(filepath.Join(programFilesX86, `MeCab`, `dic`)),
	)
}
