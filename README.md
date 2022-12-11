# Rudi 
[![CI/CD](https://github.com/DooomiT/rudi-go/actions/workflows/cicd.yml/badge.svg)](https://github.com/DooomiT/rudi-go/actions/workflows/cicd.yml)

# Project rudi

## Rudi server

The rudi server is a simple HTTP server that communicates with the assemblyai API or a local stt model. It is used to transcribe audio files and return the results to the client.
The local stt model is based on the [coqui stt](https://github.com/coqui-ai/STT) project. And does only support wav files with a sample rate of 16000.

### Preparation if you want to use the local stt

1. Download coqui stt native library
  ```bash
  ./scripts/setup-coqui.sh
  ```

2. Download model files
- Download the model files huge-vocabulary.scorer and model.tflite from [here](https://coqui.ai/english/coqui/v1.0.0-huge-vocab)

3. Add the model files to the model directory

#### Sample audio files

In the audio-files directory you can find some audio samples downloaded from https://github.com/coqui-ai/STT/releases/download/v1.4.0/audio-1.4.0.tar.gz

### Setup

1. Export library path

```bash
export CGO_LDFLAGS="-L$HOME/.coqui/"
export CGO_CXXFLAGS="-I$HOME/.coqui/"
export LD_LIBRARY_PATH="$HOME/.coqui/:$LD_LIBRARY_PATH"
# On macOSX use  export DYLD_LIBRARY_PATH="$HOME/.coqui/:$DYLD_LIBRARY_PATH"
```

2. Install the dependencies

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
  serve-local 

user
  stt         

Additional Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for rudi

Use "rudi [command] --help" for more information about a command.
```

### serve

```plain
Run this command in order to start a api server

Usage:
  rudi serve <assembly-ai-token> [port] [flags]

Flags:
  -h, --help   help for serve
```

### serve-local

```plain
Run this command in order to start a api server with a local stt instance

Usage:
  rudi serve-local <model> [scorer] [port] [flags]

Flags:
  -h, --help   help for serve-local
```

### STT

```plain
This command uploads a audio file and returns the transcribed text

Usage:
  rudi stt [filepath] [flags]

Flags:
  -h, --help   help for stt
```
