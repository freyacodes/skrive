# Skrive REST API v1.0

This document describes the Skrive REST API. The API follows these conventions:

* Unless stated otherwise, successful responses will have the status `200 OK`.
* If a response contains a body, it will have the header `Content-Type: application/json`.
* Currently, all endpoints require authentication.

## Authentication

Authentication is currently only possible by providing the server's password in the `Authorization` header. The password should be prefixed with `Password `. If your password is `FOOBAR`, then your Authorization header should be `Password FOOBAR`.

## Dose entity

| Field       | Description                                                  |
| ----------- | ------------------------------------------------------------ |
| `Id`        | A randomly generated ID. If this is provided as `null` to the server, an ID will be generated. Uniqueness is currently not guaranteed in the log. |
| `Time`      | The time at which the dose was administered, expressed as seconds since the Unix epoch. |
| `Quantity`  | A user-provided string. This field is meant to describe the amount administered, typically a number and a unit. |
| `Substance` | A user-provided string. This field is meant to describe the specific chemical substance administered. |
| `Route`     | A user-provided string. This field is meant to describe the route of administration, e.g. Oral or Sublingual. |

#### Example

```json
{
    "Id": "V3knt5xrdXagezE0mUJNU",
    "Time": 1714504074,
    "Quantity": "1,5 mg",
    "Substance": "Estradiol",
    "Route": "Transdermal"
}
```



## Requests

### `GET /v1/doses`

Returns all doses as a JSON array.

### `POST /v1/doses/append`

Requires a Skrive dose entity in the request body. The dose is added to the log. This endpoint is currently NOT idempotent, and this may change in the future without notice.

This endpoint returns the newly added dose, potentially with a generated ID if none was provided.

### `DELETE /v1/doses/:id`

Deletes the given dose if a dose with that ID is found. No body is returned on success. Returns `404 Not Found` if no such dose is found.
