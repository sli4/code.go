package adapter

import "testing"

func TestAdapter(t *testing.T) {
	cpu := &M1CPU{}
	adapter := &M1Adapter{
		cpu: cpu,
	}
	adapter.Run()
}
