package syntax

func init() {
	for k, v := range nameMapTopLevelKeyword {
		stringMapTopLevelKeyword[v] = k
	}

	for k, v := range nameMapPrimaryKeyword {
		stringMapPrimaryKeyword[v] = k
	}

	for k, v := range nameMapSecondaryKeyword {
		stringMapSecondaryKeyword[v] = k
	}

	for k, v := range nameMapDataType {
		stringMapDataType[v] = k
	}
}
