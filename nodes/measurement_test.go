package nodes

import (
	"testing"

	"strings"

	"github.com/flily/go-a2l/syntax"
)

func TestMeasurementWriteStringDefault(t *testing.T) {
	mes := NewMeasurement("some_variable_name", syntax.FLOAT32_IEEE)
	mes.ECUAddress = 0xc0debabe

	expected := strings.Join([]string{
		"/begin  MEASUREMENT",
		"  /* Name                  */        some_variable_name",
		"  /* Long Identifier       */        \"\"",
		"  /* Type                  */        FLOAT32_IEEE",
		"  /* Conversion Method     */        \"\"",
		"  /* Resolution            */        0.000000",
		"  /* Accuracy              */        0.000000",
		"  /* Lower Limit           */        -3.403E+38",
		"  /* Upper Limit           */        3.403E+38",
		"  ECU_ADDRESS        0xC0DEBABE",
		"/end    MEASUREMENT",
	}, "\n")

	var sb strings.Builder
	_, err := WriteTo(&sb, "  ", 0, mes)
	if err != nil {
		t.Fatalf("WriteTo() error: %v", err)
	}

	got := sb.String()
	if got != expected {
		t.Errorf("Measurement.WriteTo() =\n%s; want\n%s", got, expected)
	}
}
