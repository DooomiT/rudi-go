# Check os and install coqui stt

if [[ -d "${HOME}/.coqui" ]]; then
    echo "Coqui STT already installed"
    exit 0
fi

if [[ "$OSTYPE" == "darwin"* ]]; then
    # Mac OSX
    echo "Installing Coqui STT for Mac OSX"
    mkdir -p ${HOME}/.coqui
    cd ${HOME}/.coqui/
    curl -LO https://github.com/coqui-ai/STT/releases/download/v1.4.0/native_client.tflite.macOS.tar.xz
    tar -xvzf native_client.tflite.macOS.tar.xz
    cd -
fi

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Linux
    echo "Installing Coqui STT for Linux"
    mkdir -p ${HOME}/.coqui
    cd ${HOME}/.coqui/
    curl -LO https://github.com/coqui-ai/STT/releases/download/v1.4.0/native_client.tflite.Linux.tar.xz
    tar -xvf native_client.tflite.Linux.tar.xz
    cd -
fi
