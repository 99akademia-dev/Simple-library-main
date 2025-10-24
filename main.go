package main

import "fmt"

func main() {
	fmt.Println("Запуск системы управления библиотекой...")

	myLibrary := &Library{}

	fmt.Println("\n--- Наполняем библиотеку ---")
	

	reader1, err := myLibrary.AddReader("Тамерлан", "Джигкаев")
	if err != nil {
		fmt.Printf("Ошибка при добавлении читателя: %v\n", err)
		return
	}
	fmt.Printf("Зарегистрирован новый читатель: %s\n", reader1)

	reader2, err := myLibrary.AddReader("Линда", "Элбакянц")
	if err != nil {
		fmt.Printf("Ошибка при добавлении читателя: %v\n", err)
		return
	}
	fmt.Printf("Зарегистрирован новый читатель: %s\n", reader2)

	book1, err := myLibrary.AddBook(1984, "1984", "Джордж Оруэлл")
	if err != nil {
		fmt.Printf("Ошибка при добавлении книги: %v\n", err)
		return
	}
	fmt.Printf("Добавлена новая книга: %s\n", book1)

	book2, err := myLibrary.AddBook(1967, "Мастер и Маргарита", "Михаил Булгаков")
	if err != nil {
		fmt.Printf("Ошибка при добавлении книги: %v\n", err)
		return
	}
	fmt.Printf("Добавлена новая книга: %s\n", book2)

	_, err = myLibrary.AddBook(1984, "1984", "Джордж Оруэлл")
	if err != nil {
		fmt.Printf("Ошибка при добавлении дубликата книги: %v\n", err)
	}

	fmt.Println("\n--- Библиотека готова к работе ---")
	fmt.Printf("Количество читателей: %d\n", myLibrary.GetReadersCount())
	fmt.Printf("Количество книг: %d\n", myLibrary.GetBooksCount())


	fmt.Println("\n=== СЦЕНАРИИ ИСПОЛЬЗОВАНИЯ ===")

	fmt.Println("\n1. Успешная выдача книги:")
	err = myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Printf("Ошибка выдачи книги: %v\n", err)
	} else {
		fmt.Println("Книга успешно выдана читателю")
		
		book, _ := myLibrary.FindBookById(1)
		if book != nil {
			fmt.Printf("Статус книги после выдачи: %s\n", book)
		}
	}

	fmt.Println("\n2. Попытка выдать уже выданную книгу:")
	err = myLibrary.IssueBookToReader(1, 2)
	if err != nil {
		fmt.Printf("Ошибка выдачи книги: %v\n", err)
	} else {
		fmt.Println("Книга успешно выдана")


	}
	fmt.Println("\n3. Попытка выдать книгу несуществующему читателю:")
	err = myLibrary.IssueBookToReader(2, 999)
	if err != nil {
		fmt.Printf("Ошибка выдачи книги: %v\n", err)
	} else {
		fmt.Println("Книга успешно выдана")
	}

	fmt.Println("\n4. Успешный возврат книги:")
	err = myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Printf("Ошибка возврата книги: %v\n", err)
	} else {
		fmt.Println("Книга успешно возвращена в библиотеку")
		
		book, _ := myLibrary.FindBookById(1)
		if book != nil {
			fmt.Printf("Статус книги после возврата: %s\n", book)
		}
	}

	fmt.Println("\n5. Попытка вернуть книгу, которая уже в библиотеке:")
	err = myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Printf("Ошибка возврата книги: %v\n", err)
	} else {
		fmt.Println("Книга успешно возвращена")
	}

	fmt.Println("\n6. Выдача и возврат другой книги:")
	err = myLibrary.IssueBookToReader(2, 2)
	if err != nil {
		fmt.Printf("Ошибка выдачи книги: %v\n", err)
	} else {
		fmt.Println("Книга успешно выдана")
		
		book, _ := myLibrary.FindBookById(2)
		if book != nil {
			fmt.Printf("Статус книги после выдачи: %s\n", book)
		}
	}

	err = myLibrary.ReturnBook(2)
	if err != nil {
		fmt.Printf("Ошибка возврата книги: %v\n", err)
	} else {
		fmt.Println("Книга успешно возвращена")
		
		book, _ := myLibrary.FindBookById(2)
		if book != nil {
			fmt.Printf("Статус книги после возврата: %s\n", book)
		}
	}

	fmt.Println("\n--- Работа системы завершена ---")
}