# Go Web Programming

- Online Video (content purely in Chinese): https://www.bilibili.com/video/BV1Xv411k7Xn
- Another Github Repo (in Chinese): https://github.com/xpengkang/go-web-zero

### Part 0 Server

1. `func ListenAndServe(addr string, handler Handler) error`
   - server url + handler
   - default: `"http://localhost:80"` and `nil`
   - if the value of handler is `nil`, then `DefaultServeMux` will be used as the handler
2. We can also write a new `Server` struct with arguments `Addr` and `Handler`，and then call the `ListenAndServe()` method on the struct

### Part 1 Handler

1. the `Handler` interface
   - the `Handler` interface has the method `ServeHTTP`
   - ```
     type Handler interface {
         ServeHTTP(ResponseWriter, *Request)
     }
     ```
2. `http.Handle()` vs `http.HandleFunc()`
   1. `http.Handle(pattern string, handler Handler)`
   2. `func http.HandleFunc(pattern string, handler func(ResponseWriter, *Request))`
   3. the `HandlerFunc` type that implements `Handler`
     - `type HandlerFunc func(ResponseWriter, *Request)`
     - this type implements the method `func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)`
   4. summary
     - Both `http.Handler()` and `http.HandleFunc()` can be used to register a `handler`. The former one registers a `Handler` type while the latter one registers a function
     - We can also use `http.HandlerFunc()` to convert a `handleFunc` to a `Handler` type
3. The Built-in `Handler`'s
   - `func NotFoundHandler() Handler`
   - `func RedirectHandler(url string, code int) Handler`
   - `func StripPrefix(prefix string, h handler) Handler`
   - `func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler`
   - `func FileServer(root FileSystem) Handler`
     - `http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot"))` that allows the server to access the files under the local path of `wwwroot`
     - `type Dir string` implements the `Open` method, so it is a `FileSystem` type
       - `func (d Dir) Open(name string) (File, error)`

### Part 2 Request

1. Request vs Response
2. `r.URL`
   1. The general pattern: `scheme://[userinfo@]host/path[?query][#fragment]`
   2. `query`/`RawQuery` means the parameters after the url, e.g., `id=123&thread_id=456`
   3. `fragment` will be removed if the request is sent from the browser

#### Read the Request from Backend

1. Get the header of `r *Request`
   1. `r.Header["Accept-Encoding"]` returns `[]string` type
   2. `r.Header.Get("Accept-Encoding")` returns `string` type
2. Get the body of `r *Request`: define a buffer and read
```go
http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
    length := r.ContentLength
    body := make([]byte, length)
    r.Body.Read(body)
    fmt.Fprintln(w, string(body))
})
```
3. Get the parameters of `r *Request`
   - `r.URL.RawQuery` returns the raw string of `query`
   - `r.URL.Query()` returns a `map[string][]string` type
```go
url := r.URL
query := url.Query()
id := query["id"] // []string{"123"} (a slice)
threadID := query.Get("thread_id") // "456" (the string itself)
```

#### Form Data

1. Submit a form
   - `enctype`: `application/x-www-form-urlencoded` (default) or `multipart/form-data`
   - `multipart/form-data` converts every `name-value` pair to a MIME part where each part has its own `Content Type` and `Content Disposition`
   - We can also use GET method to submit a form, where such requests do not have bodies
2. Read a form. Attributes can come from both the form fields and url params
   - `r.ParseForm()`
   - `r.Form` returns both form fields and url params
   - `r.PostForm` returns only form fields
3. Multipart
   - `r.ParseMultipartForm()`
   - `r.FormValue("first_name")` returns both form fields and url params
   - `r.PostFormValue("first_name")` returns only form fields
4. `MultipartReader()` processes form data in a way of streaming
   - `func (r *Request) MultipartReader() (*multipart.Reader, error)`

#### Upload a file

```go
func process(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(1024)

    fileHeader := r.MultipartForm.File["uploaded"][0]
    file, err := fileHeader.Open()
    if err == nil {
        data, err := ioutil.ReadAll(file)
        if err == nil {
            fmt.Fprintln(w, string(data))
        }
    }
}
```

```go
func process(w http.ResponseWriter, r *http.Request) {
    // File, FileHeader, Error
    // the first file of all files uploaded
    file, _, err := r.FormFile("uploaded")
    if err == nil {
        data, err := ioutil.ReadAll(file)
        if err == nil {
            fmt.Fprintln(w, string(data))
        }
    }
}
```

### Part 3 Response

1. `ResponseWriter` writes to the response body
2. `WriteHeader` writes to the response header
3. `Header` returns a mutable map of the headers

```go
func writeExample(w http.ResponseWriter, r *http.Request) {
    str := `<html>
<head><title>Go Web</title></head>
<body><h1>Hello World</h1></body>
</html>
`
    w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(501)
    fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Location", "http://google.com")
    w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    post := &Post{
        User: "Sau Sheong",
        Threads: []string{"first", "second", "third"}
    }
    json, _ := json.Marshal(post)
    w.Write(json)
}

func main() {
    server := http.Server{
        Addr: "localhost:8080",
    }

    http.HandleFunc("/write", writeExample)
    http.HandleFunc("/writeheader", writeHeaderExample)
    http.HandleFunc("/redirect", headerExample)
    http.HandleFunc("/json", jsonExample)

    server.ListenAndServe()
}
```

4. Built-in Responses (functions)
    - `NotFound()`
    - `ServeFile()`
    - `ServeContent()`
    - `Redirect()`

### Part 4 Template

1. Logic-less vs. Logic
2. Usage
   - Parse
     - `ParseFiles()`, `ParseGlob()`, `Parse()`
     - Only returns the first one if multiple exist
   - Execute the parsed template and pass ResponseWriter and data
     - `Must()` function
     - `Execute()`, `ExecuteTemplate()`
3. Example

```go
package main

import (
	"net/http"
	"text/template"
)

func main() {
	server := http.Server{
		Addr := "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	// "Hello World!" will be passed
	// to the "action" of the template
	t.Execute(w, "Hello World!")
}
```

4. An advanced example

```go

func main() {
    templates := loadTemplates()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fileName := r.URL.Path[1:]
        t := template.Lookup(fileName)
        if t != nil {
            err := t.Execute(w, nil)
            if err != nil {
                log.Fatalln(err.Error())
            }
        } else {
            w.WriteHeader(http.StatusNotFound)
        }
    })

    http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))
    http.Handle("/img/", http.FileServer(http.Dir("wwwroot")))

    http.ListenAndServe("localhost:8080", nil)
}

func loadTemplates() *template.Template {
    result := template.New("templates")
    // result, err := result.ParseGlob("templates/*.html")
    // template.Must(result, err)
    template.Must(result.ParseGlob("templates/*.html"))
    return result
}
```

5. Action

- `{{ if . }}`, `{{ else }}`, `{{ end }}` (condition)
- `{{ range . }}`, `{{ else }}`, `{{ end }}` (iteration)
- `{{ with "world" }}`, `{{ else }}`, `{{ end }}` (set)
- `{{ template "t2.html" }}` (include another template)

6. Functions and Pipes

- Set new variables in the Action
  - `$key`, `$value`
- pipe
  - `{{ p1 | p2 | p3 }}`
  - `{{ 12.3456 | printf "%.2f" }}`
- function
  - built-in functions
    - define, template, block
    - html, js, urlquery
    - index
    - print/printf/println
    - len
    - with, set all the local `.` as the value following the `with` keyword
- example: `template.New("").Funcs(funcMap).Parse(...)`

```go
func main() {
    server := http.Server{
        Addr: "localhost:8080"
    }
    http.HandleFunc("/process", process)
    server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
    funcMap := template.FuncMap{"fdate": formatDate}
    t := template.New("t1.html").Funcs(funcMap)
    t.ParseFiles("t1.html")
    t.Execute(w, time.Now())
}

func formatDate(t time.Time) string {
    layout := "2006-01-02"
    return t.Format(layout)
}
```

7. `Layout` template

- `{{ define "content" }}`, `{{ end }}`
- `{{ template "content" . }}`
- `{{ block "content" . }}`, `{{ end }}`

8. Logical operators

- `eq`, `ne`
- `lt`, `gt`
- `le`, `ge`
- `and`
- `or`
- `not`

### Part 5 Database

1. Connection

```go

var db *sql.DB

func main() {
    connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
        server, user, password, port, database)
    
    // driver and connStr, "sqlserver" as an example
    var err error
    db, err = sql.Open("sqlserver", connStr)
    if err != nil {
        log.Fatalln(err.Error())
    }
    ctx := context.Background()

    err = db.PingContext(ctx)
    if err != nil {
        log.Fatalln(err.Error())
    }

    fmt.Println("Connected!")
}
```

2. Multi-row: `Query`

Queries: `Query`, `QueryRow`, `QueryContext`, `QueryRowContext`

https://pkg.go.dev/database/sql#Rows

- `func (rs *Rows) Close() error`, close
- `func (rs *Rows) ColumnTypes() ([]*ColumnType, error)`, types of columns
- `func (rs *Rows) Columns() ([]string, error)`, names of columns
- `func (rs *Rows) Err() error`, error
- `func (rs *Rows) Next() bool`, iterates
- `func (rs *Rows) NextResultSet() bool`, iterates over the result set
- `func (rs *Rows) Scan(dest ...interface{}) error`, copy and set to the dest

3. Single-row: `QueryRow`

https://pkg.go.dev/database/sql#Row

- `func (r *Row) Err() error`, error
- `func (r *Row) Scan(dest ...interface{}) error`, copy and set to the dest

4. Example

Folder structure

- `models.go` where the entity structs are put

```go
type app struct {
    ID     int
    name   string
    status int
    level  int
    order  int
}
```

- `services.go`

```go
// func getOne(id int) (app, error) {
//     a := app{}
//     err := db.QueryRow("SELECT Id, Name, Status, Level, Order FROM dbo.App").Scan(
//         &a.ID, &a.name, &a.status, &a.level, &a.order)
//     return a, err
// }

func getOne(id int) (a app, err error) {
    a = app{}
    err = db.QueryRow("SELECT Id, Name, Status, Level, [Order] FROM dbo.App WHERE Id=@Id",
        sql.Named("Id", id)).Scan(
        &a.ID, &a.name, &a.status, &a.level, &a.order)
    return
}

func getMany(id int) (apps []app, err error) {
    rows, err := db.Query("SELECT Id, Name, Status, Level, [Order] FROM dbo.App WHERE Id>@Id",
        sql.Named("Id", id))
    for rows.Next() {
        a := app{}
        err = rows.Scan(&a.ID, &a.name, &a.status, &a.level, &a.order)
        if err != nil {
            log.Fatalln(err.Error())
        }
        apps = append(apps, a)
    }
    return
}
```

5. Update and Delete

```go
func (a *app) Update() (err error) {
    // result, error
    _, err = db.Exec("UPDATE dbo.App SET Name=@Name, [Order]=@Order WHERE Id=@Id",
        sql.Named("Name", a.name), sql.Named("Order", a.order), sql.Named("Id", a.ID))
    if err != nil {
        log.Fatalln(err.Error())
    }
    return
}

func (a *app) Delete() (err error) {
    // result, error
    _, err = db.Exec("DELETE FROM dbo.App WHERE Id=@Id", sql.Named("Id", a.ID))
    if err != nil {
        log.Fatalln(err.Error())
    }
    return
}
```

6. `Ping`, `PingContext`, `Prepare`, `PrepareContext`, `Transactions` (`Begin`, `BeginTx`)

```go
func (a *app) Insert() (err error) {
    statement := `INSERT INTO dbo.App (Name, Nickname, Status, Level, [Order], Pinyin) VALUES (@Name, 'Nick', @Status, @Level, @Order, '...');
        SELECT isNull(SCOPE_IDENTITY(), -1);`
    stmt, err := db.Prepare(statement)
    if err != nil {
        log.Fatalln(err.Error())
    }
    defer stmt.Close()
    err = stmt.QueryRow(
        sql.Named("Name", a.name),
        sql.Named("Status", a.status),
        sql.Named("Level", a.level),
        sql.Named("Order", a.order)
    ).Scan(&a.ID)
    
    if err != nil {
        log.Fatalln(err.Error())
    }

    return
}
```

### Part 6 Controller

1. Folder structure
   - controller
     - controller.go
       - ```go
         package controller

         // RegisterRoutes ...
         func registerRoutes() {
             // static resources

             registerHomeRoutes()
             registerAboutRoutes()
             registerContactRoutes()
         }
         ```
     - home.go
       - ```go
         package controller

         func registerHomeRoutes() {
             http.HandleFunc("/home", handleHome)
         }

         func handleHome(w http.ResponseWriter, r *http.Request) {
             // handler
         }
         ```
     - about.go
       - ```go
         package controller

         func registerAboutRoutes() {
             http.HandleFunc("/about", handleAbout)
         }

         func handleAbout(w http.ResponseWriter, r *http.Request) {
             // handler
         }
         ```
     - contact.go
       - ```go
         package controller

         func registerContactRoutes() {
             http.HandleFunc("/contact", handleContact)
         }

         func handleContact(w http.ResponseWriter, r *http.Request) {
             // handler
         }
         ```

2. Routes with params

```go
func registerCompanyRoutes() {
    http.HandleFunc("/companies", handleCompanies)
    http.HandleFunc("/companies/", handleCompany)
}

func handleCompanies(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("layout.html", "companies.html")
    t.ExecuteTemplate(w, "layout", nil)
}

func handleCompany(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("layout.html", "company.html")

    pattern, _ := regexp.Compile(`/companies/(\d+)`)
    matches := pattern.FindStringSubmatch(r.URL.path)
    if len(matches) > 0 {
        companyID, _ := strconv.Atoi(matches[i])
        t.ExecuteTemplate(w, "layout", companyID)
    } else {
        w.WriteHeader(http.StatusNotFound)
    }
}
```

3. Third-part router / multiplexer

- `gorilla/mux`
- `httprouter`


### Part 7 JSON

1. Mapping from Go struct to JSON body

```go
type Company struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Country string `json:"country"`
}
```

2. Data types

- Go bool -> JSON boolean
- Go float64 -> JSON numerics
- Go string -> JSON strings
- Go nil -> JSON null

3. JSON in an unknown/dynamic structure

- `map[string]interfaces{}`, a JSON
- `[]interface{}`, an JSON array

4. read & write
    1. encoder and decoder (io stream)
        1. decoder
            - `dec := json.NewDecoder(r.Body)`
            - `dec.Decode(&query)`
        2. encoder
            - `enc := json.NewEncoder(w)`
            - `enc.Encode(results)`
    2. marshal and unmarshal (string/bytes)
```go
func main() {
    jsonStr := `
    {
        "id": 123,
        "name": "Google",
        "country": "USA"
    }`

    c := Company{}
    _ = json.Unmarshal([]byte(jsonStr), &c)
    fmt.Println(c)

    // json without indent
    // e.g., {"id":123,"name":"Google","country":"USA"} 
    bytes, _ := json.Marshal(c)
    fmt.Println(string(bytes))

    // data, prefix, indent
    // json with indent
    // e.g.
    // {
    //   "id": 123,
    //   "name": "Google",
    //   "country": "USA"
    // } 
    bytes1, _ := json.MarshalIndent(c, "", "  ")
    fmt.Println(string(bytes1))
}
```

5. example

```go
func main() {
    http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case http.MethodPost:
                dec := json.NewDecoder(r.Body)
                company := Company{}
                err := dec.Decode(&company)
                if err != nil {
                    log.Println(err.Error())
                    w.WriteHeader(http.StatusInternalServerError)
                    return
                }

                enc := json.NewEncoder(w)
                err = enc.Encode(company)
                if err != nil {
                    log.Println(err.Error())
                    w.WriteHeader(http.StatusInternalServerError)
                    return
                }
            default:
                w.WriteHeader(http.StatusMethodNotAllowed)
        }
    })
}
```

### Part 8 Middleware

1. example

```go
type MyMiddleware struct {
    Next http.Handler
}

func (m MyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // do something before the Next handler
    m.NextServeHttp(w, r)
    // do something after the Next handler
}
```

2. Possible usage

- `Logging`
- Security
- Timeout
- Compress the response

3. An `AuthMiddleware` example

```go
package middleware

import "net/http"

// AuthMiddleware ...
type AuthMiddleware struct {
    Next http.Handler
}

func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if am.Next == nil {
        am.Next = http.DefaultServeMux
    }

    auth := r.Header.Get("Authorization")
    if auth != "" {
        // requires the request to have an Authorization header
        am.Next.ServeHTTP(w, r)
    } else {
        w.WriteHeader(http.StatusUnauthorized)
    }
}
```

### Part 9 Request Context

1. Request Context

- `func (*Request) Context() context.Context`
- `func (*Request) WithContext(ctx context.Context) context.Context`


2. `context.Context` (read only)

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

3. Context APIs

- `WithCancel()` with a `CancelFunc`
- `WithDeadline()` with a timestamp (`time.Time`)
- `WithTimeout()` with a duration (`time.Duration`)
- `WithValue()` with some values

4. Example

```go
type TimeoutMiddleware struct {
    Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if tm.Next == nil {
        tm.Next = http.DefaultServeMux
    }

    // create a context
    ctx := r.Context()

    // set timeout
    ctx, _ = context.WithTimeout(ctx, 3*time.Second)
    r.WithContext(ctx)
    ch := make(chan struct{})
    go func() {
        tm.Next.ServeHTTP(w, r)
        // if the request is finished within 3s
        // a new struct will be sent to the channel
        ch <- struct{}()
    }()
    select {
    case <-ch:
        // the goroutine finished in time
        return
    case <-ctx.Done():
        // the goroutine didn't finish in time
        // ctx.Done() is called
        // timeout
        w.WriteHeader(http.StatusRequestTimeout)
    }
    ctx.Done()
}
```

### Part 10 HTTPS

1. Generate a new certificate
   - `go run ~/go/src/crypto/tls/generate_cert.go -h`, help
   - `go run ~/go/src/crypto/tls/generate_cert.go -host localhost`, generate a certificate
2. `http.ListenAndServeTLS`

```go
func main() {
    controller.RegisterRoutes()
    http.ListenAndServeTLS("localhost:8080", "cert.pem", "key.pem", nil)
}
```

3. Server Push

```go
func registerHomeRoutes() {
    http.HandleFunc("/home", handleHome)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    // server push
    // push "/css/app.css"
    if pusher, ok := w.(http.Pusher); ok {
        pusher.Push("/css/app.css", &http.PushOptions{
            Header : http.Header {"Content-Type": []string{"text/css"}},
        })
    }
    t, _ := template.ParseFiles("layout.html", "home.html")
    t.ExecuteTemplate(w, "layout", "Hello World")
}
```

### Part 11 Tests

- `user_test.go`
  - The filename of a test code should end with `test`
  - Production: do not include files of test code
  - Development: include files of test code
- `func TestUpdatesModifiedTime(t *testing.T){ ... }`
  - The test function should start with `Test`
  - The naming of the function should display its usage
  - The arg type, `*testing.T`, provides some relevant tools
- Example to test a model with `go test {package_name}`
  - `go test request/model`
  - `go test request/controller`

`company.go`

```go
package model

import "strings"

// Company ...
type Company struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Country string `json:"country"`
}

// GetCompanyType ...
func (c *Company) GetCompanyType() (result string) {
    if strings.HasSuffix(c.Name, ".LTD") {
        result = "Limited Liability Company"
    } else {
        result = "Others"
    }
    return
}
```

`company_test.go`

```go
package model

import "testing"

func TestCompanyTypeCorrect (t *testing.T) {
    c := Company{
        ID: 12345,
        Name: "ABCD .LTD",
        Country: "China",
    }

    companyType := c.GetCompanyType()

    if companyType != "Limited Liability Company" {
        t.Errorf("Company's GetCompanyType Method failed to get the correct company type")
    }
}
```

- Example to test a controller

`company.go`

```go
package controller

import (
    "net/http"
    "request/model"
)

// RegisterRoutes ...
func RegisterRoutes() {
    http.HandleFunc("/companies", handleCompany)
}

func handleCompany(w http.ResponseWriter, r *http.Request) {
    c := model.Company{
        ID: 123,
        Name: "Google",
        Country: "USA",
    }

    enc := json.NewEncoder(w)
    enc.Encode(c)
}

```

`company_test.go`

```go
package controller

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "request/model"
    "testing"
)

func TestHandleCompanyCorrect(t *testing.T) {
    r := httptest.NewRequest(http.MethodGet, "/companies", nil)
    w := httptest.NewRecorder()

    handleCompany(w, r)

    result, _ := ioutil.ReadAll(w.Result().Body)

    c := model.Company{}
    json.Unmarshal(result, &c)

    if c.ID != 123 {
        t.Errorf("Failed to handle company correctly!")
    }
}
```

### Part 12 Profiling

- Memory usage
- CPU usage
- Blocked goroutines
- Tracing
- Web UI

Usage:
- `import _ "net/http/pprof"`
- Then start a new goroutine and use a new port to start another server for profiling, say "localhost:8000".
- Go to "localhost:8000/debug/pprof" for the Web UI of profiling.
- Or use CLI tools
  - `go tool pprof http://localhost:8000/debug/pprof/heap`, memory
  - `go tool pprof http://localhost:8000/debug/pprof/profile`, CPU
  - `go tool pprof http://localhost:8000/debug/pprof/block`, goroutines
  - `go tool pprof http://localhost:8000/debug/pprof/trace?seconds=5`, trace
