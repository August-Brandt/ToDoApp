// A generated module for TodoApp functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"math/rand/v2"
	"dagger/todo-app/internal/dagger"
	"fmt"
	"math"
)

type TodoApp struct{}

func (m *TodoApp) Publish(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	client := m.BuildClient(source.Directory("client"))
	server := m.BuildServer(source.Directory("server/src"))

	return dag.Container().
		From("ubuntu:latest").
		WithDirectory("/app/client/dist/", client).
		WithFile("/app/server/src/server", server).
		WithExec([]string{"cp", "server/database/ToDoDatabase.db", "/app/server/database/ToDoDatabase.db"}).
		WithExposedPort(8080).
		WithWorkdir("/app/server/src").
		WithEntrypoint([]string{"./server"}).
		Publish(ctx, fmt.Sprintf("ttl.sh/todo-app-%.0f", math.Floor(rand.Float64()*10000000)))
}

func (m *TodoApp) BuildClient(source *dagger.Directory) *dagger.Directory {
	return m.ClientBuildEnv(source).
		WithExec([]string{"npm", "run", "build"}).
		Directory("./dist")
}

func (m *TodoApp) ClientBuildEnv(source *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("node:24").
		WithDirectory("/client", source).
		WithWorkdir("/client").
		WithExec([]string{"npm", "install"})
}

func (m *TodoApp) BuildServer(source *dagger.Directory) *dagger.File {
	return m.ServerBuildEnv(source).
		WithExec([]string{"go", "build", "-o", "server", "."}).
		File("./server")
}

func (m *TodoApp) ServerBuildEnv(source *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang:1.22").
		WithDirectory("/server/src", source).
		WithWorkdir("/server/src").
		WithExec([]string{"go", "mod", "tidy"})
}
