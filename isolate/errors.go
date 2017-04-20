package isolate

import (
	"errors"
	"syscall"
)

const (
	systemCategory     = 1
	isolateErrCategory = 42
)

const (
	codeSpawnEAGAIN = int(syscall.EAGAIN) + iota
	codeBadMsg
	codeBadProfile
	codeUnknownIsolate
	codeSpoolingFailed
	codeSpawningFailed
	codeMetricsFailed
	codeOutputError
	codeKillError
	codeSpoolCancellationError
)

var (
	errBadMsg                 = [2]int{isolateErrCategory, codeBadMsg}
	errBadProfile             = [2]int{isolateErrCategory, codeBadProfile}
	errUnknownIsolate         = [2]int{isolateErrCategory, codeUnknownIsolate}
	errSpoolingFailed         = [2]int{isolateErrCategory, codeSpoolingFailed}
	errSpawningFailed         = [2]int{isolateErrCategory, codeSpawningFailed}
	errMetricsError           = [2]int{isolateErrCategory, codeMetricsFailed}
	errOutputError            = [2]int{isolateErrCategory, codeOutputError}
	errKillError              = [2]int{isolateErrCategory, codeKillError}
	errSpoolCancellationError = [2]int{isolateErrCategory, codeSpoolCancellationError}
	errSpawnEAGAIN            = [2]int{systemCategory, codeSpawnEAGAIN}
)

var (
	ErrSpawningCancelled = errors.New("spawning has been cancelled")
)
