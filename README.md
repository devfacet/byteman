# Byteman

[![Godoc][doc-image]][doc-url] [![Release][release-image]][release-url] [![Build][build-image]][build-url]

A Golang library that provides functions for bytes and bits.

## Usage

See [byteman_test.go](byteman_test.go), [numbers_test.go](numbers_test.go) and [strings_test.go](strings_test.go).

## Test

```shell
# Test everything:
make test

# For BDD development:
# It will open a new browser window. Make sure:
#   1. There is no errors on the terminal window.
#   2. There is no other open GoConvey page.
make test-ui

# Benchmarks
make test-benchmarks
```

## Release

```shell
# Update and commit CHANGELOG.md first (i.e. git add CHANGELOG.md && git commit -m "v1.0.0").
# Set GIT_TAG using semver (i.e. GIT_TAG=v1.0.0)
make release GIT_TAG=
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md)

## License

Licensed under The MIT License (MIT)  
For the full copyright and license information, please view the LICENSE.txt file.

[doc-url]: https://pkg.go.dev/github.com/devfacet/byteman
[doc-image]: https://pkg.go.dev/badge/github.com/devfacet/byteman

[release-url]: https://github.com/devfacet/byteman/releases/latest
[release-image]: https://img.shields.io/github/release/devfacet/byteman.svg?style=flat-square

[build-url]: https://github.com/devfacet/byteman/actions/workflows/test.yaml
[build-image]: https://github.com/devfacet/byteman/workflows/Test/badge.svg
