# dio-expert-session-pre-class

## Pré Desenvolvimento

1. Vamos criar um projeto no Github chamado `dio-expert-session-finance`
    - Depois voltamos aqui para configurar mais nosso repositório. Por enquanto vamos ter apenas arquivos básicos
        - `.gitignore`
        - `README.md`

2. Depois do nosso projeto criado no Github, vamos criar um projeto `Go` básico. 
    - `$ mkdir dio-expert-session-finance`
    - `$ cd dio-expert-session-finance`
    - `$ go mod init github.com/marcopollivier/dio-expert-session-finance`
    
    Esses comandos vão criar dentro da pasta `dio-expert-session-finance` um arquivo `go.mod`. 
    E esse arquivo vai ser a base do nosso projeto. 
    
    Depois disso, você já pode abrir seu projeto na IDE de sua escolha
    - GoLand 
    - VSCode 
    - Vim 
    
 3. Agora que temos o nosso projeto funcionando corretamente, vamos criar um `Hello, World!` para 
 termos certeza que tudo está de acordo com o que esperávamos. 
     - Dentro da pasta `dio-expert-session-finance/cmd/server/` vamos criar o arquivo `main.go`
     
    ```go
    package main
    
    import "fmt"
    
    func main() {
        fmt.Print("Olá, Mundo!")
    }    
    ```

Com isso já vemos que nosso ambiente está ok e funcional, mas ainda não é isso que queremos exatamente, correto? 

4. Mas precisamos evoluir nosso código para começar a tomar forma de uma API. 

    ```go
    package main
    
    import (
           "fmt"
           "net/http"
    )
    
    func main() {
           http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                   fmt.Fprintf(w, "Olá. Bem vindo a minha página!")
           })
    
           http.ListenAndServe(":8080", nil)
    }
    ```
    `$ curl curl http://localhost:8080/`
    
    
 
        
        
    
