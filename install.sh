#!/bin/bash

set -e

DEST="/usr/local/bin"  # Default install location

# Parse command-line arguments
while [[ $# -gt 0 ]]; do
    case "$1" in
        --to)
            DEST="$2"
            shift 2
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Ensure destination directory exists
if [[ ! -d "$DEST" ]]; then
    echo "Error: Destination directory $DEST does not exist."
    exit 1
fi

# Define repository and latest release API
REPO="sfkleach/monogram"
LATEST_RELEASE_URL="https://api.github.com/repos/$REPO/releases/latest"

# Detect OS and architecture
ARCH=$(uname -m)
OS=$(uname | tr '[:upper:]' '[:lower:]')

# Determine the correct asset name
if [[ "$OS" == "linux" ]]; then
    if [[ "$ARCH" == "x86_64" ]]; then
        ASSET_NAME="monogram-linux-x86_64.tar.gz"
    elif [[ "$ARCH" == "aarch64" ]]; then
        ASSET_NAME="monogram-linux-arm64.tar.gz"
    else
        echo "Unsupported architecture: $ARCH"
        exit 1
    fi
elif [[ "$OS" == "darwin" ]]; then
    if [[ "$ARCH" == "x86_64" ]]; then
        ASSET_NAME="monogram-macos-intel.zip"
    elif [[ "$ARCH" == "arm64" ]]; then
        ASSET_NAME="monogram-macos-arm64.zip"
    else
        echo "Unsupported Mac architecture: $ARCH"
        exit 1
    fi
elif [[ "$OS" == "windows" ]]; then
    ASSET_NAME="monogram-windows.zip"
else
    echo "Unsupported OS: $OS"
    exit 1
fi

# Create a temporary working directory
TMP_DIR=$(mktemp -d)

# Cleanup function to safely remove files
cleanup() {
    echo "Cleaning up temporary files..."
    rm -f "$TMP_DIR"/monogram*
    rm -f "$TMP_DIR"/archive_*
    rmdir "$TMP_DIR"
}
# trap cleanup EXIT  # Ensure cleanup runs on script exit

# Fetch the latest release data
echo "Fetching latest Monogram release..."
ASSET_URL=$(curl -s $LATEST_RELEASE_URL | grep "browser_download_url" | cut -d '"' -f 4 | grep "$ASSET_NAME")

if [[ -z "$ASSET_URL" ]]; then
    echo "Error: Unable to find Monogram binary for your system."
    exit 1
fi

# Download archive into the temporary directory
echo "Downloading Monogram from $ASSET_URL into $TMP_DIR..."
curl -L -o "$TMP_DIR/archive_$ASSET_NAME" "$ASSET_URL"

echo "Extracting files..."
cd "$TMP_DIR"
if [[ "$ASSET_NAME" == *.tar.gz ]]; then
    tar -xzf "archive_$ASSET_NAME"
elif [[ "$ASSET_NAME" == *.zip ]]; then
    unzip "archive_$ASSET_NAME"
fi

# Move both binaries to the destination, checking write permissions
chmod +x monogram*

if [[ -w "$DEST" ]]; then
    mv monogram* "$DEST/"
else
    echo "No write permissions for $DEST, using sudo..."
    sudo mv monogram* "$DEST/"
fi

echo "Installation complete! The monogram tools are installed to $DEST."
