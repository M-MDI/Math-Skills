# stat-projects-bin

An infrastructure to run binaries to test the `stat-projects` on all platforms
that support `docker-engine`.

The binaries present in the directory `bin/` can be run directly on most Linux
distributions without using the below instructions.

## How to run

> To be able to run this script you need to have installed locally
> [`docker-engine`](https://docs.docker.com/engine/install/)

Run the script `./run.sh <bin-name>`. The `<bin-name>` must be inside the
`/bin` directory.
