# DMX Webserver

Offers simplistic API and converts calls into DMX commands.

- [DMX Webserver](#dmx-webserver)
- [Running](#running)
  - [Configuring Actions](#configuring-actions)
  - [Example CURLs](#example-curls)
    - [DMX API](#dmx-api)
    - [Trigger API](#trigger-api)
  - [Example WWW](#example-www)
- [DEV Todos](#dev-todos)

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

## Configuring Actions

Using a [config file](./configs/example.yaml) you can define 
trigger sources, chases and event sequences to interact with the DMX output and bridge.

The `config` flag allows to specify a different config file.

## Example CURLs

### DMX API

```sh
# Set DMX Channel '1' to value '150'
curl -v -X PATCH -H "Content-Type: application/json" -d "{\"list\": [{\"channel\": 1, \"value\": 150}]}" http://localhost:8080/api/v1/dmx

# Fade DMX Channel '1' to value '150' over '2500' milliseconds
curl -v -X PATCH -H "Content-Type: application/json" -d "{\"fadeTimeMillis\": 2500, \"scene\": {\"list\": [{\"channel\": 1, \"value\": 150}]}}" http://localhost:8080/api/v1/dmx/fade

# Clear DMX (set all DMX channels to 0 immediately)
curl -v -X PUT http://localhost:8080/api/v1/dmx/clear
```

### Trigger API

```sh
# Send trigger signal from source '35406887899400'
curl -v -X POST -H "Content-Type: application/json" -d "{\"source\": \"35406887899400\"}" localhost:8080/api/v1/trigger
```

## Example WWW

Run with `-static ./www` as example to also have a static file server, serving a demo page.

```sh
go run cmd\dmxweb\dmxweb.go -dmx-write-port COM5 -log-level debug -static ./www
```

# DEV Todos

- [ ] Bridge Fade IN/OUT at chase
- [ ] Log who is responsible for DMX updates
- [ ] 