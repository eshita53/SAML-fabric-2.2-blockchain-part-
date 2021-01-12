cd ..
cd ..
cd test-network
./network.sh down
./network.sh up createChannel -c mychannel -ca
./network.sh deployCC -ccn basic -ccl go
cd ..
cd asset-transfer-basic/application-javascript
rm -rf wallet
node app.js







