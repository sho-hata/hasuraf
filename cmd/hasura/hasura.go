package hasura

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
)

type HasuraCmd struct {
	called    string
	command   []string
	fileNames []string
	options   map[string]interface{}
	target    string
}

func NewHasuraCmd(called string, options map[string]interface{}) *HasuraCmd {
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
	fmt.Println("running... ", "hasura", strings.Join(h.command, " "))
	fmt.Println("")

	r, err := exec.Command("hasura", h.command...).CombinedOutput()
	return string(r), err
}

func (h *HasuraCmd) setCommand() *HasuraCmd {
	if h.target == "" {
		return h
	}

	if h.called == "seed" {
		h.command = []string{"seed", "apply", "--file", h.target}
	} else {
		return h
	}
	// set optional flags
	for k, v := range h.options {
		h.command = append(h.command, fmt.Sprintf("--%s", k))
		switch v := v.(type) {
		case string:
			h.command = append(h.command, v)
		case bool:
			h.command = append(h.command, strconv.FormatBool(v))
		}
	}
	return h
}

func (h *HasuraCmd) setTarget() error {
	target, err := h.findOne()
	if err != nil {
		return err
	}
	h.target = target
	return nil
}

func (h *HasuraCmd) setFileNames() error {
	seedFilePath := fmt.Sprintf("./seeds/%s", h.options["database-name"])
	files, err := ioutil.ReadDir(seedFilePath)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return errors.New("no file")
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
	var fileNames []fileName //nolint:prealloc // Since filenames include directory names, they are less in length than in capacity.
	for _, f := range h.fileNames {
		fileNames = append(fileNames, fileName{f})
	}
	i, err := fuzzyfinder.Find(fileNames, func(i int) string { return fileNames[i].name })
	if err != nil {
		return "", err
	}
	return fileNames[i].name, nil
}
