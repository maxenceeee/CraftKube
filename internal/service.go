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
	Enable       bool
	MinInstances int
	MaxInstances int
	Cooldown     int   // in seconds
	TriggerLogic Logic // AND / OR on multiple Policies by autoScaling policy type
	Upscale      UpscalePolicy
	Downscale    DownscalePolicy
}

type UpscalePolicy struct {
	Policies []Policy
}

type DownscalePolicy struct {
	Policies []Policy
}

type Policy struct {
	Name        string
	Type        string
	Inhibitor   Inhibitor
	ScaleAmount int
}

type ValueType string

const (
	ValueTypeNumber     ValueType = "Number"
	ValueTypePercentage ValueType = "Percentage"
	ValueTypeString     ValueType = "String"
)

type Inhibitor struct {
	Threshold interface{}
	Operator  string
	ValueType ValueType
}
