package main

import "fmt"

func main() {
	var mapaPrecioDescuento = map[uint]uint{150:35,120:87,256:65,3002:14}

	imprimirDescuento(mapaPrecioDescuento)

}

func imprimirDescuento(mapa map[uint]uint) {
for precio,descuento	 := range mapa {
	imprimir:=precio-(precio*descuento/100)
	fmt.Println(imprimir)
}


}