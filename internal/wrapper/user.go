package wrapper

import (
	"log"
	"os/exec"
	"os/user"
	"slices"
	"strings"

	"github.com/pterm/pterm"
)

func GetUserName() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return user.Username
}

// AddGroup adds the user to group if the user is not part of the group already.
func AddGroup(group string) {
	out, err := exec.Command("groups").Output()
	if err != nil {
		pterm.Fatal.Println("Failed to add user to", group, "group:", err)
	}

	groups := strings.Fields(string(out))
	if !slices.Contains(groups, group) {
		err = Run("sudo", "usermod", "-aG", group, GetUserName()) // add current user to group
		if err != nil {
			pterm.Fatal.Println("Failed to add user to", group, "group:", err)
		}
	}
}
