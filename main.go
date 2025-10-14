package main

import "fmt"

func main() {
	user1 := Reader{
		ID:        1,
		FirstName: "Агунда",
		LastName:  "Абаева",
		IsActive:  true,
	}

	book1 := Book{
		ID:       1,
		Title:    "Преступление и наказание",
		Author:   "Федор Достоевский",
		Year:     1866,
		IsIssued: false,
	}

	fmt.Println("=== Демонстрация работы библиотеки ===")
	fmt.Printf("Книга: %s\n", book1.String())
	fmt.Printf("Статус выдачи: %t\n", book1.IsIssued)
	fmt.Println("\n--- Выдача книги активному читателю ---")
	book1.IssueBook(&user1)
	fmt.Printf("Статус выдачи после попытки выдачи: %t\n", book1.IsIssued)
	fmt.Println("\n--- Повторная выдача книги ---")
	book1.IssueBook(&user1)
	fmt.Println("\n--- Возврат книги ---")
	book1.ReturnBook()
	fmt.Printf("Статус выдачи после возврата: %t\n", book1.IsIssued)
	fmt.Println("\n--- Выдача книги неактивному читателю ---")
	user1.Deactivate()
	book1.IssueBook(&user1)
	fmt.Println("\n--- Выдача книги после активации ---")
	user1.Activate()
	book1.IssueBook(&user1)
	fmt.Println("\n--- Информация о читателе ---")
	fmt.Println(user1.String())
}
