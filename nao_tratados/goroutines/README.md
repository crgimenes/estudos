# Goroutines

Go tem um sistema de multitarefa leve e bastante poderoso e simples de usar. Basicamente você só precisa usar a clausula go na frente da função.

Exemplo:

```go
go nomeDaRotina()
```

No exemplo abaixo demostramos que Go termina todas as threads quando main termina.
Para ver o efeito remova a linha com o timer dentro de main.

```go
package main

import "time"

func foo() {
	println("foo")
}

func main() {
	go foo()

	// experimente comentar a linha abaixo
	time.Sleep(2 * time.Second)
	println("fim")
}
```
[Playground](https://play.golang.org/p/4NYROfW7Ub)

### Canais
Agora vamos ver um exemplo mais complexo e aprender um pouco sobre canais.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printValue(value int, c chan int) {
	defer fmt.Printf("%v terminou\r\n", value)
	rt := rand.Int31n(1000)
	time.Sleep(time.Duration(rt) * time.Millisecond)
	c <- value
}

func main() {
	rand.Seed(time.Now().Unix())

	c := make(chan int)

	for i := 0; i < 10; i++ {
		go printValue(i, c)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("main: %vº colocado, valor = %v\r\n", i+1, <-c)
	}

}
```
[Playground](https://play.golang.org/p/kUgYknGoeN)

[Neste vídeo](https://www.youtube.com/watch?v=H2jULD66BII) eu mostro esse exemplo funcionando.

[Neste artigo](http://divan.github.io/posts/go_concurrency_visualize/) você vai ver uma ótima representação gráfica e muitos exemplos de Goroutines e canais.
