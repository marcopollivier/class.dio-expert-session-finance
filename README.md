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
    - `$ go get -u golang.org/x/lint/golint`
        
    10.2. E vamos executar nossos primeiros comandos relacionados 
    - `$ go test ./...`
    - `$ golint ./...`

11. Vamos configurar o CircleCI 

12. Vamos colocar métricas na nossa aplicação 

    ```
    $ go get github.com/prometheus/client_golang/prometheus
    $ go get github.com/prometheus/client_golang/prometheus/promauto
    $ go get github.com/prometheus/client_golang/prometheus/promhttp 
    ```

13. Para os passos seguintes, nós vamos fazer uma integração com um BD qualquer. 
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

14. Já com o banco acessível via Docker, vamos criar a base que utilizaremos no nosso teste
   
    ```sql
    CREATE TABLE transactions (
        id SERIAL PRIMARY KEY,
        title varchar(100),
        amount decimal,
        type smallint,
        installment smallint,
        created_at timestamp 
    );
    
    insert into transactions (title, amount, type, installment, created_at)
    values ('Freela', '100.0', 0, 1, '2020-04-10 04:05:06'); 
   
    select * from transactions;
    ```

15. Agora com a estrutra de banco criada, vamos fazer as alterações necessárias no código. E a primeira delas é baixar a
dependência do driver do Postgres. 

    [Lista de SQLDrivers disponíveis](https://github.com/golang/go/wiki/SQLDrivers) 

    Execute o seguinte comando dentro da pasta do projeto 
    
    ```shell script
    $ go get -u github.com/lib/pq
    ```

16. E esse é o código que vai manipular as informações do banco de fato

    ```go
    package postgres
    
    import (
    	"database/sql"
    	"fmt"
    	"github.com/marcopollivier/dio-expert-session-pre-class/model/transaction"
    
    	_ "github.com/lib/pq"
    )
    
    const (
    	host     = "localhost"
    	port     = 5432
    	user     = "postgres"
    	password = "postgres"
    	dbname   = "diodb"
    )
    
    func connect() *sql.DB {
    	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    	var db, _ = sql.Open("postgres", psqlInfo)
    	return db
    }
    
    func Create(transaction transaction.Transaction) int {
    	var db = connect()
    	defer db.Close()
    
    	var sqlStatement = `INSERT INTO  transactions (title, amount, type, installment, created_at)
    						VALUES ($1, $2, $3, $4, $5)
    						RETURNING id;`
    
    	var id int
    	_ = db.QueryRow(sqlStatement,
    					transaction.Title,
    					transaction.Amount,
    					transaction.Type,
    					transaction.Installment,
    					transaction.CreatedAt).Scan(&id)
    	fmt.Println("New record ID is:", id)
    
    	return id
    }
      
    func main() {
    	//log.Fatal(http.Init())
    
    	//"Outro freela", 400.0, 0, 1, "2020-04-20 12:00:06"
    	postgres.Create(transaction.Transaction{Title: "Outro freela", Amount: 600.0, Type: 0, Installment: 1, CreatedAt: util.StringToTime("2020-04-20T12:00:06")})
    }
    
    func FetchAll() transaction.Transactions {
    	var db = connect()
    	defer db.Close()
    
    	rows, _ := db.Query("SELECT title, amount, type, installment, created_at FROM transactions")
    	defer rows.Close()
    
    	var transactionSlice []transaction.Transaction
    	for rows.Next() {
    		var transaction transaction.Transaction
    		_ = rows.Scan(&transaction.Title,
    					  &transaction.Amount,
    					  &transaction.Type,
    					  &transaction.Installment,
    					  &transaction.CreatedAt)
    
    		transactionSlice = append(transactionSlice, transaction)
    	}
    
    	return transactionSlice
    }
    
    fmt.Print(postgres.FetchAll())
    ```
    
