package helpers

import (
	"github.com/kardianos/osext"
)

// GetExecutablePath return root path for application
func GetExecutablePath() (string, error) {
	var executablePath string
	executablePath, err := osext.ExecutableFolder()
	if err != nil {
		return executablePath, err
	}
	return executablePath, nil
}
