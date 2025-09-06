package keywords

type (
	Keyword  int
	DataType int
)

const (
	InvalidKeyword Keyword = iota
	// Top-level keywords
	A2ML_VERSION
	ASAP2_VERSION
	HEADER
	MODULE
	PROJECT
	// Primary keywords
	A2ML
	AXIS_PTS
	BLOB
	CHARACTERISTIC
	COMPU_METHOD
	COMPU_TAB
	COMPU_VTAB
	COMPU_VTAB_RANGE
	FRAME
	FUNCTION
	GROUP
	IF_DATA
	INSTANCE
	MEASUREMENT
	MOD_COMMON
	MON_PAR
	RECORD_LAYOUT
	TRANSFORMER
	TYPEDEF_AXIS
	TYPEDEF_BLOB
	TYPEDEF_CHARACTERISTIC
	TYPEDEF_MEASUREMENT
	TYPEDEF_STRUCTURE
	UNIT
	USER_RIGHTS
	VARIANT_CODING
	// Secondary keywords
	ADDR_EPK
	ALIGNMENT_BYTE
	ALIGNMENT_FLOAT32_IEEE
	ALIGNMENT_FLOAT64_IEEE
	ALIGNMENT_INT64
	ALIGNMENT_LONG
	ALIGNMENT_WORD
	ANNOTATION
	ANNOTATION_LABEL
	ANNOTATION_ORIGIN
	ANNOTATION_TEXT
	ARRAY_SIZE
	AXIS_DESCR
	AXIS_PTS_REF
	AXIS_PTS_X
	AXIS_PTS_Y
	AXIS_PTS_Z
	AXIS_PTS_4
	AXIS_PTS_5
	AXIS_RESCALE_X
	BIT_MASK
	BIT_OPERATION
	BYTE_ORDER
	CALIBRATION_ACCESS
	CALIBRATION_HANDLE
	CALIBRATION_HANDLE_TEXT
	CALIBRATION_METHOD
	COEFFS
	COEFFS_LINEAR
	COMPARISON_QUANTITY
	COMPU_TAB_REF
	CPU_TYPE
	CURVE_AXIS_REF
	CUSTOMER
	CUSTOMER_NO
	DATA_SIZE
	DEFAULT_VALUE
	DEFAULT_VALUE_NUMERIC
	DEF_CHARACTERISTIC
	DEPENDENT_CHARACTERISTIC
	DEPOSIT
	DISCRETE
	DISPLAY_IDENTIFIER
	DIST_OP_X
	DIST_OP_Y
	DIST_OP_Z
	DIST_OP_4
	DIST_OP_5
	ECU
	ECU_ADDRESS
	ECU_ADDRESS_EXTENSION
	ECU_CALIBRATION_OFFSET
	EPK
	ERROR_MASK
	EXTENDED_LIMITS
	FIX_AXIS_PAR
	FIX_AXIS_PAR_DIST
	FIX_AXIS_PAR_LIST
	FIX_NO_AXIS_PTS_X
	FIX_NO_AXIS_PTS_Y
	FIX_NO_AXIS_PTS_Z
	FIX_NO_AXIS_PTS_4
	FIX_NO_AXIS_PTS_5
	FNC_VALUES
	FORMAT
	FORMULA
	FORMULA_INV
	FRAME_MEASUREMENT
	FUNCTION_LIST
	FUNCTION_VERSION
	GUARD_RAILS
	IDENTIFICATION
	// IF_DATA,
	IN_MEASUREMENT
	LAYOUT
	LEFT_SHIFT
	LOC_MEASUREMENT
	MAP_LIST
	MATRIX_DIM
	MAX_GRAD
	MAX_REFRESH
	MEMORY_LAYOUT
	MEMORY_SEGMENT
	MONOTONY
	NO_AXIS_PTS_X
	NO_AXIS_PTS_Y
	NO_AXIS_PTS_Z
	NO_AXIS_PTS_4
	NO_AXIS_PTS_5
	NO_OF_INTERFACES
	NO_RESCALE_X
	NUMBER
	OFFSET_X
	OFFSET_Y
	OFFSET_Z
	OFFSET_4
	OFFSET_5
	OUT_MEASUREMENT
	PHONE_NO
	PHYS_UNIT
	PROJECT_NO
	READ_ONLY
	READ_WRITE
	REF_CHARACTERISTIC
	REF_GROUP
	REF_MEASUREMENT
	REF_MEMORY_SEGMENT
	REF_UNIT
	RESERVED
	RIGHT_SHIFT
	RIP_ADDR_W
	RIP_ADDR_X
	RIP_ADDR_Y
	RIP_ADDR_Z
	RIP_ADDR_4
	RIP_ADDR_5
	ROOT
	SHIFT_OP_X
	SHIFT_OP_Y
	SHIFT_OP_Z
	SHIFT_OP_4
	SHIFT_OP_5
	SIGN_EXTEND
	SI_EXPONENTS
	SRC_ADDR_X
	SRC_ADDR_Y
	SRC_ADDR_Z
	SRC_ADDR_4
	SRC_ADDR_5
	STATIC_RECORD_LAYOUT
	STATUS_STRING_REF
	STEP_SIZE
	SUB_FUNCTION
	SUB_GROUP
	SUPPLIER
	SYMBOL_LINK
	SYSTEM_CONSTANT
	UNIT_CONVERSION
	USER
	VAR_ADDRESS
	VAR_CHARACTERISTIC
	VAR_CRITERION
	VAR_FORBIDDEN_COMB
	VAR_MEASUREMENT
	VAR_NAMING
	VAR_SELECTION_CHARACTERISTIC
	VAR_SEPERATOR
	VERSION
	VIRTUAL
	VURTUAL_CHARACTERISTIC

	InvalidDataType DataType = iota
	UBYTE
	SBYTE
	UWORD
	SWORD
	ULONG
	SLONG
	A_UINT64
	A_INT64
	FLOAT32_IEEE
	FLOAT64_IEEE
)

var (
	nameMapTopLevelKeyword = map[string]Keyword{
		"A2ML_VERSION":  A2ML_VERSION,
		"ASAP2_VERSION": ASAP2_VERSION,
		"HEADER":        HEADER,
		"MODULE":        MODULE,
		"PROJECT":       PROJECT,
	}

	// reverse map for nameMapTopLevelKeyword, initialized in init()
	stringMapTopLevelKeyword = map[Keyword]string{}

	nameMapPrimaryKeyword = map[string]Keyword{
		"A2ML":                   A2ML,
		"AXIS_PTS":               AXIS_PTS,
		"BLOB":                   BLOB,
		"CHARACTERISTIC":         CHARACTERISTIC,
		"COMPU_METHOD":           COMPU_METHOD,
		"COMPU_TAB":              COMPU_TAB,
		"COMPU_VTAB":             COMPU_VTAB,
		"COMPU_VTAB_RANGE":       COMPU_VTAB_RANGE,
		"FRAME":                  FRAME,
		"FUNCTION":               FUNCTION,
		"GROUP":                  GROUP,
		"IF_DATA":                IF_DATA,
		"INSTANCE":               INSTANCE,
		"MEASUREMENT":            MEASUREMENT,
		"MOD_COMMON":             MOD_COMMON,
		"MON_PAR":                MON_PAR,
		"RECORD_LAYOUT":          RECORD_LAYOUT,
		"TRANSFORMER":            TRANSFORMER,
		"TYPEDEF_AXIS":           TYPEDEF_AXIS,
		"TYPEDEF_BLOB":           TYPEDEF_BLOB,
		"TYPEDEF_CHARACTERISTIC": TYPEDEF_CHARACTERISTIC,
		"TYPEDEF_MEASUREMENT":    TYPEDEF_MEASUREMENT,
		"TYPEDEF_STRUCTURE":      TYPEDEF_STRUCTURE,
		"UNIT":                   UNIT,
		"USER_RIGHTS":            USER_RIGHTS,
		"VARIANT_CODING":         VARIANT_CODING,
	}

	// reverse map for nameMapPrimaryKeyword, initialized in init()
	stringMapPrimaryKeyword = map[Keyword]string{}

	nameMapSecondaryKeyword = map[string]Keyword{
		"ADDR_EPK":                     ADDR_EPK,
		"ALIGNMENT_BYTE":               ALIGNMENT_BYTE,
		"ALIGNMENT_FLOAT32_IEEE":       ALIGNMENT_FLOAT32_IEEE,
		"ALIGNMENT_FLOAT64_IEEE":       ALIGNMENT_FLOAT64_IEEE,
		"ALIGNMENT_INT64":              ALIGNMENT_INT64,
		"ALIGNMENT_LONG":               ALIGNMENT_LONG,
		"ALIGNMENT_WORD":               ALIGNMENT_WORD,
		"ANNOTATION":                   ANNOTATION,
		"ANNOTATION_LABEL":             ANNOTATION_LABEL,
		"ANNOTATION_ORIGIN":            ANNOTATION_ORIGIN,
		"ANNOTATION_TEXT":              ANNOTATION_TEXT,
		"ARRAY_SIZE":                   ARRAY_SIZE,
		"AXIS_DESCR":                   AXIS_DESCR,
		"AXIS_PTS_REF":                 AXIS_PTS_REF,
		"AXIS_PTS_X":                   AXIS_PTS_X,
		"AXIS_PTS_Y":                   AXIS_PTS_Y,
		"AXIS_PTS_Z":                   AXIS_PTS_Z,
		"AXIS_PTS_4":                   AXIS_PTS_4,
		"AXIS_PTS_5":                   AXIS_PTS_5,
		"AXIS_RESCALE_X":               AXIS_RESCALE_X,
		"BIT_MASK":                     BIT_MASK,
		"BIT_OPERATION":                BIT_OPERATION,
		"BYTE_ORDER":                   BYTE_ORDER,
		"CALIBRATION_ACCESS":           CALIBRATION_ACCESS,
		"CALIBRATION_HANDLE":           CALIBRATION_HANDLE,
		"CALIBRATION_HANDLE_TEXT":      CALIBRATION_HANDLE_TEXT,
		"CALIBRATION_METHOD":           CALIBRATION_METHOD,
		"COEFFS":                       COEFFS,
		"COEFFS_LINEAR":                COEFFS_LINEAR,
		"COMPARISON_QUANTITY":          COMPARISON_QUANTITY,
		"COMPU_TAB_REF":                COMPU_TAB_REF,
		"CPU_TYPE":                     CPU_TYPE,
		"CURVE_AXIS_REF":               CURVE_AXIS_REF,
		"CUSTOMER":                     CUSTOMER,
		"CUSTOMER_NO":                  CUSTOMER_NO,
		"DATA_SIZE":                    DATA_SIZE,
		"DEFAULT_VALUE":                DEFAULT_VALUE,
		"DEFAULT_VALUE_NUMERIC":        DEFAULT_VALUE_NUMERIC,
		"DEF_CHARACTERISTIC":           DEF_CHARACTERISTIC,
		"DEPENDENT_CHARACTERISTIC":     DEPENDENT_CHARACTERISTIC,
		"DEPOSIT":                      DEPOSIT,
		"DISCRETE":                     DISCRETE,
		"DISPLAY_IDENTIFIER":           DISPLAY_IDENTIFIER,
		"DIST_OP_X":                    DIST_OP_X,
		"DIST_OP_Y":                    DIST_OP_Y,
		"DIST_OP_Z":                    DIST_OP_Z,
		"DIST_OP_4":                    DIST_OP_4,
		"DIST_OP_5":                    DIST_OP_5,
		"ECU":                          ECU,
		"ECU_ADDRESS":                  ECU_ADDRESS,
		"ECU_ADDRESS_EXTENSION":        ECU_ADDRESS_EXTENSION,
		"ECU_CALIBRATION_OFFSET":       ECU_CALIBRATION_OFFSET,
		"EPK":                          EPK,
		"ERROR_MASK":                   ERROR_MASK,
		"EXTENDED_LIMITS":              EXTENDED_LIMITS,
		"FIX_AXIS_PAR":                 FIX_AXIS_PAR,
		"FIX_AXIS_PAR_DIST":            FIX_AXIS_PAR_DIST,
		"FIX_AXIS_PAR_LIST":            FIX_AXIS_PAR_LIST,
		"FIX_NO_AXIS_PTS_X":            FIX_NO_AXIS_PTS_X,
		"FIX_NO_AXIS_PTS_Y":            FIX_NO_AXIS_PTS_Y,
		"FIX_NO_AXIS_PTS_Z":            FIX_NO_AXIS_PTS_Z,
		"FIX_NO_AXIS_PTS_4":            FIX_NO_AXIS_PTS_4,
		"FIX_NO_AXIS_PTS_5":            FIX_NO_AXIS_PTS_5,
		"FNC_VALUES":                   FNC_VALUES,
		"FORMAT":                       FORMAT,
		"FORMULA":                      FORMULA,
		"FORMULA_INV":                  FORMULA_INV,
		"FRAME_MEASUREMENT":            FRAME_MEASUREMENT,
		"FUNCTION_LIST":                FUNCTION_LIST,
		"FUNCTION_VERSION":             FUNCTION_VERSION,
		"GUARD_RAILS":                  GUARD_RAILS,
		"IDENTIFICATION":               IDENTIFICATION,
		"IF_DATA":                      IF_DATA,
		"IN_MEASUREMENT":               IN_MEASUREMENT,
		"LAYOUT":                       LAYOUT,
		"LEFT_SHIFT":                   LEFT_SHIFT,
		"LOC_MEASUREMENT":              LOC_MEASUREMENT,
		"MAP_LIST":                     MAP_LIST,
		"MATRIX_DIM":                   MATRIX_DIM,
		"MAX_GRAD":                     MAX_GRAD,
		"MAX_REFRESH":                  MAX_REFRESH,
		"MEMORY_LAYOUT":                MEMORY_LAYOUT,
		"MEMORY_SEGMENT":               MEMORY_SEGMENT,
		"MONOTONY":                     MONOTONY,
		"NO_AXIS_PTS_X":                NO_AXIS_PTS_X,
		"NO_AXIS_PTS_Y":                NO_AXIS_PTS_Y,
		"NO_AXIS_PTS_Z":                NO_AXIS_PTS_Z,
		"NO_AXIS_PTS_4":                NO_AXIS_PTS_4,
		"NO_AXIS_PTS_5":                NO_AXIS_PTS_5,
		"NO_OF_INTERFACES":             NO_OF_INTERFACES,
		"NO_RESCALE_X":                 NO_RESCALE_X,
		"NUMBER":                       NUMBER,
		"OFFSET_X":                     OFFSET_X,
		"OFFSET_Y":                     OFFSET_Y,
		"OFFSET_Z":                     OFFSET_Z,
		"OFFSET_4":                     OFFSET_4,
		"OFFSET_5":                     OFFSET_5,
		"OUT_MEASUREMENT":              OUT_MEASUREMENT,
		"PHONE_NO":                     PHONE_NO,
		"PHYS_UNIT":                    PHYS_UNIT,
		"PROJECT_NO":                   PROJECT_NO,
		"READ_ONLY":                    READ_ONLY,
		"READ_WRITE":                   READ_WRITE,
		"REF_CHARACTERISTIC":           REF_CHARACTERISTIC,
		"REF_GROUP":                    REF_GROUP,
		"REF_MEASUREMENT":              REF_MEASUREMENT,
		"REF_MEMORY_SEGMENT":           REF_MEMORY_SEGMENT,
		"REF_UNIT":                     REF_UNIT,
		"RESERVED":                     RESERVED,
		"RIGHT_SHIFT":                  RIGHT_SHIFT,
		"RIP_ADDR_W":                   RIP_ADDR_W,
		"RIP_ADDR_X":                   RIP_ADDR_X,
		"RIP_ADDR_Y":                   RIP_ADDR_Y,
		"RIP_ADDR_Z":                   RIP_ADDR_Z,
		"RIP_ADDR_4":                   RIP_ADDR_4,
		"RIP_ADDR_5":                   RIP_ADDR_5,
		"ROOT":                         ROOT,
		"SHIFT_OP_X":                   SHIFT_OP_X,
		"SHIFT_OP_Y":                   SHIFT_OP_Y,
		"SHIFT_OP_Z":                   SHIFT_OP_Z,
		"SHIFT_OP_4":                   SHIFT_OP_4,
		"SHIFT_OP_5":                   SHIFT_OP_5,
		"SIGN_EXTEND":                  SIGN_EXTEND,
		"SI_EXPONENTS":                 SI_EXPONENTS,
		"SRC_ADDR_X":                   SRC_ADDR_X,
		"SRC_ADDR_Y":                   SRC_ADDR_Y,
		"SRC_ADDR_Z":                   SRC_ADDR_Z,
		"SRC_ADDR_4":                   SRC_ADDR_4,
		"SRC_ADDR_5":                   SRC_ADDR_5,
		"STATIC_RECORD_LAYOUT":         STATIC_RECORD_LAYOUT,
		"STATUS_STRING_REF":            STATUS_STRING_REF,
		"STEP_SIZE":                    STEP_SIZE,
		"SUB_FUNCTION":                 SUB_FUNCTION,
		"SUB_GROUP":                    SUB_GROUP,
		"SUPPLIER":                     SUPPLIER,
		"SYMBOL_LINK":                  SYMBOL_LINK,
		"SYSTEM_CONSTANT":              SYSTEM_CONSTANT,
		"UNIT_CONVERSION":              UNIT_CONVERSION,
		"USER":                         USER,
		"VAR_ADDRESS":                  VAR_ADDRESS,
		"VAR_CHARACTERISTIC":           VAR_CHARACTERISTIC,
		"VAR_CRITERION":                VAR_CRITERION,
		"VAR_FORBIDDEN_COMB":           VAR_FORBIDDEN_COMB,
		"VAR_MEASUREMENT":              VAR_MEASUREMENT,
		"VAR_NAMING":                   VAR_NAMING,
		"VAR_SELECTION_CHARACTERISTIC": VAR_SELECTION_CHARACTERISTIC,
		"VAR_SEPERATOR":                VAR_SEPERATOR,
		"VERSION":                      VERSION,
		"VIRTUAL":                      VIRTUAL,
		"VURTUAL_CHARACTERISTIC":       VURTUAL_CHARACTERISTIC,
	}

	// reverse map for nameMapSecondaryKeyword, initialized in init()
	stringMapSecondaryKeyword = map[Keyword]string{}

	nameMapDataType = map[string]DataType{
		"UBYTE":        UBYTE,
		"SBYTE":        SBYTE,
		"UWORD":        UWORD,
		"SWORD":        SWORD,
		"ULONG":        ULONG,
		"SLONG":        SLONG,
		"A_UINT64":     A_UINT64,
		"A_INT64":      A_INT64,
		"FLOAT32_IEEE": FLOAT32_IEEE,
		"FLOAT64_IEEE": FLOAT64_IEEE,
	}

	stringMapDataType = map[DataType]string{}
)

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

func (k Keyword) String() string {
	if s, ok := stringMapTopLevelKeyword[k]; ok {
		return s
	}

	if s, ok := stringMapPrimaryKeyword[k]; ok {
		return s
	}

	if s, ok := stringMapSecondaryKeyword[k]; ok {
		return s
	}

	return "InvalidKeyword"
}

func (t DataType) String() string {
	if s, ok := stringMapDataType[t]; ok {
		return s
	}

	return "InvalidDataType"
}

func GetKeywordByName(name string) Keyword {
	if kw, ok := nameMapTopLevelKeyword[name]; ok {
		return kw
	}

	if kw, ok := nameMapPrimaryKeyword[name]; ok {
		return kw
	}

	if kw, ok := nameMapSecondaryKeyword[name]; ok {
		return kw
	}

	return InvalidKeyword
}

func GetDataTypeByName(name string) DataType {
	if dt, ok := nameMapDataType[name]; ok {
		return dt
	}

	return InvalidDataType
}
