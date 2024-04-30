# Skrive REST API

This document describes the Skrive REST API. The API follows these conventions:

* Unless stated otherwise, successful responses will respond with `200 OK`.
* If a response contains a body, the header `Content-Type: application/json` will follow.
* Currently, all endpoints require authentication.

## Authentication

Authentication is currently only possible by providing the server's password in the `Authorization` header. The password should be prefixed with `Password `. If your password is `FOOBAR`, then your Authorization header should be `Password FOOBAR`.

## Dose entity

```json
{
    "Id": "V3knt5xrdXagezE0mUJNU",
    "Time": 1714504074,
    "Quantity": "1,5 mg",
    "Substance": "Estradiol",
    "Route": "Transdermal"
}
```

`Id` is a globally unique ID. If `null` is provided to the server, then an ID will be generated. `Time` is Unix epoch in seconds. `Quantity`, `Substance`, and `Route` are all user-provided strings that are stored and displayed as-is.


## Requests

### `GET /v1/doses`

Returns all doses as a JSON array.

### `POST /v1/doses/append`

Requires a Skrive dose entity in the request body. The dose is added to the log. This endpoint is currently NOT idempotent, and this may change in the future without notice.

This endpoint returns the newly added dose, potentially with a generated ID if none was provided.

### `DELETE /v1/doses/:id`

Deletes the given dose if a dose with that ID is found. No body is returned on success. Returns `404 Not Found` if no such dose is found.
