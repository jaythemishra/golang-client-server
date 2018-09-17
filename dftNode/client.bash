#!/bin/bash

SERVERURL=http://localhost:12345/api/v1
CLIENT_ID=4909
ACCEPT1=application/json
GETJOBURL=${SERVERURL}/client/get-job/${CLIENT_ID}

JOB_ID=`curl -s --header "Content-Type: ${ACCEPT1}" ${GETJOBURL} | jq -r '.job_id'`
POSTRESULTURL=${SERVERURL}/client/post-result/${JOB_ID}


curl -s --header "Content-Type: ${ACCEPT1}" ${GETJOBURL} | jq -r '.tasks[] | .task_id' > task_ids.txt
curl -s --header "Content-Type: ${ACCEPT1}" ${GETJOBURL} | jq -r '.tasks[] | .details' > tasks.txt
curl -s --header "Content-Type: ${ACCEPT1}" ${GETJOBURL} | jq -r '.tasks[] | .timeout' > timeouts.txt
curl -s --header "Content-Type: ${ACCEPT1}" ${GETJOBURL} | jq -r '.tasks[] | .start_time' > start_times.txt
paste -d \\n task_ids.txt tasks.txt timeouts.txt start_times.txt > taskinfo.txt

while read TASK_ID && read TASK_DETAILS && read TIMEOUTS && read START_TIME; do
	echo Task ID: ${TASK_ID} >> results.txt
	echo Details: ${TASK_DETAILS} >> results.txt
	echo Timeout: ${TIMEOUTS} >> results.txt
	echo Start Time: ${START_TIME} >> results.txt
	echo Results: >> results.txt
	gtimeout ${TIMEOUTS} ${TASK_DETAILS} >> results.txt || echo Timeout: Command took too long >> results.txt
	echo "" >> results.txt
done < taskinfo.txt


RESULT="{
    \"client_id\": \"${CLIENT_ID}\",
    \"job_id\": \"${JOB_ID}\",
    \"results\": \"`cat results.txt | tr '\n' ' '`\"
}"

rm task_ids.txt tasks.txt timeouts.txt start_times.txt taskinfo.txt results.txt
echo "${RESULT}"

#curl -s --data ${RESULT} --request POST --header "Content-Type: ${ACCEPT1}" ${POSTRESULTURL}
curl -v \
--data "${RESULT}" \
--request POST \
--header "Content-Type: ${ACCEPT1}" \
--header "Accept: ${ACCEPT1}" \
--insecure \
${POSTRESULTURL}
#eof
