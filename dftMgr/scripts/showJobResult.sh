#!/bin/bash

SERVER_URL=http://localhost:12345/api/v1
ACCEPT=application/json

GET_JOB_RESULT_URL=${SERVER_URL}/user/get-result/$1

curl -s --request GET "Content-Type: ${ACCEPT}" ${GET_JOB_RESULT_URL} | sed 's/\\n/\
/g'