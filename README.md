# Short URL
Go short url service.

## Prerequisites

### Container Engine
This project uses [Podman](https://podman.io) as the container engine, but OCI compatible container engines should work. 

### Taskfile

[Taskfile](https://taskfile.dev) is introduced as an alternative to Makefile. Instead of typing a cluster of long long commands, you can use 

- `task up` to pull up containers
- `task down` to shut down containers
- `task clean` to shut down containers and **remove all data** (be careful!)
- `task migrate` to create tables, insert values from the SQL files under `sql/`.
- `task database` to connect to the interactive shell of the database in container.

Taskfile uses the YAML format, and you can not be more familiar with it if you've read GitHub Actions workflow before.

### Go
*The language we Gophers love*. The [Go](https://go.dev) version of this project is 1.26.2.