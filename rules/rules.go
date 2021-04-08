package rules



//Enunciado 1: Viagem
//Regra 01: Se Umidade é baixa e Temperatura é média ou Temperatura é alta, então Viajar é sim
//Regra 02: Se Umidade é baixa e Temperatura é baixa, então Viajar é não
//Regra 03: Se Umidade é alta e previsão de chuva é baixa, então Viajar é sim
//Regra 04: Se Humidade é alta e previsão de chuva é média ou previsão de chuva é alta e
//temperatura é média, então viajar é sim
//Regra 05: Se Umidade é alta e previsão de chuva é média ou previsão de chuva é alta e
//temperatura é alta, então Viajar é não

const (
	Rules = `
rule regra1 "Se humidade é baixa e Temperatura é média ou Temperatura é alta, então Viajar é sim" {
   when
      Travel.Humidity == "LOW"  && (
         Travel.Temperature == "MEDIUM" || Travel.Temperature == "HIGH"
      )
    then
      Travel.ShouldTravel = true;
      Retract("regra1");
}

rule regra2 "Se humidade é baixa e Temperatura é baixa, então Viajar é não" {
   when
      Travel.Humidity == "LOW"  && 
      Travel.Temperature == "LOW" 
    then
      Travel.ShouldTravel = false;
      Retract("regra2");
}

rule regra3 "Se humidade é alta e previsão de chuva é baixa, então Viajar é sim" {
    when
      Travel.Humidity == "HIGH"  && 
      Travel.WeatherStormForecast == "LOW"
    then
      Travel.ShouldTravel = true;
      Retract("regra3");
}

rule regra4 "Se Humidade é alta e previsão de chuva é média ou previsão de chuva é alta e temperatura é média, então viajar é sim" {
    when
      Travel.Humidity == "HIGH"  && Travel.Temperature == "MEDIUM" &&
      (
         Travel.WeatherStormForecast == "HIGH" || Travel.WeatherStormForecast == "MEDIUM" 
      )
    then
      Travel.ShouldTravel = true;
      Retract("regra4");
}

rule regra5 "Se humidade é alta e previsão de chuva é média ou previsão de chuva é alta e temperatura é alta, então Viajar é não" {
    when
      Travel.Humidity == "HIGH"  && Travel.Temperature == "HIGH" &&
      (
         Travel.WeatherStormForecast == "HIGH" || Travel.WeatherStormForecast == "MEDIUM" 
      )
    then
      Travel.ShouldTravel = false;
      Retract("regra5");
}
`

)

type Qualifier string

const (
	HIGH   Qualifier = "HIGH"
	MEDIUM           = "MEDIUM"
	LOW              = "LOW"
)


type Travel struct {
	WeatherStormForecast Qualifier
	Humidity             Qualifier
	Temperature          Qualifier
	ShouldTravel         bool
}
