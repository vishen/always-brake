# Always Break

`always-break` is a simple program that always breaks and exits
with a non-zero exit code. This is useful when testing how
something behaves when a program breaks.

## Configuration

The following env can be set to control the behaviour:

- TIMEOUT_SECONDS=int  # How long to wait before exiting
- MESSAGE=string  # Message to display before exiting
- EXIT_CODE=int  # Exit code to exit with

## Running

```
$ always-break
2019/04/17 14:53:22 Waiting 5s to break...
2019/04/17 14:53:27 always-break has broken!
exit status 1
```

## Testing Kubernetes

This will create a deployment that will continually break and get
restarted.

```
$ kubectl run always-break --image=vishen/always-break:0.0.1 
```

## Docker 

https://hub.docker.com/r/vishen/always-break
