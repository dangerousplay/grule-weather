package rules

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTravel_1(t *testing.T) {
	scenarios := []Travel{
		{
			WeatherStormForecast: "",
			Humidity:             "LOW",
			Temperature:          "MEDIUM",
			ShouldTravel:         false,
		},
		{
			WeatherStormForecast: "",
			Humidity:             "LOW",
			Temperature:          "HIGH",
			ShouldTravel:         false,
		},
	}

	for _, scenario := range scenarios {
		executeTest(t, &scenario, true)
	}
}

func TestTravel_2(t *testing.T) {
	scenarios := []Travel{
		{
			WeatherStormForecast: "",
			Humidity:             "LOW",
			Temperature:          "LOW",
			ShouldTravel:         true,
		},
	}

	for _, scenario := range scenarios {
		executeTest(t, &scenario, false)
	}
}

func TestTravel_3(t *testing.T) {
	scenarios := []Travel{
		{
			WeatherStormForecast: "LOW",
			Humidity:             "HIGH",
			Temperature:          "",
			ShouldTravel:         false,
		},
	}

	for _, scenario := range scenarios {
		executeTest(t, &scenario, true)
	}
}

func TestTravel_4(t *testing.T) {
	scenarios := []Travel{
		{
			WeatherStormForecast: "MEDIUM",
			Humidity:             "HIGH",
			Temperature:          "MEDIUM",
			ShouldTravel:         false,
		},
		{
			WeatherStormForecast: "HIGH",
			Humidity:             "HIGH",
			Temperature:          "MEDIUM",
			ShouldTravel:         false,
		},
	}

	for _, scenario := range scenarios {
		executeTest(t, &scenario, true)
	}
}

func TestTravel_5(t *testing.T) {
	scenarios := []Travel{
		{
			WeatherStormForecast: "MEDIUM",
			Humidity:             "HIGH",
			Temperature:          "HIGH",
			ShouldTravel:         true,
		},
		{
			WeatherStormForecast: "HIGH",
			Humidity:             "HIGH",
			Temperature:          "HIGH",
			ShouldTravel:         true,
		},
	}

	for _, scenario := range scenarios {
		executeTest(t, &scenario, false)
	}
}

func executeTest(t *testing.T, travel *Travel, expected bool) {
	dataContext := ast.NewDataContext()
	err := dataContext.Add("Travel", travel)
	if err != nil {
		t.Fatal(err)
	}

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	err = ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(Rules)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	eng1 := &engine.GruleEngine{MaxCycle: 100}
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)

	assert.Equal(t, travel.ShouldTravel, expected)
}
