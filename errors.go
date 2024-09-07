package neolauncher

import "errors"

// ErrLauncherIsNil for when nil is passed to the launcher
var ErrLauncherIsNil = errors.New("launcher implementation is nil")
