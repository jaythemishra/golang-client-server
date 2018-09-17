#!/bin/bash

SERVER_URL=http://192.168.254.61:12345/api/v1
ACCEPT=application/json
CREATE_CLIENT_URL=${SERVER_URL}/user/create-client
CLIENT_ID=`curl -s --request POST ${CREATE_CLIENT_URL} | jq -r '.client_id'`
echo Client ID: ${CLIENT_ID}
../../dftNode/dftNode -client_id=${CLIENT_ID} -server_url=${SERVER_URL}
