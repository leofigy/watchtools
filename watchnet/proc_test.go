package watchnet

import (
	"context"
	"testing"
	"time"
)

func TestWatch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Log("running")
	Watch(ctx)
	time.Sleep(5 * time.Second)
	cancel()
}
