package main
import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
//    "strconv"
    "regexp"
    "strconv"
    "time"
    "crypto/md5"
    "io"
    "os"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    //注意:如果没有调用ParseForm方法, 下面无法获取表单数据
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])

    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("value:", strings.Join(v, ""))
    }

    fmt.Fprintf(w, "Hello mushroom")

}


func appdeploy(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    //注意:如果没有调用ParseForm方法, 下面无法获取表单数据

    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))
        //token := fmt.Sprintf("%x", h.Sum(strconv.FormatInt(crutime, 10)))

        fmt.Println(token)

        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, token)
        //t.Execute(w, nil)
    } else {

        if len(r.Form["password"][0]) < 5 {
            fmt.Fprintf(w, "password too short\n")
        }

        /*if _, err := strconv.Atoi(r.Form.Get("age")); err != nil {
            fmt.Fprintf(w, "Age must be number!")
        }*/

        if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("username")); !m {
            fmt.Fprintf(w, "username must be English\n")
        }

        if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
            fmt.Fprintf(w, "Age must be number!")
        }

        if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
            fmt.Fprintf(w, "realname must be Chinese!\n")
        }

        if m, _ := regexp.MatchString(`^([\w\.\_]{2, 10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
            fmt.Fprintf(w, "email must be email!\n")
        }
        //注意:正则表达式""和``的区别

        slice := []string{"apple", "pear", "banana"}
        for _, v := range slice {
            if v == r.Form.Get("fruit") {
                fmt.Fprintln(w, v)
            }
        }

        gslice := []int{1, 2}
        for _, v := range gslice {
            if strconv.Itoa(v) == r.Form.Get("gender") {
                fmt.Fprintln(w, v)
            }
        }

        /*islice := []string{"football", "basketball", "tennis"}
        a := Slice_diff(r.Form["interest"], slice)
        if a == nil {
            fmt.Fprintln(w, a)
        }*/
        //注意:Slice_diff 来自库github.com/astaxie/beeku

        token := r.Form.Get("token")
        if token != "abc" {
            fmt.Printf("token: %s\n", token)
        } else {
            fmt.Println("token not existed!")
        }


        fmt.Println("username:", r.Form["username"])
        fmt.Println("username length:", len(r.Form["username"][0]))
        //fmt.Println("password:", r.Form["password"])
        //fmt.Println("usernaem:", template.HTMLEscapeString(r.Form.Get("username")))
        fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
        fmt.Println("age:", r.Form["age"])
        fmt.Println("realname:", r.Form["realname"])
        fmt.Println("email:", r.Form["email"])
    }
}

func applist(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
         
    }
}

func main() {
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/applist", applist)
    http.HandleFunc("/appdeploy", appdeploy)

    err := http.ListenAndServe(":10000", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
