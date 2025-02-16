#!/usr/bin/env bash
# install.sh - A dynamic installer for kbgen

set -e

# Check required commands
command -v curl >/dev/null 2>&1 || { echo >&2 "Error: curl is required but not installed."; exit 1; }
command -v jq >/dev/null 2>&1 || { echo >&2 "Error: jq is required but not installed. Please install jq and retry."; exit 1; }
command -v tar >/dev/null 2>&1 || { echo >&2 "Error: tar is required but not installed."; exit 1; }

echo "Fetching latest kbgen release version..."
KBGEN_VERSION=$(curl -s "https://api.github.com/repos/eminaktas/kbgen/tags" | jq -r '.[0].name')
if [ -z "$KBGEN_VERSION" ]; then
  echo "Error: Unable to fetch latest version from GitHub." >&2
  exit 1
fi
echo "Latest version is $KBGEN_VERSION"

# Determine OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
if [ "$ARCH" = "x86_64" ]; then
  ARCH="amd64"
elif [ "$ARCH" = "arm64" ] || [ "$ARCH" = "aarch64" ]; then
  ARCH="arm64"
else
  echo "Error: Unsupported architecture: $ARCH" >&2
  exit 1
fi

# Construct download URL (removes the leading 'v' from the version)
DOWNLOAD_URL="https://github.com/eminaktas/kbgen/releases/download/${KBGEN_VERSION}/kbgen_${KBGEN_VERSION#v}_${OS}_${ARCH}.tar.gz"
echo "Downloading kbgen from: $DOWNLOAD_URL"

# Download the binary archive
curl -L "$DOWNLOAD_URL" -o kbgen.tar.gz

# Extract the archive while stripping the top-level folder
echo "Extracting kbgen..."
tar -xzvf kbgen.tar.gz --strip-components=1

# Move the binary to /usr/local/bin (requires sudo)
echo "Installing kbgen to /usr/local/bin..."
sudo mv kbgen /usr/local/bin/
sudo chmod +x /usr/local/bin/kbgen

# Cleanup
rm kbgen.tar.gz

echo "kbgen installed successfully!"
kbgen version
