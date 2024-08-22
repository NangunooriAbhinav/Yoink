#!/bin/bash

# Define variables
APP_NAME="yoink"
INSTALL_DIR="/usr/local/bin"
TARGET="$INSTALL_DIR/$APP_NAME"

# Step 1: Build the Go executable
echo "Building the Go program..."
go build -o $APP_NAME

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Build failed. Please check your Go code."
    exit 1
fi

# Step 2: Check if the executable already exists and remove it
if [ -f $TARGET ]; then
    echo "Removing existing $APP_NAME from $INSTALL_DIR..."
    sudo rm -f $TARGET

    # Check if the removal was successful
    if [ $? -ne 0 ]; then
        echo "Failed to remove existing $APP_NAME. Please check your permissions."
        exit 1
    fi
fi

# Step 3: Move the new executable to the installation directory
echo "Installing $APP_NAME to $INSTALL_DIR..."
sudo mv $APP_NAME $INSTALL_DIR

# Check if the move was successful
if [ $? -ne 0 ]; then
    echo "Failed to install $APP_NAME. Please check your permissions."
    exit 1
fi

# Step 4: Make sure the executable is in the PATH and executable
sudo chmod +x $TARGET

# Step 5: Confirm installation
if command -v $APP_NAME >/dev/null 2>&1; then
    echo "$APP_NAME has been installed successfully and is available globally."
else
    echo "Failed to install $APP_NAME globally."
    exit 1
fi

echo "Installation complete. You can now use the command '$APP_NAME'."
