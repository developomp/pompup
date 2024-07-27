package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "System76 Scheduler",
		Desc: "Improves responsiveness",
		Setup: func() {
			wrapper.ParuOnce("system76-scheduler")
			wrapper.SystemctlEnable("com.system76.Scheduler", wrapper.EnableNow)
		},
	})
}
