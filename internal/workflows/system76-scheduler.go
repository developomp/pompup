package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Workflow{
		Name:  "System76 Scheduler",
		Desc:  "Improves responsiveness",
		Tags:  []Tag{System},
		Setup: setupSystem76Scheduler,
	})
}

func setupSystem76Scheduler() {
	wrapper.Paru("system76-scheduler")
	wrapper.Run("sudo", "systemctl", "enable", "--now", "com.system76.Scheduler.service")
}
