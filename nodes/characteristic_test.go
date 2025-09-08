package nodes

import (
	"testing"

	"strings"

	"github.com/flily/go-a2l/syntax"
)

func TestCharacteristicWriteStringDefault(t *testing.T) {
	cha := NewCharacteristic("some_variable_name", syntax.SWORD)
	cha.ECUAddress = 0xc0debabe

	expected := strings.Join([]string{
		"/begin  CHARACTERISTIC",
		"  /* Name                  */        some_variable_name",
		"  /* Long Identifier       */        \"\"",
		"  /* Type                  */        SWORD",
		"  /* ECU Address           */        0xC0DEBABE",
		"  /* Record Layout         */        \"\"",
		"  /* Maximum Difference    */        0",
		"  /* Conversion Method     */        \"\"",
		"  /* Lower Limit           */        -32768",
		"  /* Upper Limit           */        32767",
		"/end    CHARACTERISTIC",
	}, "\n")

	got, err := WriteString("  ", 0, cha)
	if err != nil {
		t.Fatalf("WriteTo() error: %v", err)
	}

	if got != expected {
		t.Errorf("Characteristic.WriteTo() =\n%s; want\n%s", got, expected)
	}
}

func TestCharacteristicWriteString1(t *testing.T) {
	cha := NewCharacteristic("some_variable_name", syntax.SWORD)
	cha.Type = "VALUE"
	cha.LongIdentifier = "a very long identifier"
	cha.ECUAddress = 0xc0debabe
	cha.RecordLayout = "Scalar_SWORD"
	cha.ConversionMethod = "some_method_name"

	expected := strings.Join([]string{
		"/begin  CHARACTERISTIC",
		"  /* Name                  */        some_variable_name",
		"  /* Long Identifier       */        \"a very long identifier\"",
		"  /* Type                  */        VALUE",
		"  /* ECU Address           */        0xC0DEBABE",
		"  /* Record Layout         */        Scalar_SWORD",
		"  /* Maximum Difference    */        0",
		"  /* Conversion Method     */        some_method_name",
		"  /* Lower Limit           */        -32768",
		"  /* Upper Limit           */        32767",
		"/end    CHARACTERISTIC",
	}, "\n")

	got, err := WriteString("  ", 0, cha)
	if err != nil {
		t.Fatalf("WriteTo() error: %v", err)
	}

	if got != expected {
		t.Errorf("Characteristic.WriteTo() =\n%s; want\n%s", got, expected)
	}
}
