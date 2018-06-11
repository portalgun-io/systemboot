package recovery

import (
	"log"
	"os/exec"
)

const shell string = "rush"

// PermissiveRecoverer is a sad
type PermissiveRecoverer struct {
	Debug bool
}

// Recover sasd
func (pr PermissiveRecoverer) Recover(message string) error {
	if message != "" {
		log.Panicf("%s\n", message)
	}

	if pr.Debug {
		path, err := exec.LookPath(shell)
		if err != nil {
			return err
		}

		cmd := exec.Command(path)
		err = cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
