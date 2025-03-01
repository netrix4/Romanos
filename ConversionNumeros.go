package main

import (
	"fmt"
)

type diccionarioRomano struct {
    signo string
    valor int
}
type restasValidas struct {
    principal string
    coeficientesAledanos []string;
}

const entrada string = "XM";
const longitud int = len(entrada);
var valoresIndividuales [longitud]int;

var pivote string = "";
var sumaTemp int;

var contadorRepeticiones int = 0;
var sumaTotal int = 0;


var diccionarioRestas = [7] restasValidas{
    {"V", []string{"I"}},
    {"X", []string{"I"}},
    {"L", []string{"X"}},
    {"C", []string{"X"}},
    {"D", []string{"X","L"}},
    {"M", []string{"C"}},
}

var numerosValidos = [7] diccionarioRomano{
    {"I", 1},
    {"V", 5}, 
    {"X", 10}, 
    {"L", 50}, 
    {"C", 100},
    {"D", 500},
    {"M", 1000},
}

func retornarValorDeLetra(posicion int) int {
    var valorARetornar int;
    if posicion < longitud {
        digito := string(entrada[posicion])
        for _, numeroDiccionario := range numerosValidos {
            if digito == numeroDiccionario.signo {
                valorARetornar = int(numeroDiccionario.valor)
                valoresIndividuales[posicion] = valorARetornar;
            }
        }
    }else{
        valorARetornar = 0;
    }
    return valorARetornar;
}

func main(){

    // const entrada string = "XXL";

	fmt.Println("Esta es la entrada: ", entrada);

	for i := 0; i < longitud; i++ {

        var valorDigitoRomano = retornarValorDeLetra(i)
        var valorDigitoRomanoSiguiente = retornarValorDeLetra(i+1)

		fmt.Println("Letra:", string(entrada[i]));
        fmt.Println("valor del digito:", valorDigitoRomano)
        fmt.Println("valor del digito siguiente:", valorDigitoRomanoSiguiente)
        fmt.Println()

        // // sumaTemp = valorDigitoRomano
        // if valorDigitoRomanoSiguiente != 0{
        //     if valorDigitoRomano == valorDigitoRomanoSiguiente {
        //         sumaTemp += valorDigitoRomano+valorDigitoRomanoSiguiente;
        //         // contadorRepeticiones++;
    
        //         // if contadorRepeticiones == 3 {
        //         //     contadorRepeticiones=0;
        //         // }
        //     }
        //     if valorDigitoRomano < valorDigitoRomanoSiguiente{
        //         sumaTotal = valorDigitoRomanoSiguiente-sumaTemp;
        //         sumaTemp = 0;
        //     }            
        // }else{
        //     fmt.Println("Se cabaron los digitos")
        //     sumaTotal = valorDigitoRomanoSiguiente+sumaTemp;
        //     sumaTemp =0;
        // }

        // fmt.Println("Temporal acumulada: ",sumaTemp)
	}
    // const entrada string = "XXL";

    var temp int = valoresIndividuales[0];
    var contador int = 0;
    var total int =0;
    for _, item := range valoresIndividuales {
        fmt.Println("aaaa",contador)

        if contador < longitud-1 {
            itemSiguiente := valoresIndividuales[contador+1];


            if item<itemSiguiente {
                // temp+=item
                total += itemSiguiente-temp;
                contador++
                temp=0
            }
            if item == itemSiguiente {
                temp += itemSiguiente
            }
            // if item>itemSiguiente {

            //     total += itemSiguiente;
            //     contador++;
            // }
        }else{
            //Lista en el ultimo index
            fmt.Println("ultimo valor", temp, total)
            // if temp>valoresIndividuales[contador] {
            //     total+= valoresIndividuales[longitud-1]+temp
            // }
        }
        // item+=1;
        contador++;
    }

    

    fmt.Println("Estos son todos los valores de cada digito romano ", valoresIndividuales)
    fmt.Println("Suma2 ", total)

}



// una idea es partirlo en numero (digitos) validos



