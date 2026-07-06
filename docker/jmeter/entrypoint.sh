#!/bin/sh

echo "======================================"
echo " PT Enterprise JMeter Container"
echo "======================================"

# Service Mode
if [ $# -eq 0 ]; then
    echo "Running in service mode..."
    tail -f /dev/null
fi

# Command Mode
exec "$@"