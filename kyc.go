package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "polygonid"
)

func main() {
    // Importamos la biblioteca de Polygon ID
    verifier := polygonid.Verifier{}

    // Obtenemos la dirección del usuario
    userAddress := "0x123456789012345678901234567890123456789"

    // Solicitamos un comprobante de identidad
    proofRequest := verifier.CreateProofRequest(userAddress)

    // Enviamos la solicitud al proveedor de KYC
    response, err := http.Post(proofRequest["url"], "application/json", json.NewEncoder().Encode(proofRequest["data"]))
    if err != nil {
        fmt.Println(err)
        return
    }

    // Verificamos la respuesta
    if response.StatusCode == http.StatusOK {
        // La solicitud se completó correctamente
        proof := json.NewDecoder(response.Body).Decode()

        // Verificamos la identidad del usuario
        if verifier.VerifyProof(proof) {
            // La identidad del usuario se ha verificado correctamente
            fmt.Println("La identidad del usuario se ha verificado correctamente")
        } else {
            // La identidad del usuario no se ha verificado correctamente
            fmt.Println("La identidad del usuario no se ha verificado correctamente")
        }
    } else {
        // Se produjo un error en la solicitud
        fmt.Println("Se produjo un error en la solicitud")
    }
}