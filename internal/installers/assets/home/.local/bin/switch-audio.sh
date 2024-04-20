#!/usr/bin/env bash

# Headphone: Arctis Nova Pro Wireless
# A Wireless dock is connectected to the computer via USB and Headphone wirelessly connects to the dock.
#
# This script switches audio output between wireless mode and wired (audio jack) mode.
# You can get the list of available sinks by running: pactl list short sinks

if
    pactl list short sinks | grep alsa_output.pci-0000_06_00.6.analog-stereo | grep RUNNING
then
    pactl set-default-sink alsa_output.usb-SteelSeries_Arctis_Nova_Pro_Wireless-00.iec958-stereo
else
    pactl set-default-sink alsa_output.pci-0000_06_00.6.analog-stereo
fi
