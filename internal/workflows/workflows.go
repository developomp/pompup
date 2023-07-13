package workflows

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/pterm/pterm"
)

// DefaultPerm is equivalent to -rw-r--r--
const DefaultPerm = 0644

type Tag string

const (
	Gnome        Tag = "GNOME"
	System       Tag = "System"
	Dev          Tag = "Dev"
	Cli          Tag = "CLI"
	Gui          Tag = "GUI"
	Configurator Tag = "Configurator"
)

// Workflow can be many things. It could install and configure an application,
// modify the operating system (both kernel and user space), etc etc...
// Think of it as a loose term for installer.
type Workflow struct {
	Name      string   // Name is the display name of the workflow.
	Desc      string   // Desc briefly explains what the workflow does. Usually in one sentence.
	Tags      []Tag    // Tags are used to categorize different workflows
	Setup     func()   // Setup contains logic regarding the setup process.
	Reminders []string // Reminders for manual tasks user has to perform. Shows after all Setup functions are executed.
}

// list of workflows
var Workflows []*Workflow

func register(workflow *Workflow) {
	for i, reminder := range workflow.Reminders {
		workflow.Reminders[i] = fmt.Sprintf("[%s] %s", workflow.Name, reminder)
	}

	Workflows = append(Workflows, workflow)
}

func getHomeDir() (homeDir string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		pterm.Fatal.Println("Failed to get user's home directory: ", err)
	}

	return
}

func writeFile(path string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, DefaultPerm)
}

func inHome(path string) string {
	return fmt.Sprintf("%s/%s", getHomeDir(), path)
}

func getUserName() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return user.Username
}
