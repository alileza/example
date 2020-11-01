# Example go application

## Getting started

```sh
// generate protobuf definition & api documentation into ./autogen directory
// this step is not required if there's no changes on `example.proto` file
$ make gen

// build go application into a binary
// binary will be output into ./bin/example
$ make build

// verify that binary is okay
$ ./bin/example version
example, version  (branch: refs/heads/master, revision: 07f75fac73ec90c1c95f84f71eeb2dd90db08a09)
  build user:       alileza
  build date:       2020-11-01T14:21:17CET
  go version:       go1.15.2

// start serving application
// you should be able to access your application on http://localhost:9000/api/docs
$ ./bin/example serve
# OUTPUTS:
# 
# INFO[0000] [core] parsed scheme: ""                      system=system
# INFO[0000] [core] scheme "" not registered, fallback to default scheme  system=system
# INFO[0000] [core] ccResolverWrapper: sending update to cc: {[{0.0.0.0:9000  <nil> 0 <nil>}] <nil> <nil>}  system=system
# INFO[0000] [core] ClientConn switching balancer to "pick_first"  system=system
# INFO[0000] [core] Channel switches to new LB policy "pick_first"  system=system
# INFO[0000] [core] Subchannel Connectivity change to CONNECTING  system=system
# INFO[0000] Start serving on 0.0.0.0:9000
##################################
```
