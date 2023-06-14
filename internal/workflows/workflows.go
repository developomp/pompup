package workflows

// Workflow can be many things. It could install and configure an application,
// modify the operating system (both kernel and user space), etc etc...
// Think of it as a loose term for installer.
type Workflow struct {
	Name      string // Name is the display name of the workflow.
	Desc      string // Desc briefly explains what the workflow does. Usually in one sentence.
	Setup     func() // Setup contains logic regarding the setup process.
	PostSetup func() // PostSetup contains code that will run after ALL selected Setup functions are executed.
}

// list of workflows
var Workflows []*Workflow

func register(workflow *Workflow) {
	Workflows = append(Workflows, workflow)
}
