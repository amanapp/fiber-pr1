#!/bin/bash

# Check if SONAR_TOKEN is set
if [ -z "$SONAR_TOKEN" ]; then
    echo "Error: SONAR_TOKEN environment variable is not set."
    echo "Please set it using: export SONAR_TOKEN=sqp_0935bbb98c83d7ad0af21d2c37a0a79e3bdb2dce"
    exit 1
fi

echo "DEBUG: Using token starting with: ${SONAR_TOKEN:0:7}..."

# Run Sonar Scanner using Docker
docker run \
    --rm \
    --network="host" \
    -v "$(pwd):/usr/src" \
    sonarsource/sonar-scanner-cli \
    -Dsonar.host.url="http://localhost:9000" \
    -Dsonar.token="$SONAR_TOKEN"
