#!/bin/bash
PORT=${1:-6021}

if [ "${_SERVICE_HOST_PORT}" == "" ]; then
	__IGNORE=1 #do nothing
else
	PORT=${_SERVICE_HOST_PORT}
fi

# prepare service
CAN_START_SERVICE="true"

if [ "${CAN_START_SERVICE}" == "false" ]; then
	exit 2
fi

# start service
echo '>>>>>'
echo "mode: ${_SERVICE_MODE}"
echo '<<<<<'
if [ "$_SERVICE_MODE" == "deploy" ]; then
	echo "deploy now"
  ./teamdo
else
	echo "dev now"
	BEEGO_RUNMODE=dev ENABLE_DEV_TEST_RESOURCE=1 BEEGO_MODE=dev bee run .
fi
