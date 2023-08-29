#!/bin/bash

# Load environment variables from the .env file
set -a
[ -f .env ] && . .env
set +a

# Execute the provided command or start the default command
if [ "$1" != "" ]; then
    exec "$@"
else
    exec ./main
fi