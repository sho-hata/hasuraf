package command

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/ktr0731/go-fuzzyfinder"
)

type file struct {
	Name string
}

const seedFilePath string = "seeds/default/"

func Run() (string, error) {
	fileNames, err := fileNames()
	if err != nil {
		return "", err
	}
	fs := []file{}
	for _, f := range fileNames {
		fs = append(fs, file{f})
	}
	idx, err := fuzzyfinder.Find(fs, func(i int) string { return fs[i].Name })
	if err != nil {
		return "", err
	}
	// TODO: Get the optional arguments.
	return run(fs[idx].Name)
}

func run(target string) (string, error) {
	r, err := exec.Command("hasura", "seed", "apply", "--file", fmt.Sprintf("%s%s", seedFilePath, target)).CombinedOutput()
	return string(r), err
}

func fileNames() ([]string, error) {
	files, err := ioutil.ReadDir(fmt.Sprintf("./%s", seedFilePath))
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, errors.New("no seed file")
	}

	var fileNames []string
	for _, f := range files {
		if !f.IsDir() {
			fileNames = append(fileNames, f.Name())
		}
	}
	return fileNames, nil
}
