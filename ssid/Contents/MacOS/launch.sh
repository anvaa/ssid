#!/bin/bash

# Get the path to the app bundle
DIR="$(cd "$(dirname "$0")" && pwd)"

# Run your Go binary (background)
"$DIR/ssid" &
# echo "Go bin DIR is $DIR/MacOS/ssid.arm64"
# Open web browser to the desired address
open https://localhost:5005

exit 0
