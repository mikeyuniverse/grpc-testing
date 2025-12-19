# Wrong enum value

Grpc enum is an integer.
Any integer can be passed as value.

## Getting started

Client sends request with invalid status field (enum).

Run server: prints accepted status

```sh
go run cmd/server/main.go
```

Run client: sends status `22`

```sh
go run cmd/client/main.go
```

## Logs

1. Send request

    ```go
    _, err = statusClient.CheckStatus(
        context.Background(), 
        &status.CheckStatusRequest{
            Status: status.Status_STATUS_SECOND + 20,
        })
    ```

2. Server will print

    ```sh
    Raw status:  22
    String:  22
    Number:  22
    ```
