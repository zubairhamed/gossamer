cd /home/pi/iot/server
rm -rf output.log
nohup ./iotserver > output.log 2>&1&
