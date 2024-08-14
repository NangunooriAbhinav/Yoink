#!/bin/bash

# Define variables
APP_NAME="yoink"
INSTALL_DIR="/usr/local/bin"

# Step 1: Build the Go executable
echo "Building the Go program..."
go build -o $APP_NAME

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Build failed. Please check your Go code."
    exit 1
fi

# Step 2: Move the executable to the installation directory
echo "Installing the $APP_NAME to $INSTALL_DIR..."
sudo mv $APP_NAME $INSTALL_DIR

# Check if the move was successful
if [ $? -ne 0 ]; then
    echo "Failed to install $APP_NAME. Please check your permissions."
    exit 1
fi

# Step 3: Make sure the executable is in the PATH and executable
sudo chmod +x $INSTALL_DIR/$APP_NAME

# Step 4: Confirm installation
if command -v $APP_NAME >/dev/null 2>&1; then
    echo "$APP_NAME has been installed successfully and is available globally."
else
    echo "Failed to install $APP_NAME globally."
    exit 1
fi

echo "Installation complete. You can now use the command '$APP_NAME'."
