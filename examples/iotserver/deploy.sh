clear
env GOARM=5 GOARCH=arm GOOS=linux go build

echo "********** START ********** "

ssh pi@appserver 'pkill iotserver'
ssh pi@appserver 'mkdir -p /home/pi/iot/server'
scp config.yml pi@appserver:/home/pi/iot/server
scp run_iotserver.sh pi@appserver:/home/pi
scp iotserver pi@appserver:/home/pi/iot/server
ssh pi@appserver 'chmod +x ./run_iotserver.sh'
ssh pi@appserver './run_iotserver.sh > done.txt'

echo "*********** END ***********"