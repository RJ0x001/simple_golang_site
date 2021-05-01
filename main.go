package main

import ("fmt"
        "net/http"
        "html/template")

type User struct {
  Name string
  Age uint8
  Money int8
  Avg_grades, Happines float64
  Hobbies []string
}

func (u User) getAllInfo() string {
  return fmt.Sprintf("User: %s, age: %d, money: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string){
  u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request)  {
  bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Foo", "Bar"}}
  // fmt.Fprintf(w, )
  tmpl, _ := template.ParseFiles("templates/home_page.html")
  tmpl.Execute(w, bob)
}

func contacs_page(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Go home page")
}

func handleRequest()  {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts_page/", home_page)
  http.ListenAndServe(":8081", nil)
}

func main()  {
  handleRequest()
}
