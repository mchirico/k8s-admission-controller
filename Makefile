
ifndef $(GOPATH)
  export GOPATH=${HOME}/gopath
  ${shell mkdir -p ${GOPATH}}
endif

ifndef $(GOBIN)
  export GOBIN=${GOPATH}/bin
endif


docker-build:
	docker build -t gcr.io/mchirico/warden:v1 -f Dockerfile .


.PHONY: kind
kind:
	kind load docker-image gcr.io/mchirico/warden:v1


.PHONY: curltest
curltest:
	curl -v --cacert certs/clientCom.pem -X POST -H "Content-Type: application/json" -d '{"request":{"uid": "1","object":{"metadata":{"labels":{"billing":"y"}}}}}' POST  https://localhost:5000/validate


.PHONY: load
load:
	kubectl apply -f warden-k8s.yaml
	sleep 10
	kubectl apply -f webhook.yaml


.PHONY: unload
unload:
	kubectl delete -f warden-k8s.yaml
	kubectl delete -f webhook.yaml



.PHONY: k8s119
k8s119:
	go get k8s.io/kubernetes || true
	cd ${GOPATH}/src/k8s.io/kubernetes && git checkout v1.19.1
	go get sigs.k8s.io/kind
#     Node image
	kind build node-image --image=master


cluster:
	kind delete cluster
	kind create cluster --config kind.yaml --image=master



build:
	go build -v .

run:
	docker run --name aibot --rm -it -p 443:443  gcr.io/mchirico/warden:v1






