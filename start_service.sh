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
	if [ "${_TRACING_MODE}" == "disable" ]; then
		echo "[jaeger] DISABLE agent"
	else
		echo "[jaeger] start agent with endpoint:${_ALIYUN_JAEGER_ENDPOINT}, token:${_ALIYUN_JAEGER_TOKEN}"
		./jaeger-agent --reporter.grpc.host-port=${_ALIYUN_JAEGER_ENDPOINT} --jaeger.tags=Authentication=${_ALIYUN_JAEGER_TOKEN} &
	fi
    ./teamdo
else
	echo "dev now"
	BEEGO_RUNMODE=dev ENABLE_DEV_TEST_RESOURCE=1 BEEGO_MODE=dev bee run teamdo
fi
