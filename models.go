package main

import "fmt"

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (r Reader) DisplayReader() {
	fmt.Printf("читатель: %s %s (ID: %d)\n", r.FirstName, r.LastName, r.ID)
}

func (r *Reader) Deactivate() {
	r.IsActive = false
	fmt.Printf("читатель %s %s деактивирован\n", r.FirstName, r.LastName)
}

func (r *Reader) Activate() {
	r.IsActive = true
	fmt.Printf("читатель %s %s активаирован\n", r.FirstName, r.LastName)
}

func (r Reader) String() string {
	status := ""
	if r.IsActive {
		status = "Активен"
	} else {
		status = "Неактивен"
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, статус: %s", r.FirstName, r.LastName, r.ID, status)
}

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
}

func (b Book) String() string {
	return fmt.Sprintf("\"%s\" (%s, %d)", b.Title, b.Author, b.Year)
}

func (b *Book) IssueBook(reader *Reader) {
	if b.IsIssued {
		fmt.Printf("Книга %s уже выдана\n", b.Title)
		return
	}
	if !reader.IsActive {
		fmt.Printf("читатель %s %s не активен и не может получить книгу\n", reader.FirstName, reader.LastName)
		return
	}
	b.IsIssued = true
	fmt.Printf("книга %s выдана читателю %s %s\n", b.Title, reader.FirstName, reader.LastName)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("книга %s и так в библиотеке\n", b.Title)
		return
	}
	b.IsIssued = false
	fmt.Printf("книга %s возвращена в библиотеку\n", b.Title)
}
