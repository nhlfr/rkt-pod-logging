# rkt-pod-logging

```
go build -o hello .
acbuild begin
acbuild set-name example.com/hello
acbuild copy hello /bin/hello
acbuild set-exec /bin/hello
acbuild write hello-latest-linux-amd64.aci
acbuild end
sudo RKT_EXPERIMENT_APP=true RKT_EXPERIMENT_ATTACH=true $(which rkt) run --insecure-options=image hello-latest-linux-amd64.aci --stdin=stream --stdout=stream --stderr=stream
sudo RKT_EXPERIMENT_APP=true RKT_EXPERIMENT_ATTACH=true $(which rkt) attach <id_of_pod>
```
