package main

import (
    "context"
)

type Fileutils struct{}

func (m *Fileutils) Tree(ctx context.Context, dir *Directory) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", dir).
		WithWorkdir("/mnt").
		WithExec([]string{"tree"}).
		Stdout(ctx)
}
