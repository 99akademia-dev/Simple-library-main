package main

import "fmt"

func main() {
    fmt.Println("Проект 'Простая библиотека' запущен!")
    
    book := Book{
        ID:       1,
        Title:    "Война и мир",
        Author:   "Лев Толстой",
        IsIssued: false,
    }
    
    reader := Reader{
        ID:   1,
        Name: "Иван Иванов",
    }
    
    fmt.Printf("Книга: %s, Автор: %s\n", book.Title, book.Author)
    fmt.Printf("Читатель: %s\n", reader.Name)
    
    if book.IsIssued {
        fmt.Println("Книга выдана")
    } else {
        fmt.Println("Книга доступна")
    }
}