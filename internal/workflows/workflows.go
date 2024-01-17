package workflows

import (
	"fmt"
)

type Tag string

const (
	Gnome        Tag = "GNOME"
	System       Tag = "System"
	Gaming       Tag = "Gaming"
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
