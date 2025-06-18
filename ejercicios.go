package main


import (
"fmt"
"slices"
"sort"

)

func main (){


lista :=[5] int {5 , 8 , 10 , 20 , 30}
fmt.Println(lista)
//copiar una lista en un sline con susn datos y rango 
sline := make([]int , len(lista))
copy(sline , lista [:])
fmt.Println(sline)
//agregar un elemento
sline = append(sline , 40)
fmt.Println(sline)
//eliminar un elemento 
sline=slices.Delete(sline , 5 , 5+1)
fmt.Println(sline)
//buscar si un valor existe dentro de sline 
valorabuscar:= 21
if slices.Contains(sline , valorabuscar){
	fmt.Printf("el valor %d si existe en la lista.\n" , valorabuscar)
}else {
	fmt.Printf("el valor %d no esta en la lista.\n " , valorabuscar)
}
//ordenar un sline de float de menor a mayor 
precios:= [] float64{10.34 , 66.44 , 66.45 , 100.23 , 1.99}
fmt.Println(precios)
sort.Float64s(precios)
fmt.Println(precios)
//revertir un sliner 
slices.Reverse(precios)
fmt.Println(precios)

//leer e eimprimir datos de una factura 
var nombre string
var apellido string
var monto float64
var iba float64
var total float64
var cedula int
//ingreso de los datos del cliente 
fmt.Println("ingrese nombre del cliente ")
fmt.Scan(&nombre)
fmt.Println("ingrese apellido del cliente ")
fmt.Scan(&apellido)
fmt.Println("ingrese cedula del cliente ")
fmt.Scan(&cedula)
fmt.Println("ingrese monto a pagar del cliente ")
fmt.Scan(&monto) 
//calculo del iba 
iba = monto * 0.16
total = monto + iba
//montrar datos de la factura 
fmt.Println(nombre)
fmt.Println(apellido)
fmt.Println(cedula)
fmt.Println(monto)
fmt.Println(iba)
fmt.Println(total)





}
