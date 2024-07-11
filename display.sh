kubectl --context kind-tracetest create deployment hello --image=registry.cn-hangzhou.aliyuncs.com/acejilam/alpine:latest -- sleep infinity

kubectl --context kind-tracetest scale deployment hello --replicas=5

open -a '/Applications/Google Chrome.app' http://127.0.0.1:16686


#resource=deployments name=hello