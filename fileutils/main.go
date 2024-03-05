package main

import (
    "context"
)

type Fileutils struct{}

func (f *Fileutils) Ls(ctx context.Context, dir *Directory) (string ,error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", dir).
		WithWorkdir("/mnt").
		WithExec([]string{"ls"}).
		Stdout(ctx)
}
