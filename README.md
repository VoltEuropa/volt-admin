# volt-admin
An administration CLI to run common commands on our database

```
➜  volt-admin
Help you manage the volt.team database

Usage:
  volt-admin [flags]
  volt-admin [command]

Available Commands:
  help        Help about any command
  login       Log into the server
  member      Manage members
  team        Manage teams

Flags:
  -h, --help   help for volt-admin

Use "volt-admin [command] --help" for more information about a command.
```

# Using it

You need to be an approved administrator to use this program. If you are,
simply login and start using the command line interface:

```
➜  volt-admin login
[Authentication process, only happens once]
➜  volt-admin member find 100651
[Will return info about the user with id 100651]
```

# Building it

```
go build
```
