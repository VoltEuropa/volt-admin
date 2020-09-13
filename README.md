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

Ready-to-use binaries are available at https://github.com/VoltEuropa/volt-admin/releases.
Download the program for your operating system and add it to your path in order to
be able to run `volt-admin` from your shell.

**You need to be an approved administrator to use this program**.
If you are, simply login and start using the command line interface:

```
➜  volt-admin login
[Authentication process, only happens once]
➜  volt-admin member get 100651
[Will return info about the user with id 100651]
```

## Environment variables

- `VOLT_ADMIN_SERVER`: the address of the volt-admin server (default: `http://volt.team:8090`).
- `VOLT_ADMIN_TOKEN`: the path of the token file used to communicate to the volt-admin server (default: `~/.volt-admin-token.json`).

# Building it

This requires [Go](https://go.dev/) >= 1.15.

```
go build
```

## Releasing it
To build the binaries for all operating systems, there is a script for macOS and Linux that builds volt-admin for all operating systems:

```
./build.sh
```

The binaries will be locating under the directory `bin` as `.tar.gz` archives.
