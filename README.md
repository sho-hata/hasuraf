# hasura-fzf

This command has a fzf-like UI that allows you to find and run the file version used by the [hasura cli command](https://hasura.io/docs/latest/graphql/core/hasura-cli/index.html).

<img src="https://media1.giphy.com/media/aZur3hP0ke0p7vJrH8/giphy.gif?cid=790b761159378e0f6a9e19ab11981ecac39458e1a3f2db2f&rid=giphy.gif&ct=g">

## install

### homebrew

```
brew install sho-hata/hasura-fzf/hasura-fzf
```

### go

```
go install github.com/sho-hata/hasura-fzf@latest
```

## binary

[releases](https://github.com/sho-hata/hasura-fzf/releases)

## supported hasura cli commands

- [hasura seed apply --file](https://hasura.io/docs/latest/graphql/core/hasura-cli/hasura_seed_apply.html)
- [hasura migrate apply --version](https://hasura.io/docs/latest/graphql/core/hasura-cli/hasura_migrate_apply.html)
- [hasura migrate delete --version](https://hasura.io/docs/latest/graphql/core/hasura-cli/hasura_migrate_delete.html)

## Usage

```
___  ___  ________  ________  ___  ___  ________  ________  ________
|\  \|\  \|\   __  \|\   ____\|\  \|\  \|\   __  \|\   __  \|\  _____\
\ \  \\\  \ \  \|\  \ \  \___|\ \  \\\  \ \  \|\  \ \  \|\  \ \  \__/
 \ \   __  \ \   __  \ \_____  \ \  \\\  \ \   _  _\ \   __  \ \   __\
  \ \  \ \  \ \  \ \  \|____|\  \ \  \\\  \ \  \\  \\ \  \ \  \ \  \_|
   \ \__\ \__\ \__\ \__\____\_\  \ \_______\ \__\\ _\\ \__\ \__\ \__\
    \|__|\|__|\|__|\|__|\_________\|_______|\|__|\|__|\|__|\|__|\|__|
                       \|_________|

Usage:
  hasuraf [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  seed        Manage seed data.
  migrate     Manage migrations on the database.
  delete      clear migrations from local project and server.

Flags:
  -h, --help   help for hasuraf

Use "hasuraf [command] --help" for more information about a command.
```

## cautions when using

- As with the `hasura cli`, run it in the directory where the `config.yml` exists.
- When you use it, put the .env file with "HASURA_GRAPHQL_DATABASE_URL" in the current directory.
  - If the file is located elsewhere, use the \"--envfile\" option to specify the location of the .env file.

## features

## seed apply

Find the seed file to apply and run the \"hasura seed apply\" command.

It will convert as follows

```
hasuraf seed apply
```

↓

```
hasura seed apply --file XXX
```

### options

Compliant with [originnal](https://hasura.io/docs/latest/graphql/core/hasura-cli/hasura_seed_apply.html#hasura-seed-apply).

However, the `--file` option is not accepted.

## migrate apply

Find the migrate version to apply and run the \"hasura migrate apply\" command.

It will convert as follows

```
hasuraf migrate apply
```

↓

```
hasura migrate apply --version XXX
```

### options

Compliant with [originnal](https://hasura.io/docs/latest/graphql/core/hasura-cli/hasura_migrate_apply.html#hasura-migrate-apply).

However, the `--version` option is not accepted.

## migrate delete

Find the migrate version to delete and run the \"hasura migrate delete\" command.

It will convert as follows

```
hasuraf migrate delete
```

↓

```
hasura migrate delete --version XXX
```

### options

Compliant with [originnal](https://hasura.io/docs/latest/graphql/core/hasura-cli/hasura_migrate_delete.html#hasura-migrate-delete).

However, the `--version` option is not accepted.

## supported hasura config version

[config v3](https://hasura.io/docs/latest/graphql/core/migrations/upgrade-v3.html)

## Author

[Shoki Hata(sho-hata)](https://github.com/sho-hata) Released under the MIT License.
