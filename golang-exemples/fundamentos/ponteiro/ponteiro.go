package main

import "fmt"

func main() {

	// variavel i de tipo int32
	i := 1

	// go nao tem aritimetica de ponteiros ...

	// p é um ponteiro
	var p *int = nil
	var r **int = nil

	// p recebe o endereco da varivel i
	p = &i
	r = &p

	// endereco de memoria de i
	fmt.Println(p)

	// valor da variavel i
	fmt.Println(*p)

	// incrementa +1 pela referencia do ponteiro para a variavel i
	*p++

	fmt.Println(i)
	fmt.Println(*p)

	// incrementa +1 direto na variavel i
	i++

	fmt.Println(i)
	fmt.Println(*p)

	// ponteiro r que referencia ponteiro p que referencia variavel i

	fmt.Println(r)
	fmt.Println(*r)
	fmt.Println(**r)
}
