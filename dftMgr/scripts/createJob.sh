#!/bin/bash

SERVER_URL=http://localhost:12345/api/v1
ACCEPT=application/json

CREATE_TASK_URL=${SERVER_URL}/user/create-task
CREATE_JOB_URL=${SERVER_URL}/user/create-job/$1

JOB_ID=`curl -s --request POST ${CREATE_JOB_URL} | jq -r '.job_id'`

echo Job ID: ${JOB_ID}

while [ "$2" != "" -a "$3" != "" -a "$4" != "" -a "$5" != "" ]; do
    TASK="{
        \"type\": \"$2\",
        \"repetitions\": \"$3\",
        \"destination\": \"$4\",
        \"timeout\": \"$5\"
    }"

    TASK_ID=`curl -s --request POST --header "Content-Type: ${ACCEPT}" --data "${TASK}" ${CREATE_TASK_URL} | jq -r '.task_id'`
    echo Task ID: ${TASK_ID}
    ASSIGN_TASK_URL=${SERVER_URL}/user/assign-task/${JOB_ID}/${TASK_ID}
    curl -s --request POST ${ASSIGN_TASK_URL}
    shift
    shift
    shift
    shift
done

PUBLISH_JOB_URL=${SERVER_URL}/user/publish-job/${JOB_ID}

curl -s --request POST ${PUBLISH_JOB_URL}
