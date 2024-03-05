package main

import (
    "context"
)

type Fileutils struct{}

func (m *Fileutils) Tree(ctx context.Context) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithWorkdir("/mnt").
		WithExec([]string{"tree"}).
		Stdout(ctx)
}
