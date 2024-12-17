package cast

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

func RunCast(arg ...string) ([]byte, error) {
	cmd := exec.CommandContext(context.Background(), "cast", arg...)

	logDone := make(chan struct{})

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get handle on stdout: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start anvil: %w", err)
	}
	log := new(bytes.Buffer)

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			txt := scanner.Text()
			if _, err := fmt.Fprintln(log, txt); err != nil {
				return
			}
		}
		logDone <- struct{}{}
	}()

	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	<-logDone

	fmt.Print(string(log.Bytes()))

	return log.Bytes(), nil
}
