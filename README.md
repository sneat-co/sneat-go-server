# sneat-go-server

This repository generates Sneat server binaries for local development.

To run the Sneat.app backend server from source, execute:

```shell
go run .
```

To functional properly it requires [dev-environment setup](https://github.com/sneat-co/sneat-devenv).

The server will also be available as copiled binary in nearest future.

## Why this repo is so small and why we need `sneat-go-backend`?

The `sneat-go-backend` module contains HTTP handlers, business logic and abstract data access layer.
It does not depend on Firebase and any other vendor-specific packages.

We inject specific implementation in this `sneat-go-server` module, such as:

- [firebase.google.com/go/v4](https://github.com/firebase/firebase-admin-go)
- [cloud.google.com/go/firestore](https://github.com/googleapis/google-cloud-go)
    - [dalgo2firestore](github.com/dal-go/dalgo2firestore)
