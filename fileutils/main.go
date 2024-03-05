package main

import (
    "context"
)

type FileUtils struct{}

func (f *FileUtils) Ls(ctx context.Context, dir *Directory) (string ,error) {
    return dag.Container().
        From("alpine:latest").
		WithMountedDirectory("/mnt", dir).
		WithWorkdir("/mnt").
		WithExec([]string{"ls -la"}).
		Stdout(ctx)
}
