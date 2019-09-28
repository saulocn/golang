package matematica

//Divisor é a função que realiza a divisão entre dois números
func Divisor(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
