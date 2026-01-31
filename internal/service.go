package internal

import (
	"fmt"
	"strings"
)

// Service represents the overall structure of a service
type Service struct {
	Name        string            `yaml:"name" json:"name"`
	Image       string            `yaml:"image" json:"image"`
	AutoScaling AutoScalingConfig `yaml:"auto_scaling" json:"auto_scaling"`
}

func (s *Service) Validate() error {
	if strings.TrimSpace(s.Name) == "" {
		return fmt.Errorf("The name of the service is required")
	}
	if strings.TrimSpace(s.Image) == "" {
		return fmt.Errorf("The image of the service is required")
	}
	return s.AutoScaling.Validate()
}

type Logic string

const (
	LogicAND Logic = "AND"
	LogicOR  Logic = "OR"
)

// Define the different types of scaling policies (e.g., CPU, Memory, CustomMetric)
type PolicyType string

const (
	PolicyCPU          PolicyType = "CPU"
	PolicyMemory       PolicyType = "Memory"
	PolicyCustomMetric PolicyType = "CustomMetric"
)

// Operator defines the comparison operator for scaling conditions
type Operator string

const (
	OpGreaterThan Operator = ">"
	OpLessThan    Operator = "<"
	OpLeftEqual   Operator = "<="
	OpRightEqual  Operator = ">="
	OpEqual       Operator = "=="
)

type AutoScalingConfig struct {
	Enable       bool     `yaml:"enable"`
	MinInstances int      `yaml:"min_instances"`
	MaxInstances int      `yaml:"max_instances"`
	Cooldown     int      `yaml:"cooldown"`      // in seconds
	TriggerLogic Logic    `yaml:"trigger_logic"` // AND / OR
	Upscale      []Policy `yaml:"upscale_policies"`
	Downscale    []Policy `yaml:"downscale_policies"`
}

func (a *AutoScalingConfig) Validate() error {
	if !a.Enable {
		return nil
	}

	// Safeguards on instance numbers
	if a.MinInstances < 1 {
		return fmt.Errorf("min_instances must be at least 1")
	}
	if a.MaxInstances < a.MinInstances {
		return fmt.Errorf("max_instances (%d) cannot be less than min_instances (%d)", a.MaxInstances, a.MinInstances)
	}
	if a.Cooldown < 0 {
		return fmt.Errorf("cooldown cannot be negative")
	}

	// Validation of logic
	if a.TriggerLogic != LogicAND && a.TriggerLogic != LogicOR {
		return fmt.Errorf("invalid trigger_logic: %s", a.TriggerLogic)
	}

	// Validation of policies
	for i, p := range a.Upscale {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("upscale_policy[%d]: %w", i, err)
		}
	}
	for i, p := range a.Downscale {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("downscale_policy[%d]: %w", i, err)
		}
	}

	return nil
}

type Policy struct {
	Name        string     `yaml:"name"`
	Type        PolicyType `yaml:"type"`
	Condition   Condition  `yaml:"condition"`
	ScaleAmount int        `yaml:"scale_amount"`
}

func (p *Policy) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("the name of the policy is required")
	}
	if p.ScaleAmount <= 0 {
		return fmt.Errorf("scale_amount must be greater than 0 (received: %d)", p.ScaleAmount)
	}
	return p.Condition.Validate()
}

type ValueType string

const (
	ValueTypeNumber     ValueType = "Number"
	ValueTypePercentage ValueType = "Percentage"
	ValueTypeString     ValueType = "String"
)

type Condition struct {
	Threshold interface{} `yaml:"threshold"`
	Operator  Operator    `yaml:"operator"`
	ValueType ValueType   `yaml:"value_type"`
}

func (c *Condition) Validate() error {
	// Verification of the operator
	validOps := map[string]bool{">": true, "<": true, "==": true, ">=": true, "<=": true}
	if !validOps[string(c.Operator)] {
		return fmt.Errorf("invalid operator: %s", c.Operator)
	}

	// Verification of the value type
	if c.ValueType != ValueTypeNumber && c.ValueType != ValueTypePercentage {
		return fmt.Errorf("invalid value_type: %s", c.ValueType)
	}

	// Verification of the threshold (must be present)
	if c.Threshold == nil {
		return fmt.Errorf("threshold cannot be empty")
	}

	return nil
}
