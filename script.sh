#!/bin/bash

# Desired time in HH:MM format
DESIRED_TIME="21:11"

DESIRED_DAY="17"

DESIRED_WEEKDAY="Monday"

# Get the current time in HH:MM format
CURRENT_TIME=$(date +"%H:%M")

CURRENT_DAY=$(date +"%d")

CURRENT_WEEKDAY=$(date +"%A")

# Check if the current time matches the desired time
if [ "$CURRENT_TIME" == "$DESIRED_TIME" ] && [ "$CURRENT_WEEKDAY" == "$DESIRED_WEEKDAY" ]; then
    # Run your command and log the output
    echo "$(date): It's time! Running the desired command."
else
    echo "not set ${CURRENT_WEEKDAY} ${CURRENT_TIME} ${CURRENT_DAY}"
fi