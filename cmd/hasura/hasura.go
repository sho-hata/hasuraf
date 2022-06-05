package hasura

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

type HasuraCmd struct {
	called      string
	command     []string
	options     map[string]interface{}
	files       []fileInfo
	applyTarget string
}

type fileInfo struct {
	name     string
	headline string
}

func NewHasuraCmd(called string, options map[string]interface{}) *HasuraCmd {
	if called == calledMigrateApply || called == calledMigrateDelete {
		setRegex()
	}

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
	switch h.called {
	case CalledSeedApply:
		h.command = append(strings.Split(h.called, " "), []string{"--file", h.applyTarget}...)
	case calledMigrateApply, calledMigrateDelete:
		h.command = append(strings.Split(h.called, " "), []string{"--version", h.applyTarget}...)
	}
	if h.applyTarget == "" {
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

func (h *HasuraCmd) setTarget() error {
	if len(h.files) == 0 {
		return nil
	}

	fileName, err := h.findOne()
	if err != nil {
		return err
	}

	if h.called == calledMigrateApply || h.called == calledMigrateDelete {
		h.applyTarget = trimVersion(fileName)
	} else {
		h.applyTarget = fileName
	}

	return nil
}

func (h *HasuraCmd) setFileNames() error {
	var filePath string

	switch h.called {
	case CalledSeedApply:
		filePath = fmt.Sprintf("./seeds/%s", h.options["database-name"])
	case calledMigrateApply, calledMigrateDelete:
		filePath = fmt.Sprintf("./migrations/%s", h.options["database-name"])
	default:
		return nil
	}

	files, err := os.ReadDir(filePath)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("no such file or directory")
	}

	for _, file := range files {
		if file.IsDir() {
			if h.called == calledMigrateApply || h.called == calledMigrateDelete {
				headline, err := readFileHeadline("./migrations/default/" + file.Name() + "/up.sql")
				if err != nil {
					return err
				}
				h.files = append(h.files, fileInfo{name: file.Name(), headline: headline})
			}
		}

		if !file.IsDir() && h.called == CalledSeedApply {
			headline, err := readFileHeadline("./seeds/default/" + file.Name())
			if err != nil {
				return err
			}
			h.files = append(h.files, fileInfo{name: file.Name(), headline: headline})
		}
	}

	return nil
}

func (h *HasuraCmd) findOne() (string, error) {
	i, err := fuzzyfinder.Find(
		h.files,
		func(i int) string { return h.files[i].name },
		fuzzyfinder.WithPreviewWindow(func(i, width, height int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf(`📝 Applied SQL file

			%s
			...
			`, h.files[i].headline)
		}),
	)
	if err != nil {
		return "", err
	}
	return h.files[i].name, nil
}

func trimVersion(fileName string) string {
	return string(versionRegex.Find([]byte(fileName)))
}

func setRegex() {
	versionRegex = regexp.MustCompile(`^[0-9]+`)
}

// readFileHeadline
// Opens a file in the path passed as an argument, reads the first three lines, and returns them as a string.
func readFileHeadline(path string) (string, error) {
	var headline string
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var count int
	for scanner.Scan() {
		if count == 4 {
			headline += scanner.Text()
			break
		}
		headline += fmt.Sprintln(scanner.Text())
		count++
	}
	return headline, nil
}
