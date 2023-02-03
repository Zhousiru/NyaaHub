package storage

import (
	"fmt"
	"os/exec"

	"github.com/Zhousiru/NyaaHub/fetcher/lib/config"
)

func UploadDir(localPath, remoteDir string) error {
	cmd := exec.Command(
		"rclone",
		"--config",
		config.RcloneConfigPath,
		"copy",
		localPath,
		config.RcloneDriver+":"+remoteDir,
		"--include",
		"*",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("output: %s, err: %w", string(output), err)
	}
	return nil
}
