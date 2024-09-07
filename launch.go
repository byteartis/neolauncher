package neolauncher

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

var exit = os.Exit

type Launcher interface {
	Launch(context.Context) error
}

// Launch takes in a context and a Launcher implementation.
// The launcher implementation is expected to be a long running process that will be terminated
// once either SIGTERM or SIGINT is received.
func Launch(ctx context.Context, svc Launcher) {
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGTERM, syscall.SIGINT)

	logger := log.Ctx(ctx)
	err := launch(logger.WithContext(context.Background()), quitCh, svc)
	if err != nil {
		logger.Error().Err(err).Send()
		exit(1)
	}

	exit(0)
}

func launch(ctx context.Context, quitCh chan os.Signal, svc Launcher) error {
	if svc == nil {
		return ErrLauncherIsNil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		<-quitCh
		cancel()
	}()

	return svc.Launch(ctx)
}
