#!/bin/sh
mkdir /data
chmod 777 /data -R
#ln -s /home/node-v14.20.0/bin/node /usr/bin/node
export PATH=$PATH:/home/node-v14.20.0/bin
#export LD_LIBRARY=$LD_LIBRARY:
mkdir /data/test_app1
chmod 777 /data/test_app1
cp flow.json /data/test_app1/flows.json
cp node_modules /data/test_app1/ -rf
chmod 777 /data -R
export NODE_PATH=/home/node-v14.20.0/lib/node/node-red/node_modules:/data/test_app1/node_modules
cd /home/node-v14.20.0/lib/node/node-red
npm --no-update-notifier --no-fund start --cache /data/test_app1/.npm -- --port 1881 --userDir /data/test_app1
node test.js
