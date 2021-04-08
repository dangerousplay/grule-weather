package main

import (
	"fmt"
	"github.com/grule_weather/rules"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

func main() {

	var weather, humidity, temperature rules.Qualifier

	fmt.Println("What is the weather forecast to storm?")
	_, _ = fmt.Scanln(&weather)

	fmt.Println("What is the humidity?")
	_, _ = fmt.Scanln(&humidity)

	fmt.Println("What is the temperature?")
	_, _ = fmt.Scanln(&temperature)

	travel := rules.Travel{
		WeatherStormForecast: weather,
		Humidity:             humidity,
		Temperature:          temperature,
		ShouldTravel:         false,
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("Travel", &travel)
	if err != nil {
		panic(err)
	}

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	err = ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(rules.Rules)))

	if err != nil {
		panic(err)
	}

	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	eng1 := &engine.GruleEngine{MaxCycle: 100}
	err = eng1.Execute(dataContext, kb)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Should you travel? %t\n", travel.ShouldTravel)
}
