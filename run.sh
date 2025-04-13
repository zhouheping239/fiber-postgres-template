#!/bin/sh

echo "Starting server..."
echo "Project: ${PROJECT_NAME}"

cd /backend_project
# Run the built executable
./${PROJECT_NAME}