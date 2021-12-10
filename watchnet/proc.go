package watchnet

// NetProcess simple process that opened/closed a port
type NetProcess struct {
     Pid uint
     Ports []uint
     DisplayName string
}

func Watch() {}

func getClosed(pipe chan NetProcess) {}

func getOpenned(pipe chan NetProcess) {}
