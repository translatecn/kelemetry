// Copyright 2023 The Kelemetry Authors.
//
// Licensed under the Apache License, Version 2.0 (the"License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an"AS IS"BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/kubewharf/kelemetry/cmd/kelemetry"
	_ "github.com/kubewharf/kelemetry/pkg" // include default packages
	"os"
)

func main() {
	os.Setenv("_DEBUG", "x")
	if os.Getenv("_DEBUG") != "" {
		os.Args = append(os.Args, []string{
			"kelemetry",
			"--audit-consumer-enable",
			"--audit-producer-enable",
			"--audit-webhook-enable",
			"--event-informer-enable",
			"--annotation-linker-enable",
			"--owner-linker-enable",
			"--diff-decorator-enable",
			"--diff-controller-enable",
			"--diff-api-enable",
			"--jaeger-storage-plugin-enable",
			"--jaeger-redirect-server-enable",
			"--mq=local",
			"--audit-consumer-partition=0,1,2,3,4,5,6,7",
			"--http-address=0.0.0.0",
			"--http-port=8080",
			"--kube-target-cluster=tracetest",
			"--kube-target-rest-burst=100",
			"--kube-target-rest-qps=100",
			"--kube-config-paths=tracetest=/Users/acejilam/.kube/koord",
			"--diff-cache=etcd",
			"--diff-cache-etcd-endpoints==127.0.0.1:2379",
			"--diff-cache-wrapper-enable",
			"--diff-controller-leader-election-enable=false",
			"--event-informer-leader-election-enable=false",
			"--linker-worker-count=8",
			"--span-cache=etcd",
			"--span-cache-etcd-endpoints=127.0.0.1:2379",
			"--tracer-otel-endpoint=127.0.0.1:4317",
			"--tracer-otel-insecure",
			"--jaeger-cluster-names=tracetest",
			"--jaeger-storage-plugin-address=127.0.0.1:17271",
			"--jaeger-backend=jaeger-storage",
			"--jaeger-storage.span-storage.type=grpc-plugin",
			"--jaeger-storage.grpc-storage.server=127.0.0.1:17271",
		}...)
	}

	kelemetry.Main()
}
