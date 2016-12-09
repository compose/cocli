# CoCLI - A Compose CLI

CoCLI is a Go-based application which uses the Compose API to provide the ability
to create, monitor and delete Compose databases.

To use, an environment variable - COMPOSEAPITOKEN must be set. This token value
can be obtained from the Compose console's Account view.

Further details to follow.

```
$ cocli --help
usage: cocli [<flags>] <command> [<args> ...]

A Compose CLI application

Flags:
  --help    Show context-sensitive help (also try --help-long and --help-man).
  --raw     Output raw JSON responses
  --fullca  Show all of CA Certificates

Commands:
  help [<command>...]
    Show help.

  show account
    Show account details

  show deployments
    Show deployments

  show recipe [<recid>]
    Show recipe

  show recipes [<depid>]
    Show recipes for a deployment

  show clusters
    Show available clusters

  show user
    Show current associated user

  create deployment [<flags>] [<name>] [<type>]
    Create deployment

$ cocli --help create deployment
usage: cocli create deployment [<flags>] [<name>] [<type>]

Create deployment

Flags:
  --help                   Show context-sensitive help (also try --help-long and
                           --help-man).
  --raw                    Output raw JSON responses
  --fullca                 Show all of CA Certificates
  --cluster=CLUSTER        Cluster ID
  --datacenter=DATACENTER  Datacenter location

Args:
  [<name>]  New Deployment Name
  [<type>]  New Deployment Type

```
