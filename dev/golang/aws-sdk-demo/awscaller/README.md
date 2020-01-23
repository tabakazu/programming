# awscaller
CLI for shows aws api caller infomations.

# Installation
```
$ git clone https://github.com/tabakazu/awscaller
$ cd awscaller
$ go build
```

# Usage
```
$ ./awscaller
NAME:
   awscaller - Shows aws api caller infomations

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   x.x.x

COMMANDS:
   account   Shows a aws api caller account number
   policies  Shows a aws api caller attached policies
   userid    Shows a aws api caller user id
   username  Shows a aws api caller user name
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Commands
```
$ ./awscaller account
- Account : 000000000000

$ ./awscaller userid
- UserId : ABCDEFGHIJKLMNOPQRSTU

$ ./awscaller username
- UserName : tabakazu

$ ./awscaller policies
- AttachedPolicies :
    - AdministratorAccess
```
