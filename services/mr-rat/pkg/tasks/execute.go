package tasks

import (
	"os/exec"
)

func (t PingTask) Run() Result {
	return Result{
		ID:      t.ID,
		Output:  "PONG!",
		Success: true,
	}
}

func (t ExecuteTask) Run() Result {
	cmd := exec.Command("sh", "-c", t.Command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return Result{
			ID:      t.ID,
			Output:  err.Error(),
			Success: false,
		}
	}
	return Result{
		ID:      t.ID,
		Output:  string(output),
		Success: true,
	}
}
