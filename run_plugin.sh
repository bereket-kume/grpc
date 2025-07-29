#!/bin/bash

# Script to run different bank plugins

PLUGIN_NAME=$1
PORT=$2

if [ -z "$PLUGIN_NAME" ]; then
    echo "Usage: $0 <plugin_name> [port]"
    echo ""
    echo "Available plugins:"
    echo "  bankabc    - BankABC plugin (default port: 50051)"
    echo "  bankxyz    - BankXYZ plugin (default port: 50052)"
    echo "  template   - Template plugin (default port: 50053)"
    echo ""
    echo "Examples:"
    echo "  $0 bankabc"
    echo "  $0 bankxyz 50052"
    echo "  $0 template 50053"
    exit 1
fi

# Set default port if not provided
if [ -z "$PORT" ]; then
    case $PLUGIN_NAME in
        "bankabc")
            PORT=50051
            ;;
        "bankxyz")
            PORT=50052
            ;;
        "template")
            PORT=50053
            ;;
        *)
            echo "Unknown plugin: $PLUGIN_NAME"
            exit 1
            ;;
    esac
fi

PLUGIN_DIR="plugins/$PLUGIN_NAME"

if [ ! -d "$PLUGIN_DIR" ]; then
    echo "Plugin directory not found: $PLUGIN_DIR"
    echo "Available plugins:"
    ls -1 plugins/
    exit 1
fi

echo "Starting $PLUGIN_NAME plugin on port $PORT..."
echo "Plugin directory: $PLUGIN_DIR"
echo "Press Ctrl+C to stop"
echo ""

cd "$PLUGIN_DIR"
go run main.go 