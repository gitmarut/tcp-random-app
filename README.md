## Purpose

This simple App generates random numbers on a TCP connection.

## How to run it

1. Run "go run main.go" in your machine.

2. Or run it as `docker - run --publish 17777:17777 tcp-random-app`
(I have this in dockerhub, so get dockerfile from `docker pull gitmarut/tcp-random-app:v1.1`)

3. Or you can deploy it in Kubernetes from the yaml file given in same directory.


For methods 1 & 2, from another terminal, use nc to connect it.

    nc localhost 17777

On each press of return key(enter) at client, you get a random number.
It will ignore any text entered except "STOP". It will gracefully terminate connection of entering "STOP"
On pressing ctrl+C at client, it will error EOF and close connection to client.

For method 3 - Get of tcp-random-app service clusterIP using.

    kubectl get svc -n tcp-random-app -o=jsonpath='{.items..spec.clusterIP}'


 Bring up another sleep pod in namespace sleep. Yaml file is given. Get the shell of the sleep pod.

    sleeppod=$(kubectl get pods -n sleep -l app=sleep -o=jsonpath='{.items..metadata.name}' )
    kubectl exec -it -n sleep $sleeppod -- sh

From sleep pod's shell execute nc with tcp-random-app service IP and port 17777. Example will be 

    nc 10.96.217.49 17777
