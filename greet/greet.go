package greet

/* Variables a nivel de paquetes solo se pueden definir con la palabra reservada var */

//Variable de tipo string, global en el paquete greet
var emoji = "em"

// English retorna "Hola" en el idioma Inglés
func English() string {
	return "Hi " + emoji
}

// Italian retorna "Hola" en el idioma italiano
func Italian() string {
	return "Ciao " + emoji
}
