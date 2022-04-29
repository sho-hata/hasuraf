package hasura

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
)

const (
	CalledSeedApply     string = "seed apply"
	calledMigrateApply  string = "migrate apply"
	calledMigrateDelete string = "migrate delete"
)

var versionRegex *regexp.Regexp

type Cmd struct {
	called    string
	command   []string
	fileNames []string
	options   map[string]interface{}
	target    string
}

func NewHasuraCmd(called string, options map[string]interface{}) *Cmd {
	if called == calledMigrateApply || called == calledMigrateDelete {
		setRegex()
	}

	return &Cmd{called: called, options: options}
}

func (h *Cmd) Run() (string, error) {
	if err := h.setFileNames(); err != nil {
		return "", err
	}

	if err := h.setTarget(); err != nil {
		return "", err
	}

	return h.setCommand().exec()
}

func (h *Cmd) exec() (string, error) {
	fmt.Println("running... ", "hasura", strings.Join(h.command, " "))
	fmt.Println("")

	r, err := exec.Command("hasura", h.command...).CombinedOutput()

	return string(r), err
}

func (h *Cmd) setCommand() *Cmd {
	switch h.called {
	case CalledSeedApply:
		h.command = append(strings.Split(h.called, " "), []string{"--file", h.target}...)
	case calledMigrateApply, calledMigrateDelete:
		h.command = append(strings.Split(h.called, " "), []string{"--version", h.target}...)
	}

	if h.target == "" {
		h.command = strings.Split(h.called, " ")
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

func (h *Cmd) setTarget() error {
	if len(h.fileNames) == 0 {
		return nil
	}

	fileName, err := h.findOne()
	if err != nil {
		return err
	}

	if h.called == calledMigrateApply || h.called == calledMigrateDelete {
		h.target = trimVersion(fileName)
	} else {
		h.target = fileName
	}

	return nil
}

func (h *Cmd) setFileNames() error {
	var filePath string

	switch h.called {
	case CalledSeedApply:
		filePath = fmt.Sprintf("./seeds/%s", h.options["database-name"])
	case calledMigrateApply, calledMigrateDelete:
		filePath = fmt.Sprintf("./migrations/%s", h.options["database-name"])
	default:
		return nil
	}

	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("no file")
	}

	for _, file := range files {
		if file.IsDir() {
			if h.called == calledMigrateApply || h.called == calledMigrateDelete {
				h.fileNames = append(h.fileNames, file.Name())
			}
		}

		if !file.IsDir() && h.called == CalledSeedApply {
			h.fileNames = append(h.fileNames, file.Name())
		}
	}

	return nil
}

func (h *Cmd) findOne() (string, error) {
	type fileName struct {
		name string
	}

	fileNames := make([]fileName, 0, len(h.fileNames))

	for _, f := range h.fileNames {
		fileNames = append(fileNames, fileName{f})
	}

	i, err := fuzzyfinder.Find(fileNames, func(i int) string { return fileNames[i].name })
	if err != nil {
		return "", err
	}

	return fileNames[i].name, nil
}

func trimVersion(fileName string) string {
	return string(versionRegex.Find([]byte(fileName)))
}

func setRegex() {
	versionRegex = regexp.MustCompile(`^[0-9]+`)
}
