# memory-leak

## Purpose
Recreate scenarios which can invoke the OOM-killer when running a container.

## Build
~~~
podman build . -t memory-leak
~~~

## Run
~~~
$ podman run --memory 500m --name=memory-leak --rm --memory-swap=0 localhost/memory-leak
~~~
By default Podman will enable swap memory with a size equal to the memory limit specified. 
This will turn off swap to make sure OOMs can be recreated at the expected memory limit.