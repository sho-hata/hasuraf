# hasura-fzf

This command has a fzf-like UI that allows you to find and run the file version used by the [hasura cli command](https://hasura.io/docs/latest/graphql/core/hasura-cli/index.html).

<img src="https://media.giphy.com/media/vcTpTeXym3tOMZWsd6/giphy.gif">

## supported hasura cli commands

- (Developing)[hasura seed apply --file](https://hasura.io/docs/latest/graphql/core/hasura-cli/hasura_seed_apply.html)
- (Developing)[hasura migrate apply --version](https://hasura.io/docs/latest/graphql/core/hasura-cli/hasura_migrate_apply.html)
- (Developing)[hasura migrate delete --version](https://hasura.io/docs/latest/graphql/core/hasura-cli/hasura_migrate_delete.html)

## Usage

1. Download binary.
2. Run the confirmation command after the download is successful.

```
hasuraf
```

### features

### seed apply

Find the seed file to apply and run the \"hasura seed apply\" command.

It will convert as follows

```
hasuraf seed
```

↓

```
hasura seed apply --file XXX
```

#### options

Compliant with [originnal](https://hasura.io/docs/latest/graphql/core/hasura-cli/hasura_seed_apply.html#hasura-seed-apply).

However, the `--file` option is not accepted.

### migrate apply(Developing)

### migrate delete(Developing)

## supported hasura config version

[config v3](https://hasura.io/docs/latest/graphql/core/migrations/upgrade-v3.html)

## Author

[Shoki Hata(sho-hata)](https://github.com/sho-hata)Released under the MIT License.
