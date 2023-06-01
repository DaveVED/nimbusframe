# nimbusframe
NimbusFrame is a logical mapper between Infrastructure as Code cloud resources or services and their image icons.

## Supported IaC Languages
Right now only `terraform` is in the process of being supported. 

## Running tests
All core tests under `./internal/...` can be run as such.

```
% go test ./internal/...
ok  	github.com/DaveVED/nimbusframe/internal/scrapper	0.187s
ok  	github.com/DaveVED/nimbusframe/internal/utils	39.468s
```

For more detail you can add in the `-v` flag.