#!/bin/sh

set -e

echo "Installing JMeter Plugins..."

PLUGIN_DIR=/opt/apache-jmeter/lib/ext

wget https://jmeter-plugins.org/get/ \
-O /tmp/PluginsManager.jar

cp /tmp/PluginsManager.jar $PLUGIN_DIR/

java -cp \
$PLUGIN_DIR/PluginsManager.jar \
org.jmeterplugins.repository.PluginManagerCMDInstaller

PluginsManagerCMD.sh install \
jpgc-casutg,\
jpgc-functions,\
jpgc-standard,\
jpgc-json,\
jpgc-cmd,\
jpgc-graphs-basic,\
jpgc-graphs-additional,\
jpgc-dummy

echo "Plugins Installed"