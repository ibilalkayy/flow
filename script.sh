#!/bin/bash

# Desired time in HH:MM format
DESIRED_TIME="19:31"

DESIRED_WEEKDAY="sunday"

# Log file location
LOG_FILE="${SCRIPT_PATH}logfile.log"

# Function to install and configure NTP
install_and_sync_ntp() {
    # Check if NTP is installed
    if ! command -v ntpd &> /dev/null; then
        echo "$(date): NTP not found. Installing NTP..." >> "$LOG_FILE"
        
        # Install NTP based on the package manager available
        if command -v apt-get &> /dev/null; then
            sudo apt-get update && sudo apt-get install -y ntp
        elif command -v yum &> /dev/null; then
            sudo yum install -y ntp
        elif command -v dnf &> /dev/null; then
            sudo dnf install -y ntp
        else
            exit 1
        fi
    fi

    # Start and enable NTP service
    if command -v systemctl &> /dev/null; then
        sudo systemctl start ntpd
        sudo systemctl enable ntpd
    elif command -v service &> /dev/null; then
        sudo service ntp start
        sudo update-rc.d ntp defaults
    fi

    # Force synchronize the clock immediately
    if command -v ntpd &> /dev/null; then
        sudo ntpd -gq
    elif command -v ntpdate &> /dev/null; then
        sudo ntpdate pool.ntp.org
    fi
}

# Install and synchronize NTP
install_and_sync_ntp

# Get the current time in HH:MM format
CURRENT_TIME=$(date +"%H:%M")

CURRENT_WEEKDAY=$(date +"%A")

# Check if the current time matches the desired time
DESIRED_WEEKDAY="sunday"
    # Run your command and log the output
    echo "$(date): It's time! Running the desired command." >> "$LOG_FILE"
fi