package main

import (
	"fmt"

	"github.com/EmanuelLezcano/hello.git/greet"
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(greet.English())
	fmt.Println(quote.Hello())
	fmt.Println(quoteV3.Concurrency())
}

/*COMANDOS
> go list -m all (Lista las dependencias)
> go mod why ruta/paquete (Porque se importa o no un paquete)
> go list -m -versions ruta/paquete (Lista de todas las versiones disponibles de ese paquete)
>

 VERSIONADO


  M Me  P
v 1. 0 .0
v 1. 0 .1  -> M= Mayor, Me= Menior, P=Patch
v 1. 0 .2

->Patch solo se actualiza cuando corregimos un bug y estas correcciones son compatibles con las anteriores.
->Menior nuevos cambios, nuevas caracteristicas, pero se mantiene la compativilidad con versiones anteriores. Cuando actualizamos esta version, la version del Patch debe ser reiniciada a 0.
->Mayor nuevos cambios, que son compatibles con todas las versiones anteriores. Cuando estoy en una version y quiero decir que la nueva version no es compatible con una anterior, debo aumentar el "Mayor" y reinicializar el Menior y Patch a 0.


> go get ruta/paquete (ACTUALIZAR UN PAQUETE a una nueva version)
> go get ruta/paquete@1.3.4 (ACTUALIZAR UN PAQUETE a una nueva version EN ESPECIFICO)

> Utilizamos un paquete en su version mayor. Para hacer posible esto se coloca por ej v3 al final de la ruta de importacion del paquete.Si el pauquete ya tiene otra ruta de importacion, se puede añadir un nombre a una de las importaciones seguido de la ruta de importacion, ej.
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
	cuando quiero usar quote en la version 3 voy a utilizar quoteV3

> go mod tidy (Limpia las dependencias no utilizadas)

*/
