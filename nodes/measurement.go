package nodes

import (
	"github.com/flily/go-a2l/syntax"
)

const (
	constMeasurementFieldName           = "Name                 "
	constMeasurementFieldLongIdentifier = "Long Identifier      "
	constMeasurementFieldDataType       = "Type                 "
	constMeasurementFieldConvMethod     = "Conversion Method    "
	constMeasurementFieldResolution     = "Resolution           "
	constMeasurementFieldAccuracy       = "Accuracy             "
	constMeasurementFieldLowerLimit     = "Lower Limit          "
	constMeasurementFieldUpperLimit     = "Upper Limit          "
)

type Measurement struct {
	Name             string
	LongIdentifier   string
	Value            *syntax.Value
	ConversionMethod string
	ResolutionInt    int64
	ResolutionFloat  float64
	AccuracyInt      int64
	AccuracyFloat    float64
	ECUAddress       uint64
}

func NewMeasurement(name string, dataType syntax.DataType) *Measurement {
	m := &Measurement{
		Name:  name,
		Value: syntax.NewValue(dataType),
	}

	return m
}

var _ Node = (*Measurement)(nil)

func (m *Measurement) Keyword() Keyword {
	return syntax.MEASUREMENT
}

func (m *Measurement) NodeType() NodeType {
	return NodeTypePrimary
}

func (m *Measurement) Values() []*TaggedValue {
	dataType := m.Value.DataType()

	result := []*TaggedValue{
		NewTaggedValue(constMeasurementFieldName, m.Name),
		NewTaggedValue(constMeasurementFieldLongIdentifier, m.LongIdentifier),
		NewTaggedValue(constMeasurementFieldDataType, dataType.String()),
		NewTaggedValue(constMeasurementFieldConvMethod, m.ConversionMethod),
		NewTaggedValue(constMeasurementFieldResolution, dataType.Pick(m.ResolutionInt, m.ResolutionFloat)),
		NewTaggedValue(constMeasurementFieldAccuracy, dataType.Pick(m.AccuracyInt, m.AccuracyFloat)),
		NewTaggedValue(constMeasurementFieldLowerLimit, m.Value.LowerLimit),
		NewTaggedValue(constMeasurementFieldUpperLimit, m.Value.UpperLimit),
	}

	return result
}

func (m *Measurement) Children() []Node {
	result := []Node{
		NewGenericHexadecimalNode(syntax.ECU_ADDRESS, NodeTypeSecondary, m.ECUAddress),
	}

	return result
}

func (m *Measurement) Comment() string {
	return ""
}
