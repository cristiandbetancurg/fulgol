package main

import (
    "fmt"
    "github.com/ethereum/go-ethereum/crypto"
    "time"
    "github.com/skip2/go-qrcode"
)

// Entrada es un struct que representa una entrada para un partido de fútbol.
type Entrada struct {
    ID       uint64
    Usuario  string
    Partido  string
    Asiento   int
    Fecha     time.Time
}

// NFT es un struct que representa un NFT de una entrada para un partido de fútbol.
type NFT struct {
    Entrada  Entrada
    Hash     string
    QRCode   []byte
}

// NewNFT crea un nuevo NFT de una entrada para un partido de fútbol.
func NewNFT(entrada Entrada) *NFT {
    // Generamos un hash para la entrada.
    hash := crypto.Sha256([]byte(entrada.ID.String() + entrada.Usuario + entrada.Partido + entrada.Asiento + entrada.Fecha.String()))

    // Creamos un nuevo NFT con el hash y los datos de la entrada.
    nft := NFT{
        Entrada:  entrada,
        Hash:     hash,
    }

    // Generamos un código QR con los datos del NFT.
    qrcode, err := qrcode.Encode(nft.Hash, qrcode.Medium, qrcode.Moire)
    if err != nil {
        panic(err)
    }

    nft.QRCode = qrcode

    return &nft
}

// main es el punto de entrada del programa.
func main() {
    // Creamos una nueva entrada para un partido de fútbol.
    entrada := Entrada{
        ID:       1,
        Usuario:  "Marlin Luna",
        Partido:  "Caracas FC vs Dvo. Tachira",
        Asiento:  12,
        Fecha:    time.Now(),
    }

    // Creamos un NFT de la entrada.
    nft := NewNFT(entrada)

    // Imprimimos el código QR del NFT.
    fmt.Println(nft.QRCode)
}