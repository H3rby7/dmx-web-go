# DMX Webserver

Offers simplistic API and converts calls into DMX commands.

- [DMX Webserver](#dmx-webserver)
- [Running](#running)
  - [Configuring Actions](#configuring-actions)
  - [Example CURLs](#example-curls)
    - [DMX API](#dmx-api)
    - [Trigger API](#trigger-api)
  - [Example Static Serving](#example-static-serving)
- [DEV Todos](#dev-todos)
  - [Mocking / Web Console](#mocking--web-console)
  - [Config Mode](#config-mode)

# Running

```sh
# Replace COM5 with whatever port your dmx is attached to
go run cmd\dmxweb\dmxweb.go -dmx-write-port COM5

# Does not write to DMX, instead writes state to log output.
go run cmd\dmxweb\dmxweb.go -dmx-write-port MOCK -static ./static/example

# MOCK DMX for both Read and Write and bridge with example2 config
go run cmd\dmxweb\dmxweb.go -dmx-write-port MOCK -dmx-read-port MOCK -dmx-bridge -static ./static/example -config configs/example2.yaml
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

## Example Static Serving

Run with `-static ./static/example` as example to also have a static file server, serving a demo page.

```sh
go run cmd\dmxweb\dmxweb.go -dmx-write-port COM5 -log-level debug -static ./static/example
```

# DEV Todos

- [ ] Log who is responsible for DMX updates
- [ ] Improve Documentation 
  - [ ] on the options
  - [ ] on triggers and chases

## Mocking / Web Console

- [X] Console-like Web Interface for MockRead Inputs.
- [ ] Console-like Web Interface to see MockWrite Outputs.
  - [ ] Gets its updates via Websocket

## Config Mode

- [ ] In bridged mode listen to inputs to create a config.yaml
  - [ ] Store scenes from active input
  - [ ] Find dimmers: By asking to turn on all lights, then fade down the Master fader, to see which channels get reduced
- [ ] Nice GUI to edit config.yaml files