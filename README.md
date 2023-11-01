# DMX Webserver

Offers simplistic API and converts calls into DMX commands.

- [DMX Webserver](#dmx-webserver)
- [Running](#running)
  - [Example CURLs](#example-curls)
  - [Example WWW](#example-www)

# Running

```sh
# Replace COM5 with whatever port your dmx is attached to
go run cmd\dmxweb\dmxweb.go -dmx-write-port COM5
```

If running/debugging via VS-Code, make sure to pass the necessary flags as args via [launch.json](./.vscode/launch.json). For example

```yaml
"configurations": [
    {
  "name": "my super awesome config",
  # ...
  "args": ["-dmx-write-port", "COM5"]
    }
]
```

*Note: Boolean flags require the use of `=` eg. `-myflag=[true|false]`*

## Example CURLs

Pass value for a channel:

```sh
curl -v -X PATCH -H "Content-Type: application/json" -d "{\"list\": [{\"channel\": 1, \"value\": 150}]}" http://localhost:8080/api/v1/dmx

curl -v -X PATCH -H "Content-Type: application/json" -d "{\"fadeTimeMillis\": 2500, \"scene\": {\"list\": [{\"channel\": 1, \"value\": 150}]}}" http://localhost:8080/api/v1/dmx/fade

curl -v -X PUT http://localhost:8080/api/v1/dmx/clear
```

## Example WWW

Run with `-static ./www` as example to also have a static file server, serving a demo page.

```sh
go run cmd\dmxweb\dmxweb.go -dmx-write-port COM5 -log-level debug -static ./www
```
