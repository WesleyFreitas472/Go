package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func split(str, separador string) []string {

	palavra := ""
	var result []string

	for i := 0; i < len(str); i++ {	
		char := string(str[i])
		
		if char != separador {
			palavra += char
		} else {
			result = append(result, palavra)
			palavra = ""
		}
	}

	result = append(result, palavra)

	return result
}

func testStrings(){
	str := "Hello World!"

	fmt.Println(str)

	strSplitado := split(str, " ")
	fmt.Println(strSplitado, len(strSplitado))

	strSplitado2 := split("cadeira_mesa_computador_teclado_mouse", "_")
	fmt.Println(strSplitado2)

	strSplitadoBuiltIn := strings.Split("cadeira_mesa_computador_teclado_mouse", "_")
	fmt.Println(strSplitadoBuiltIn)
}

func testArrayAndSlice() {
	
	array := [3]int{10, 20, 30} //estático
	
	slice := array[1:2] //como se fosse um array de ponteiros

	fmt.Println(array,slice)

	slice[0]++ //quando eu altero valor a alteração é refletida no array pq o slice é um vetor

	fmt.Println(array,slice)


	//like forEach
	for i, num := range array {
		fmt.Printf("%d) %d\n", i, num)
	}

	//like forEach ignorando um dos retornos
	for _, num := range array {
		fmt.Printf("%d\n", num)
	}

	//like forEach para ignorar o segundo valor do retorno basta omitir ele na declaração/atribuição
	for i := range array {
		fmt.Printf("%d\n", i)
	}

}

func testMake() {
	//esse make é doido
	//ele me cria um slice de 10 elementos que fazem referencia a elementos de um array de capacidade 20 
	//se eu omitir o terceiro parametro a capacidade do array interno é igual ao tamanho do slice
	slice := make([]int, 10, 20)
	fmt.Println(slice, len(slice), cap(slice))

	s := append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 22)
	fmt.Println(s, len(s), cap(s))

	array := [3]int{1,2,3}
	fmt.Println(array, len(array), cap(array))
	
	//newArray := append(array, 1) isso aqui não funciona pq o append só funciona com array e não com lista

	//Prove que o make cria um slice que referencia elementos de um array
	fmt.Println(slice, len(slice), cap(slice))

	for i := range slice {
		slice[i]++
	}

	fmt.Println(slice, len(slice), cap(slice))

}

func testMap(){
	//maps são estruturas chave-valor (like dict)
	//maps precisam ser inicializados
	//inicializar map usando a função make -> make(map[<tipo_da_chave>],<tipo_do_valor>)
	aprovados := make(map[int]string)

	aprovados[1] = "Wesley"
	aprovados[2] = "Freitas"
	aprovados[3] = "Ramalho"
	
	fmt.Println(aprovados)

	//percorrendo chave-valor do map
	//like dict().items()
	for chave, valor := range aprovados {
		fmt.Printf("chave: %d \t valor: %s\n", chave, valor)
	}

	//para deletar faça
	delete(aprovados, 1)
	fmt.Println("")

	for chave, valor := range aprovados {
		fmt.Printf("chave: %d \t valor: %s\n", chave, valor)
	}


	//outra forma de inicializar um map é assim (literal)
	reprovados := map[int]string{
		1: "João",
		2: "Antonio",
		3: "Pedro", //no último item precisa conter a vírgula ou chave pq se não dá erro de compilação,
	}

	fmt.Println(reprovados)

}

func testFuncaoAnonima() {
	//isso aqui é tipo lambda ou arrow function
	//e tem tambem o retorno limpo, outro recurso da linguagem
	//o result foi declarado no retorno da função
	//quando chamo o return o valor de result será retornado
	var soma = func(v1, v2 int) (result int) {
		result = v1 + v2
		return
	}
	fmt.Printf("Função anonima com retorno limpo. Resultado 2 + 2: %d\n", soma(2, 2))
}

func testFuncaoComoParametro(){
	var soma = func(v1, v2 int) (result int) {
		result = v1 + v2
		return
	}

	var somaMais10 = func(funcaoParametro func(int, int) int, v1, v2 int) int{
		return funcaoParametro(v1, v2) + 10
	}

	fmt.Printf("Soma + 10: %d", somaMais10(soma,2,2))
}

func testFuncaoVariatica() {
	var soma = func(valores ...int) (result int) {
		result = 0
		for _, num := range valores {
			result += num
		}
		return
	}
	slice := []int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("Soma de N números: %d\n",soma(slice...))
}


//struct
type produto struct {
	nome string
	preco float64
	desconto float64
}

//funcao com receiver -> é como se a struct fosse uma classe e a função com o receiver fosse o método
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func testReceiver(){
	//primeira forma de inicializar uma variavel do tipo produto (struct)
	tipoProduto := produto{"Lapis", 100.00, 0.5}
	fmt.Println(tipoProduto)
	fmt.Println(tipoProduto.precoComDesconto())
	//segunda forma de inicializar uma variavel do tipo produto (struct)
	var tipoProduto2 = produto{
		nome: "Caneta",
		preco: 200.00,
		desconto: 0.2,
	}
	fmt.Println(tipoProduto2)
	fmt.Println(tipoProduto2.precoComDesconto())

}

//struct aninhada

type item struct {
	itemID int
	nome string
	valor float64
}

type pedido struct{
	pedidoID int
	itens []item
}

func (p pedido) getValorTotalPedido() float64 {
	result := 0.0

	for _, item := range p.itens {
		result += item.valor
	}

	return result
}

func testStructAninhada() {
	tipoPedido := pedido{
		pedidoID: 1,
		itens: []item{
			{
				itemID: 1,
				nome: "Caneta",
				valor: 10.0,
			},{
				itemID: 2,
				nome: "Lápis",
				valor: 20.0,
			},{
				itemID: 3,
				nome: "Borracha",
				valor: 30.0,
			},{
				itemID: 4,
				nome: "Apagador",
				valor: 40.0,
			},
		},
	}

	fmt.Printf("Valor total do pedido: %f\n",tipoPedido.getValorTotalPedido())
}


type pessoa struct {
	nome string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

//perceba que para alterar algum valor da struct o receiver precisa ser um ponteiro
func (p *pessoa) setNomeSobrenome(nomeCompleto string) {
	nomes := strings.Split(nomeCompleto, " ")
	p.nome = nomes[0]
	p.sobrenome = nomes[1]
}

func testMetodoQueAlteraValorDaStruct() {
	var p1 = pessoa{"Wesley", "Freitas"}
	fmt.Println(p1.getNomeCompleto())
	p1.setNomeSobrenome("Wesley Ramalho")
	fmt.Println(p1.getNomeCompleto())
}

//interface
type imprimivel interface {
	toString() string
}

type carro struct {
	marca string
	modelo string
}

//vou usar também a struct pessoa que foi declarada mais acima

func (p pessoa) toString() string {
	return "Pessoa: " + p.nome + " " + p.sobrenome
}

func (c carro) toString() string {
	return "Carro: " + c.marca + " " + c.modelo
}

func imprimir(x imprimivel){
	fmt.Println(x.toString())
}

func testInterface(){
	//a coisa implementa a interface imprimvel
	var coisa imprimivel = pessoa{"Wesley", "Freitas"}
	fmt.Println(coisa.toString())

	//Olha o polimorfismo aí
	coisa = carro{"Toyota", "Corolla"}
	fmt.Println(coisa.toString())

	imprimir(coisa)
}

//struct to JSON (algo útil para quando for trabalhar com web api)

type produto2 struct {
	Nome string `json:"nome"` //Letra maiúscula é o que é exportado para fora do pacote
	Preco float64 `json:"preco"`
	Tags []string `json:"tags"`
} 

func testStructToJSON() {
	tipoProduto := produto2{"Celular", 14000.00, []string{"Roubo", "Apple"}}

	tipoProdutoJSON, _ := json.Marshal(tipoProduto) //isso retorna um slice de bytes, precisa converter para string
	fmt.Println(string(tipoProdutoJSON))
}

func testJSONToString() {
	var tipoProduto produto2
	jsonString := `{"nome":"Celular","preco":14000,"tags":["Roubo","Apple"]}`
	json.Unmarshal([]byte(jsonString), &tipoProduto) // precisa converter a string para slice de bytes para funcionar

	fmt.Println(tipoProduto)
}

//Em GO existe a função init que executa antes da main. Mas o ponto de entrada é a função main
func main() {
	testStrings()
	testArrayAndSlice()
	testMake()
	testMap()
	testFuncaoAnonima()
	testFuncaoComoParametro()
	testFuncaoVariatica()
	testReceiver()
	testStructAninhada()
	testMetodoQueAlteraValorDaStruct()
	testInterface()
	testStructToJSON()
	testJSONToString()
}