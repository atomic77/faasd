package handlers

import (
	"fmt"
	"testing"

	"github.com/openfaas/faas-provider/types"
)

func Test_BuildLabels_WithAnnotations(t *testing.T) {
	request := &types.FunctionDeployment{
		Labels:      &map[string]string{"function_name": "echo"},
		Annotations: &map[string]string{"current-time": "Wed 25 Jul 06:41:43 BST 2018"},
	}

	val, err := buildLabels(request)

	if err != nil {
		t.Fatalf("want: no error got: %v", err)
	}

	if len(val) != 4 {
		t.Errorf("want: %d entries in combined label annotation map got: %d", 4, len(val))
	}

	if _, ok := val[fmt.Sprintf("%scurrent-time", annotationLabelPrefix)]; !ok {
		t.Errorf("want: '%s' entry in combined label annotation map got: key not found", "annotation: current-time")
	}
}
