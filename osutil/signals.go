package osutil

import (
	"context"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"

	"github.com/maja42/mlog"
)

// SignalHandler is used to manage and react to OS signals.
type SignalHandler struct {
	Signals             []os.Signal
	MaxShutdownDuration time.Duration
	ExitFunc            func(code int)
	Logger              mlog.Logger
}

// NewSignalHandler creates a new signal manager with useful default options.
func NewSignalHandler(logger mlog.Logger) *SignalHandler {
	return &SignalHandler{
		Signals:             []os.Signal{syscall.SIGINT, syscall.SIGTERM},
		MaxShutdownDuration: 10 * time.Second,
		ExitFunc:            os.Exit,
		Logger:              logger.New("sig-handler"),
	}
}

// HandleShutdown waits for an OS signal and then calls the given shutdown function.
// If the context is not cancelled within the MaxShutdownDuration, the application is terminated (ExitFunc).
// Aborts if the context is cancelled.
func (h *SignalHandler) HandleShutdown(ctx context.Context, shutdownFunc func()) {
	sigchan := make(chan os.Signal, 2)
	signal.Notify(sigchan, h.Signals...)

	go func() {
		defer signal.Stop(sigchan)

		select {
		case <-ctx.Done():
			return
		case sig := <-sigchan:
			h.Logger.Warnf("Received signal %q - shutting down.", sig.String())
			shutdownFunc()
		}

		select {
		case <-ctx.Done():
			return
		case sig := <-sigchan:
			h.Logger.Warnf("Received signal %q (second attempt) - terminating.", sig.String())
			h.ExitFunc(2)

		case <-time.After(h.MaxShutdownDuration):
			// application is still running --> log goroutine dump and terminate
			errWriter := h.Logger.WriterLevel(mlog.ErrorLevel)
			_ = pprof.Lookup("goroutine").WriteTo(errWriter, 1)
			_ = errWriter.Close()

			h.Logger.Errorf("Application did not shutdown after %s - terminating.", h.MaxShutdownDuration)
			h.ExitFunc(1)
		}
	}()
}
