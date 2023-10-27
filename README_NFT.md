# fulgol
Proyecto Blockchain &amp; Web3.0

Este código funciona de la siguiente manera:

El struct Entrada representa una entrada para un partido de fútbol. Tiene los mismos campos que el código anterior.

El struct NFT representa un NFT de una entrada para un partido de fútbol. Tiene los siguientes campos:

```
Entrada: La entrada a la que representa el NFT.
Hash: El hash de la entrada.
QRCode: El código QR del NFT.
```

La función "NewNFT()" crea un nuevo NFT de una entrada para un partido de fútbol. Para ello, genera un hash para la entrada y lo almacena en el campo Hash del NFT. Luego, genera un código QR con los datos del NFT y lo almacena en el campo QRCode del NFT.

Para usar este código, podemos crear una nueva instancia de la clase Entrada y llamar a la función "NewNFT()" para crear un NFT de la entrada.

Por ejemplo, el siguiente código crearía un NFT de una entrada para un partido de fútbol:

```
Go
entrada := Entrada{
    ID:       1,
    Usuario:  "Juan Pérez",
    Partido:  "Real Madrid vs Barcelona",
    Asiento:  10,
    Fecha:    time.Now(),
}

nft := NewNFT(entrada)
```


Para obtener el código QR del NFT, podemos llamar al campo QRCode del NFT:

```
Go
qrcode := nft.QRCode
```

Este código es un punto de partida para crear NFT de entradas para partidos de fútbol. A medida que el proyecto FULGOL avance debemos hacerlo más robusto, podríamos agregar las siguientes características:

Almacenar los NFT en una blockchain para que sean más seguros y verificables.
Agregar una capa de tokenización para que los NFT puedan ser comprados y vendidos.
Agregar características adicionales a los NFT, como imágenes o vídeos del
