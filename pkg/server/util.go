package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
)

func Sleep(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	select {
	case <-ctx.Done():
		timer.Stop()
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func StartStopServer(ctx context.Context, srv *http.Server, shutdownTimeout time.Duration) error {
	return StartStopTLSServer(ctx, srv, shutdownTimeout, "", "")
}

func StartStopTLSServer(ctx context.Context, srv *http.Server, shutdownTimeout time.Duration, certFile, keyFile string) error {
	var wg sync.WaitGroup
	defer wg.Wait() // wait for goroutine to shutdown active connections
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		c, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if srv.Shutdown(c) != nil {
			srv.Close()
		}
	}()

	var err error
	if certFile == "" || keyFile == "" {
		err = srv.ListenAndServe()
	} else {
		err = srv.ListenAndServeTLS(certFile, keyFile)
	}

	if err != http.ErrServerClosed {
		// Failed to start or dirty shutdown
		return errors.WithStack(err)
	}
	// Clean shutdown
	return nil
}
