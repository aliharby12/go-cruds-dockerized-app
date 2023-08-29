#!/bin/bash

case "$1" in
    up)
        docker-compose up --build
        ;;
    down)
        docker-compose down
        ;;
    *)
        echo "Usage: $0 {up|down}"
        ;;
esac