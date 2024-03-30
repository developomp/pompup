// Package wrapper abstracts away OS-native operations like package installation
// and writing to file as root.
package wrapper

import (
	"log"
	"os/user"
)

func GetUserName() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return user.Username
}
