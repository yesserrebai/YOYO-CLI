#!/bin/bash

# Define the directory containing the main Go file
MAIN_DIR="$(dirname "$(dirname "${BASH_SOURCE[0]}")")/cmd"

# Check if the directory exists
if [ ! -d "$MAIN_DIR" ]; then
  echo "Error: Directory not found: $MAIN_DIR"
  exit 1
fi

# Define the path to the main Go file
MAIN_FILE="$MAIN_DIR/main.go"

# Check if the main Go file exists
if [ ! -f "$MAIN_FILE" ]; then
  echo "Error: Main Go file not found: $MAIN_FILE"
  exit 1
fi

# Define the name of the output binary
OUTPUT_BINARY="yoyo-cli"

# Determine the operating system
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

# Compile the binary based on the operating system
case "$OS" in
  linux)
    GOOS=linux GOARCH=amd64 go build -o "${OUTPUT_BINARY}" "$MAIN_FILE"
    ;;
  darwin)
    GOOS=darwin GOARCH=amd64 go build -o "${OUTPUT_BINARY}" "$MAIN_FILE"
    ;;
  windows)
    GOOS=windows GOARCH=amd64 go build -o "${OUTPUT_BINARY}.exe" "$MAIN_FILE"
    ;;
  *)
    echo "Error: Unsupported operating system: $OS"
    exit 1
    ;;
esac

# Move the binary to /usr/local/bin on Linux and macOS
if [ "$OS" == "linux" ] || [ "$OS" == "darwin" ]; then
  sudo mv "${OUTPUT_BINARY}" /usr/local/bin/
  echo "Binary moved to /usr/local/bin for global accessibility."
fi

echo "Binary generated: ${OUTPUT_BINARY}"
echo "you can use yoyo-cli from anywhere, type yoyo-cli init to create a project"
