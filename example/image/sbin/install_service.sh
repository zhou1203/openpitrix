#!/bin/bash

chmod +x ./*
./install_docker.sh
mv ../conf/openpitrix.service /lib/systemd/system/openpitrix.service
systemctl enable openpitrix.service
mkdir -p /opt/openpitrix/conf
mkdir -p /opt/openpitrix/bin
mv start.sh /opt/openpitrix/bin/start.sh
