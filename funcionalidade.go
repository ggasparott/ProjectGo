// Declaração Explicita COM VAR	
var nome string = "Gabriel"
var idade int = 25

// Declaração Implicita COM :=
cidade := "São Paulo"
altura := 1.75

// Multiplas declarações
var (
    nome  = "Gabriel"
    idade = 25
    nota  = 9.5
)

// Tipos de Dados
Números Inteiros: int, int8, int16, int32, int64
Números Decimais: float32, float64
Texto: string
Booleanos: bool (true ou false)

// Ler uma dados
fmt.Scanln(&string) // Scanln para ler strings

fmt.Scan(&int, &float32 ou float64) // Scan para ler int ou float

// Condicionais If else, else if, else
if  {
    // não precisa dos parenteses
} else {
    }


// Laços de repetição

for inicialização; condição; incremento {
    // código a ser repetido
}

// while 
    for numero <= 0 {
        fmt.Print("Número inválido! Digite novamente: ")
        fmt.Scan(&numero)
    }




