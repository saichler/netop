set -e
cd security
./build.sh
cd ../vnet
./build.sh
cd ../websvr
./build.sh
cd ../device-inv
./build.sh