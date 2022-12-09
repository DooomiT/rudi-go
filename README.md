# Rudi 

# Project rudi

## Rudi server

The rudi server is a simple HTTP server that communicates with the assemblyai API. It is used to transcribe audio files and return the results to the client.

### Setup

1. Install the dependencies

```bash
go install
```

2. Build the binary

```bash
go build
```

### Usage


### CLI
```plain
Rudi is a basic api for speech transcription

Usage:
  rudi [command]

general
  serve       

user
  stt         

Additional Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for rudi

Use "rudi [command] --help" for more information about a command.
```

### Serve

```plain
Run this command in order to start a api server

Usage:
  rudi serve <assembly-ai-token> [port] [flags]

Flags:
  -h, --help   help for serve
```

### STT

```plain
This command uploads a audio file and returns the transcribed text

Usage:
  rudi stt [filepath] [flags]

Flags:
  -h, --help   help for stt
```
