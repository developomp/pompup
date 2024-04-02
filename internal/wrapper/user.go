package wrapper

import (
	"log"
	"os/exec"
	"os/user"
	"slices"
	"strings"
)

func GetUserName() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return user.Username
}

// AddGroup adds the user to group if the user is not part of the group already.
func AddGroup(group string) error {
	out, err := exec.Command("groups").Output()
	if err != nil {
		return err
	}

	groups := strings.Fields(string(out))
	if !slices.Contains(groups, group) {
		username := GetUserName()
		err = Run("sudo", "usermod", "-aG", group, username) // add current user to group
		if err != nil {
			return err
		}
	}

	return nil
}
