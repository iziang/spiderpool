#!/usr/bin/env bash
#
# Copyright (C) 2022-2024 ApeCloud Co., Ltd
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT="$(cd "$(dirname $0)/../../" && pwd -P)"

if [ -d "${SCRIPT_ROOT}/vendor" ]; then
  export GOFLAGS="-mod=readonly"
fi

CODE_GENERATOR_PATH=$(go list -f '{{.Dir}}' -m k8s.io/code-generator)
source "${CODE_GENERATOR_PATH}/kube_codegen.sh"

APIS_PACKAGE="kbapi"
OUTPUT_PACKAGE=github.com/spidernet-io/spiderpool/kbapi/client/workloads/v1alpha1

kube::codegen::gen_client "${APIS_PACKAGE}" \
  --with-watch \
  --boilerplate "${SCRIPT_ROOT}/tools/scripts/boilerplate.txt" \
  --clientset-name client \
  --versioned-name v1alpha1 \
  --output-dir "${APIS_PACKAGE}" \
  --output-pkg "${OUTPUT_PACKAGE}"


