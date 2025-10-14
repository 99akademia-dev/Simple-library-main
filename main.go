package main

import "fmt"

func main() {
	user1 := Reader{
		ID:        1,
		FirstName: "Агунда",
		LastName:  "Абаева",
		IsActive:  true,
	}

	user2 := Reader{
		ID:        2,
		FirstName: "Сергей",
		LastName:  "Меняйло",
		IsActive:  true,
	}
	book1 := Book{
		ID:       1,
		Title:    "Преступление и наказание",
		Author:   "Федор Достоевский",
		Year:     1866,
		IsIssued: false,
	}

	book2 := Book{
		ID:       2,
		Title:    "Война и мир",
		Author:   "Лев Толстой",
		Year:     1869,
		IsIssued: false,
	}

	library := NewLibrary()
	library.AddBook(&book1)
	library.AddBook(&book2)
	library.AddReader(&user1)
	library.AddReader(&user2)

	fmt.Println("=== Демонстрация работы библиотеки ===")
	fmt.Printf("Книга 1: %s\n", book1.String())
	fmt.Printf("Книга 2: %s\n", book2.String())
	fmt.Printf("Статус выдачи книг: %t, %t\n", book1.IsIssued, book2.IsIssued)
	fmt.Println("\n--- Выдача книги активному читателю ---")
	book1.IssueBook(&user1)
	fmt.Printf("Статус книги после выдачи: %s\n", book1.String())
	fmt.Println("\n--- Попытка выдать уже выданную книгу ---")
	book1.IssueBook(&user2)
	fmt.Println("\n--- Выдача второй книги другому читателю ---")
	book2.IssueBook(&user2)
	fmt.Printf("Статус книги после выдачи: %s\n", book2.String())
	fmt.Println("\n--- Возврат книги ---")
	book1.ReturnBook()
	fmt.Printf("Статус книги после возврата: %s\n", book1.String())
	fmt.Println("\n--- Выдача книги неактивному читателю ---")
	user1.Deactivate()
	book1.IssueBook(&user1)
	fmt.Println("\n--- Выдача книги после активации ---")
	user1.Activate()
	book1.IssueBook(&user1)
	fmt.Println("\n--- Информация о читателях ---")
	fmt.Println(user1.String())
	fmt.Println(user2.String())
	fmt.Println("\n--- Демонстрация метода AssignBook ---")
	user2.AssignBook(&book2)
	fmt.Println("\n--- Финальный статус всех книг ---")
	fmt.Printf("Книга 1: %s\n", book1.String())
	fmt.Printf("Книга 2: %s\n", book2.String())
}
