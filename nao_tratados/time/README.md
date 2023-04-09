# time

Go tem uma forma... peculiar de formatar datas, e é simplesmente terrível!

Você usa uma data específica para formatar datas no lugar do tradicional "dd/MM/yyyy" você usa "02/01/2006" isso parece altamente ilógico... se você é americano você pode lembrar de uma sequencia simples como  "01/02 03:04:05PM '06 -0700", mas não é nada intuitivo para o resto do planeta. Fazer o que né tinha que ter um WTF!

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("t é uma variavel tipo %T\r\n",t)
	fmt.Println((t.Format(time.RFC3339)))

	// só pode ser piada! "01/02 03:04:05 PM 2006 -07:00 MST" sério?
	// MST é Mountain Standard Time

	fmt.Println(t.Format("2006-01-02 15:04:05-0700"))
}
```
[Playground](https://play.golang.org/p/gBbxN8DLVL)

Felizmente essa ca... vai servir para aprender um conceito importante de Go.

**Leia os fontes!**

https://golang.org/src/time/format.go


Existem também muitas implementações de strftime para mascarar esse problema e formatar datas de um modo mais tradicional.

https://godoc.org/?q=strftime

---

Uma coisa interessante é usar time junto com channels para criar um timer.

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		// espera pelo sinal
		<-sc

		fmt.Printf("\r\nliberando recursos...\r\n")

		// mostra o cursor novamente
		fmt.Print("\033[?25h")

		fmt.Printf("have a nice day!\r\n")
		os.Exit(0)
	}()

	// esconde o cursor
	fmt.Print("\033[?25l")

	timer := time.Tick(time.Duration(300) * time.Millisecond)
	//◐◓◑◒
	//▏▎▍▋
	s := []rune(`◐◓◑◒`)
	slen := len(s)
	i := 0
	for {
		<-timer
		fmt.Print("\r")
		fmt.Print(string(rune(s[i])))
		i++
		if i == slen {
			i = 0
		}
	}
}
```
