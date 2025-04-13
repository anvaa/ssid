#!/bin/bash

# Get the path to the app bundle
DIR="$(cd "$(dirname "$0")" && pwd)"

# Run your Go binary (background)
"$DIR/ssid" &
# Open web browser to the desired address
open https://localhost:5443

exit 0
