package hasura

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os/exec"

	"github.com/ktr0731/go-fuzzyfinder"
)

type HasuraCmd struct {
	called    string
	command   []string
	fileNames []string
	options   map[string]string
	target    string
}

func NewHasuraCmd(called string, options map[string]string) *HasuraCmd {
	return &HasuraCmd{called: called, options: options}
}

func (h *HasuraCmd) Run() (string, error) {
	if err := h.setFileNames(); err != nil {
		return "", err
	}
	if err := h.setTarget(); err != nil {
		return "", err
	}
	return h.setCommand().exec()
}

func (h *HasuraCmd) exec() (string, error) {
	fmt.Println(h.command)
	r, err := exec.Command("hasura", h.command...).CombinedOutput()
	return string(r), err
}

func (h *HasuraCmd) setCommand() *HasuraCmd {
	if h.called == "seed" {
		h.command = []string{"seed", "apply", "--file", h.target}
	}
	// set flags
	for k, v := range h.options {
		h.command = append(h.command, fmt.Sprintf("--%s", k))
		h.command = append(h.command, v)
	}
	return h
}

func (h *HasuraCmd) setTarget() error {
	target, err := h.findOne()
	if err != nil {
		return err
	}
	h.target = fmt.Sprintf("seed/%s/%s", h.options["database-name"], target)
	return nil
}

func (h *HasuraCmd) setFileNames() error {
	const seedFilePath string = "./seeds/default"
	files, err := readFiles(seedFilePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			h.fileNames = append(h.fileNames, file.Name())
		}
	}
	return nil
}

func (h *HasuraCmd) findOne() (string, error) {
	type fileName struct {
		name string
	}
	var fileNames []fileName
	for _, f := range h.fileNames {
		fileNames = append(fileNames, fileName{f})
	}
	i, err := fuzzyfinder.Find(fileNames, func(i int) string { return fileNames[i].name })
	if err != nil {
		return "", err
	}
	return fileNames[i].name, nil
}

func readFiles(path string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, errors.New("no file")
	}
	return files, nil
}
