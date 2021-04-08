# Exemplo de sistema especialista
Implementação de exemplo utilizando [grule-engine](https://github.com/hyperjumptech/grule-rule-engine)

Para executar o projeto precisará de um compilador golang, executando no projeto:
```shell
go run main.go
```

Exemplo de input:
```
What is the weather forecast to storm?
HIGH
What is the humidity?
HIGH
What is the temperature?
MEDIUM
Should you travel? true
```



## Regras implementadas

- Regra 01: Se Umidade é baixa e Temperatura é média ou Temperatura é alta, então Viajar é sim
- Regra 02: Se Umidade é baixa e Temperatura é baixa, então Viajar é não
- Regra 03: Se Umidade é alta e previsão de chuva é baixa, então Viajar é sim
- Regra 04: Se Humidade é alta e previsão de chuva é média ou previsão de chuva é alta e temperatura é média, então viajar é sim
- Regra 05: Se Umidade é alta e previsão de chuva é média ou previsão de chuva é alta e temperatura é alta, então Viajar é não

