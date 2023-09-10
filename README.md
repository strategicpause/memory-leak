# memory-leak

## Purpose
Recreate scenarios which can invoke the OOM-killer when running a container.

## Build
~~~
podman build . -t memory-leak
~~~

## Run
Locally
~~~
go run main.go memory --max-memory 1GB --block-size 100MB --pause=1s
~~~

Via Podman
~~~
$ podman run \
--memory=1024m \
--name=memory-leak \
--rm \
--memory-swap=0 \
localhost/memory-leak \
./leak memory --max-memory 1GB --block-size 100MB --pause=1s
~~~
By default Podman will enable swap memory with a size equal to the memory limit specified.
This will turn off swap to make sure OOMs can be recreated at the expected memory limit.

TCP Leak
~~~
$ podman run \
--memory=1024m \
--name=tcp-leak \
--rm \
localhost/memory-leak \
./leak tcp
~~~
