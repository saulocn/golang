package matematica

//Calculo funcão q realiza um cálculo qualquer, basta que se passe a função
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}
