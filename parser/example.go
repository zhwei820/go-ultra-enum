package parser

//go:generate go-ultra-enum
type ColorEnumEnum struct {
	Red   int `enum:"1"`
	Blue  int `enum:"2"`
	Grenn int `enum:"3"`
}

type TradeResultEnumEnum struct {
	Hedging    string `enum:"hedging"`     // * 对冲中(不主动回调): hedging；
	WaitDeduct string `enum:"wait_deduct"` // * 等待扣款: wait_deduct, 等待扣款；
	Success    string `enum:"success"`     // * 对冲成功: success, 包含期权、现货交易明细；
	Fail       string `enum:"fail"`        // * 对冲失败: fail, 包含失败原因，期权、现货交易明细(失败也有明细)
}

type CountryEnumEnum struct {
	China    int64 `enum:"1"`
	America  int64 `enum:"2"`
	Sinapore int64 `enum:"3"`
}

// generate an enum with default display values. The display values are set to the field names, e.g. `On` and `Off`
type StatusEnumEnum struct {
	On  bool `enum:"true"`
	Off bool `enum:"false"`
}

// generate an enum with display values and descriptions.
type SushiEnumEnum struct {
	Maki    string `enum:"MAKI,Rice and filling wrapped in seaweed"`
	Temaki  string `enum:"TEMAKI,Hand rolled into a cone shape"`
	Sashimi string `enum:"SASHIMI,Fish or shellfish served alone without rice"`
}
