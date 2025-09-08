package nodes

import (
	"github.com/flily/go-a2l/strutil"
	"github.com/flily/go-a2l/syntax"
)

const (
	constCharacteristicFieldName           = "Name                 "
	constCharacteristicFieldLongIdentifier = "Long Identifier      "
	constCharacteristicFieldType           = "Type                 "
	constCharacteristicFieldEcuAddress     = "ECU Address          "
	constCharacteristicFieldRecordLayout   = "Record Layout        "
	constCharacteristicFieldMaxDiff        = "Maximum Difference   "
	constCharacteristicFieldConvMethod     = "Conversion Method    "
	constCharacteristicFieldLowerLimit     = "Lower Limit          "
	constCharacteristicFieldUpperLimit     = "Upper Limit          "
)

type Characteristic struct {
	Name              string
	LongIdentifier    string
	Type              string
	Value             *syntax.Value
	ECUAddress        uint64
	RecordLayout      string
	MaximumDifference int64
	ConversionMethod  string
}

var _ Node = (*Characteristic)(nil)

func NewCharacteristic(name string, dataType syntax.DataType) *Characteristic {
	c := &Characteristic{
		Name:  name,
		Type:  dataType.String(),
		Value: syntax.NewValue(dataType),
	}

	return c
}

func (c *Characteristic) Keyword() Keyword {
	return syntax.CHARACTERISTIC
}

func (c *Characteristic) NodeType() NodeType {
	return NodeTypePrimary
}

func (c *Characteristic) Values() []*TaggedValue {
	result := []*TaggedValue{
		NewTaggedValue(constCharacteristicFieldName, c.Name),
		NewTaggedValue(constCharacteristicFieldLongIdentifier, c.LongIdentifier),
		NewTaggedValue(constCharacteristicFieldType, c.Type),
		NewTaggedValue(constCharacteristicFieldEcuAddress, strutil.Hexadecimal(c.ECUAddress)),
		NewTaggedValue(constCharacteristicFieldRecordLayout, c.RecordLayout),
		NewTaggedValue(constCharacteristicFieldMaxDiff, strutil.Number(c.MaximumDifference)),
		NewTaggedValue(constCharacteristicFieldConvMethod, c.ConversionMethod),
		NewTaggedValue(constCharacteristicFieldLowerLimit, c.Value.LowerLimit),
		NewTaggedValue(constCharacteristicFieldUpperLimit, c.Value.UpperLimit),
	}

	return result
}

func (c *Characteristic) Children() []Node {
	return nil
}

func (c *Characteristic) Comment() string {
	return ""
}
