package main
import "fmt"
func main() {
    var nome string
    var idade int
    var faixa string
    fmt.Scan(&nome)
    fmt.Scan(&idade)

    if idade < 12 {
        faixa = "crianca"
    }else if idade < 18{
        faixa = "jovem"
    }else if idade < 65{
        faixa = "adulto"
    }else if idade < 1000{
        faixa = "idoso"
    }else{
        faixa = "mumia"
    }

    fmt.Println(nome, "eh" ,faixa)
}

