package dag

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/util/yaml"
)

var testData = `dag:
- name: step1
  type: shell
  command: bash
  content: echo 1024
  output_alias:
    # alias shell stdout to size, and put it into the dag context
    size: null
  debug: true

- name: step2
  type: http
  command: GET http://localhost:8023/apiv1/test
  content: |
    {
      "key1": "some value",
      "size": ${{size}}
      }
  timeout: 10
  headers:
    Authorization: "Basic base64"
    User-Agent: "this is a http task"
  output_alias:
    # for jq
    # alias http response data.obj to newObj, and put it into the dag context
    data: newObj
    codex: newCode
  depends:
    - step1
  conditions:
    - input: ${{size}}
      expression: ">="
      expected: 100
  debug: true

- name: step3
  conditions:
    - input: ${{size}}
      expression: "<"
      expected: 100
  type: shell
  command: bash
  content: echo 'this value is less than 100!'
  depends:
    - step1
    - step4
  debug: true

- name: step4
  type: shell
  command: bash
  content: echo 'step 4'
  debug: true

- name: step5
  type: shell
  command: bash
  content: echo 'step 5'
  depends:
    - step2
  debug: true`

func TestDag_Run(t *testing.T) {
	type Data struct {
		Dag []*DagNode `json:"dag"`
	}
	data := &Data{}
	if err := yaml.Unmarshal([]byte(testData), data); err != nil {
		t.Error(err)
		return
	}
	dag := BuildDag(data.Dag)
	if err := dag.Run(context.TODO()); err != nil {
		t.Error(err)
		return
	}
}
