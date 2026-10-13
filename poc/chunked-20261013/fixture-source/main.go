// A harmless native health fixture, not an embedding runtime.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var c struct {
		Version string `json:"version"`
		Healthy bool   `json:"healthy"`
	}
	b, e := os.ReadFile("config.json")
	if e != nil || json.Unmarshal(b, &c) != nil || !c.Healthy {
		os.Exit(2)
	}
	f, e := os.Open("model.bin")
	if e != nil {
		os.Exit(3)
	}
	defer f.Close()
	info, e := f.Stat()
	if e != nil || info.Size() == 0 {
		os.Exit(4)
	}
	buf := make([]byte, 32)
	if _, e = f.ReadAt(buf, info.Size()-32); e != nil {
		os.Exit(5)
	}
	fmt.Printf("{\"version\":%q,\"model_bytes\":%d}\n", c.Version, info.Size())
}
