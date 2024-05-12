package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"

	"GoEdu/internal/pkg/config"
	"GoEdu/internal/pkg/module"
	"GoEdu/internal/pkg/waiter"
)

type monolith struct {
	cfg          *config.Config
	dependencies module.Dependencies
	modules      []module.Module
	waiter       waiter.Waiter
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("server exitted abnormally: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}
	d, err := module.BuildDefaultDependencies(cfg)
	if err != nil {
		return err
	}
	m := monolith{
		cfg:          cfg,
		dependencies: d,
		modules:      []module.Module{},
		waiter:       waiter.New(waiter.CatchSignals()),
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	m.waiter.Add(
		m.ServeHTTP,
	)

	return m.waiter.Wait()
}

func (m *monolith) startupModules() error {
	for _, mod := range m.modules {
		ctx := m.waiter.Context()
		if err := mod.Init(ctx, m.dependencies); err != nil {
			return err
		}
	}

	return nil
}

func (m *monolith) ServeHTTP(ctx context.Context) error {
	l := m.dependencies.Logger()
	srv := &http.Server{
		Addr:    m.cfg.HTTP.Address(),
		Handler: m.dependencies.HTTP(),
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		l.InfoContext(ctx, fmt.Sprintf("HTTP server started; listening at http://%s", m.cfg.HTTP.Address()))
		defer fmt.Println("web server shutdown")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		l.InfoContext(ctx, "HTTP to be shutdown")
		ctx, cancel := context.WithTimeout(context.Background(), m.cfg.ShutdownTimeout)
		defer cancel()
		return srv.Shutdown(ctx)
	})

	return group.Wait()
}
