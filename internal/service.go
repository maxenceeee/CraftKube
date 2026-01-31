package internal

type Service struct {
	Name  string
	Image string
}

type Logic string

const (
	LogicAND Logic = "AND"
	LogicOR  Logic = "OR"
)

type AutoScalingConfig struct {
	Enable       bool            `yaml:"enable"`
	MinInstances int             `yaml:"minInstances"`
	MaxInstances int             `yaml:"maxInstances"`
	Cooldown     int             `yaml:"cooldown"`     // in seconds
	TriggerLogic Logic           `yaml:"triggerLogic"` // AND / OR on multiple Policies by autoScaling policy type
	Upscale      UpscalePolicy   `yaml:"upscale"`
	Downscale    DownscalePolicy `yaml:"downscale"`
}

type UpscalePolicy struct {
	Policies []Policy `yaml:"policies"`
}

type DownscalePolicy struct {
	Policies []Policy `yaml:"policies"`
}

type Policy struct {
	Name        string    `yaml:"name"`
	Type        string    `yaml:"type"` // e.g., CPU, Memory, CustomMetric
	Inhibitor   Inhibitor `yaml:"inhibitor"`
	ScaleAmount int       `yaml:"scaleAmount"`
}

type ValueType string

const (
	ValueTypeNumber     ValueType = "Number"
	ValueTypePercentage ValueType = "Percentage"
	ValueTypeString     ValueType = "String"
)

type Inhibitor struct {
	Threshold interface{} `yaml:"threshold"`
	Operator  string      `yaml:"operator"`
	ValueType ValueType   `yaml:"valueType"`
}
