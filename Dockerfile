FROM alpine

RUN apk add --no-cache coreutils iputils curl jq bash

COPY dftMgr/scripts/createClient.sh /createClient.sh
COPY dftNode/dftNode /dftNode
RUN chmod +x /createClient.sh && chmod +x /dftNode
CMD ["/createClient.sh"] 

