package main

import (
    "fmt"
    "github.com/ethereum/go-ethereum/crypto"
    "time"
)

// Entrada es un struct que representa una entrada para un partido de fútbol.
type Entrada struct {
    ID       uint64
    Usuario  string
    Partido  string
    Asiento   int
    Fecha     time.Time
}

// Blockchain es un struct que representa una blockchain de entradas para partidos de fútbol.
type Blockchain struct {
    Cadena []Entrada
}

// NewBlockchain crea una nueva blockchain de entradas para partidos de fútbol.
func NewBlockchain() *Blockchain {
    return &Blockchain{
        Cadena: []Entrada{},
    }
}

// AddBlock añade un nuevo bloque a la blockchain.
func (bc *Blockchain) AddBlock(entrada Entrada) {
    // Generamos un hash para el bloque.
    hash := crypto.Sha256([]byte(entrada.ID.String() + entrada.Usuario + entrada.Partido + entrada.Asiento + entrada.Fecha.String()))

    // Creamos un nuevo bloque con el hash y los datos de la entrada.
    block := Entrada{
        ID:       bc.Cadena[len(bc.Cadena)-1].ID + 1,
        Usuario:  entrada.Usuario,
        Partido:  entrada.Partido,
        Asiento:  entrada.Asiento,
        Fecha:    entrada.Fecha,
    }
    block.Hash = hash

    // Añadimos el bloque a la blockchain.
    bc.Cadena = append(bc.Cadena, block)
}

// GetBlocks devuelve la lista de bloques de la blockchain.
func (bc *Blockchain) GetBlocks() []Entrada {
    return bc.Cadena
}

// main es el punto de entrada del programa.
func main() {
    // Creamos una nueva blockchain.
    bc := NewBlockchain()

    // Añadimos un bloque a la blockchain.
    bc.AddBlock(Entrada{
        ID:       1,
        Usuario:  "Cristian Betancur",
        Partido:  "Caracas FC vs Dvo. Tachira",
        Asiento:  10,
        Fecha:    time.Now(),
    })

    // Imprimimos la lista de bloques de la blockchain.
    for _, bloque := range bc.GetBlocks() {
        fmt.Println(bloque)
    }
}