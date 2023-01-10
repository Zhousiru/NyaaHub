package storage

import (
	"fmt"
	"os/exec"

	"github.com/Zhousiru/NyaaHub/fetcher/lib/config"
)

func Upload(localPath, remoteDir string) error {
	cmd := exec.Command(
		"rclone",
		"--config",
		config.RcloneConfigPath,
		"copy",
		localPath,
		config.RcloneDriver+":"+remoteDir,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("output: %s, err: %w", string(output), err)
	}
	return nil
}
