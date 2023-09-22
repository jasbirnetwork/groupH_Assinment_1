# Project Name groupH_Assignment_1

Calculator

## Usage
```
go run main.go
```

### This will prompt user to enter first number, operator and second number 
```
Enter first number: 2.0
Enter an operator (+, -, *, /, %, ^): +
Enter second number: 2.0
```
### Output (base on above inputs)
```
2.0 + 2.0 = 4.0
```

## Explaination

Base on the user inputs switch function inside main.go file will call approprate function base in operator user entered.

## Functions

### Function for Mod
```
func Mod(num1 float64, num2 float64) float64 {
	return math.Mod(num1, num2)
}
```
