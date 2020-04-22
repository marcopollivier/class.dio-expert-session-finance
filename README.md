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
    
5. Vamos começar a pensar no nosso modelo. Nós queremos criar um sistema pras nossas finanças pessoais. 
Para isso vamos pensar num modelo para trabalharmos. 

    5.1. Nosso modelo financeiro vai ser bem simples, mas flexível o suficiente para evoluirmos 
    no futuro
    
    - Titulo
    - Valor 
    - Tipo (ENTRADA, SAIDA)
    - Data
 
    5.2. Vamos pensar que, antes de mais nada, queremos retornar um JSON com esse modelo.
    
    ```go
    package main
    
    import (
    	"encoding/json"
    	"net/http"
    	"time"
    )
    
    func main() {
    	http.HandleFunc("/transactions", getTransactions)
    
    	_ = http.ListenAndServe(":8080", nil)
    }
    
    type Transaction struct {
    	Title     string
    	Amount    float32
    	Type      int //0. entrada 1. saida
    	CreatedAt time.Time
    }
    
    type Transactions []Transaction
    
    func getTransactions(w http.ResponseWriter, r *http.Request) {
    	if r.Method != "GET" {
    		w.WriteHeader(http.StatusMethodNotAllowed)
    		return
    	}
    
    	w.Header().Set("Content-type", "application/json")
    
    	layout := "2006-01-02T15:04:05"
    	salaryReceived, _ := time.Parse(layout, "2020-04-05T11:45:26")
    	paidElectricityBill, _ := time.Parse(layout, "2020-04-12T22:00:00")
    	var transactions = Transactions{
    		Transaction{
    			Title:     "Salário",
    			Amount:    1200.0,
    			Type:      0,
    			CreatedAt: salaryReceived,
    		},
    		Transaction{
    			Title:     "Conta de luz",
    			Amount:    100.0,
    			Type:      1,
    			CreatedAt: paidElectricityBill,
    		},
    	}
    
    	_ = json.NewEncoder(w).Encode(transactions)
    }
    ```

    ```go
    type Tction struct {
        Title     string    `json:"title"`
        Amount    float32   `json:"amount"`
        Type      int       `json:"type"` //0. entrada 1. saida
        CreatedAt time.Time `json:"created_at"`
    }
    ```

6. E agora vamos fazer um método de inserção (POST)
    
    ```go
    http.HandleFunc("/transactions/create", createATransaction)
   
    ...
   
    func createATransaction(w http.ResponseWriter, r *http.Request) {
    	if r.Method != "POST" {
    		w.WriteHeader(http.StatusMethodNotAllowed)
    		return
    	}
    
    	var res = Transactions{}
    	var body, _ = ioutil.ReadAll(r.Body)
    	_ = json.Unmarshal(body, &res)
    
    	fmt.Println(res)
    	fmt.Println(res[0].Title)
    	fmt.Println(res[1].Title)
    }
    ```
   
    ```json
    [
       {
           "title": "Salário",
           "amount": 1200,
           "type": 0,
           "created_at": "2020-04-05T11:45:26Z"
       }
    ]
    ```
   
   ```shell script
   $ curl -X POST 'http://localhost:8080/transactions/create' \
   -H 'Content-Type: application/json' \
   -d '[
           {
               "title": "Salário",
               "amount": 1200,
               "type": 0,
               "created_at": "2020-04-05T11:45:26Z"
           }
       ]'
   ```
   
7. Hora de refatorar. Vamos colocar em memória e isolar em arquivos e pacotes
 
8. Vamos começar a pensar em monitoramento? Então vamos criar um arquivo de Healthcheck

    ```go
    package actuator
    
    import (
        "encoding/json"
        "net/http"
    )
    
    func Health(responseWriter http.ResponseWriter, request *http.Request) {
        responseWriter.Header().Set("Content-Type", "application/json")
    
        profile := HealthBody{"alive"}
    
        returnBody, err := json.Marshal(profile)
        if err != nil {
            http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
            return
        }
    
        _, err = responseWriter.Write(returnBody)
        if err != nil {
            http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
            return
        }
    }
    
    type HealthBody struct {
        Status string
    }
    ``` 

9. Vamos aproveitar já que criamos um util e escrever um teste unitário para ele 
   
   ```go
   package util
   
   import (
   	"testing"
   )
   
   func TestStringToDate(testing *testing.T) {
   	var convertedTime = StringToTime("2019-02-12T10:00:00")
   
   	if convertedTime.Year() != 2019 {
   		testing.Errorf("Converter StringToDate is failed. Expected Year %v, got %v", 2019, convertedTime.Year())
   	}
   
   	if convertedTime.Month() != 2 {
   		testing.Errorf("Converter StringToDate is failed. Expected Month %v, got %v", 2, convertedTime.Month())
   	}
   
   	if convertedTime.Hour() != 10 {
   		testing.Errorf("Converter StringToDate is failed. Expected Hour %v, got %v", 10, convertedTime.Hour())
   	}
   
   }
   ```  
10. Vamos começar a pensar em um pouco de qualidade de código também

    10.1. Para fazer análise de código estática, vamos instalar a dependencia do lint 
        `$ go get -u golang.org/x/lint/golint`
        
    10.2. E vamos executar nossos primeiros comandos relacionados 
    	`$ go test ./...`
    	`$ golint ./...`

11. Para os passos seguintes, nós vamos fazer uma integração com um BD qualquer. 
Para isso, vamos subir uma imagem Docker do Postgres pra poder fazer o nosso teste. 

    Vamos subir o BD via Docker Compose
    
    host: localhost
    user: postgres
    pass: postgres
    DB: diodb
    
    ```yaml
    version: "3"
    services:
      postgres:
        image: postgres:9.6
        container_name: "postgres"
        environment:
          - POSTGRES_DB=diodb
          - POSTGRES_USER=postgres
          - TZ=GMT
        volumes:
          - "./data/postgres:/var/lib/postgresql/data"
        ports:
          - 5432:5432
    ```
   
    ```makefile
    prepare-tests:
       docker-compose -f .devops/postgres.yml up -d
    ```



