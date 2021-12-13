package watchnet

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	proc "github.com/mitchellh/go-ps"
)

// TODO
// run in proc sudo lsof -i -P -n
// NetProcess simple process that opened/closed a port
type NetProcess struct {
	Pid         uint
	Ports       []uint
	DisplayName string
}

func Watch(ctx context.Context) (alive chan NetProcess, dead chan NetProcess) {
	alive = make(chan NetProcess)
	dead = make(chan NetProcess)
	go monitor(ctx, alive, dead)
	return
}

func monitor(ctx context.Context, alive, dead chan NetProcess) {
	fmt.Println("inspecting running processes")

eventLoop:
	for {
		current, err := proc.Processes()

		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 1)
		}

		select {
		case <-ctx.Done():
			break eventLoop
		default:
			for _, i := range current {
				fmt.Println(i)
			}
			inspect(defaultArgs...)
		}
	}

}

func inspect(args ...string) (map[int]NetProcess, error) {
	netdata := make(map[int]NetProcess)

	output, err := exec.Command(lsofCommand, args...).Output()

	if err != nil {
		return netdata, err
	}

	items := strings.Split(string(output), NEWLINE)

	for _, line := range items {
		fmt.Println(line)
	}

	return netdata, nil
}
