#!/bin/bash

apt-get update
apt-get install -y wget libicu55 liblttng-ust0
wget -q https://packages.microsoft.com/config/ubuntu/16.04/packages-microsoft-prod.deb
dpkg -i packages-microsoft-prod.deb
apt-get install -y apt-transport-https
apt-get update
apt-get install -y powershell
pwsh -c 'Install-Module -Name Vmware.PowerCLI -Force'
