package cronsun

import (
	"fmt"
	"runtime"
)

const Binary = "v0.1.1"

var (
	Version = fmt.Sprintf("%s (build %s)", Binary, runtime.Version())
)
