package main

import (
	"fmt"
	"regexp"
)

func main() {

	var conteudoHTML = `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Titulo</title>
</head>

<body>
<img src='imagem1.jpg'>
<img src="imagem2.jpg">
</body>

</html>`

	imagens := pegaImagens(conteudoHTML)

	for indice, valor := range imagens {
		fmt.Println(indice, valor)
	}

}

func pegaImagens(textoHTML string) (retorno []string) {
	encontraImagem := regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	listaDeImagens := encontraImagem.FindAllStringSubmatch(textoHTML, -1)

	retorno = make([]string, len(listaDeImagens))

	for i := range listaDeImagens {
		retorno[i] = listaDeImagens[i][1]
	}
	return
}
