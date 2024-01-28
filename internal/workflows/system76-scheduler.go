package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
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
	install.Paru("system76-scheduler")
	helper.Run("sudo", "systemctl", "enable", "--now", "com.system76.Scheduler.service")
}
