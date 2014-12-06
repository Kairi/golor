golor
=====

golang lib for colorlize ouput

**Instllation

```bash
go get github.com/Kairi/golor
```

**Sample
```go
import (
    "fmt"
    "github.com/Kairi/golor"
)

func main() {
    fmt.Println("basic string")
    fmt.Println(golor.ColorString("yellow string", golor.YELLOW))
    golor.ColorPrintln("blue string", golor.YELLOW)    
}
```
    


