package app

import (
	"WAD-2022/internal/app/repository"
	"context"
)

type Application struct {
	ctx  context.Context
	repo *repository.Repository
}

func (app *Application) Run(ctx context.Context) error {
	app.StartServer()
	return nil
}

func New(ctx context.Context) (*Application, error) {
	app := &Application{
		ctx: ctx,
	}
	repo, err := repository.New(ctx)
	if err != nil {
		return nil, err
	}
	app.repo = repo
	return app, nil
}
