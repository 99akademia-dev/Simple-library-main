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
	fmt.Printf("читатель %s %s активирован\n", r.FirstName, r.LastName)
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

func (r *Reader) AssignBook(book *Book) {
	if book == nil {
		fmt.Println("Ошибка: книга не существует")
		return
	}
	fmt.Printf("Читатель %s %s взял книгу '%s'\n", r.FirstName, r.LastName, book.String())
}

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус: %s", b.Title, b.Author, b.Year, status)
}

func (b *Book) IssueBook(reader *Reader) {
	if !reader.IsActive {
		fmt.Printf("Читатель %s %s не активен и не может получить книгу.\n", reader.FirstName, reader.LastName)
		return
	}

	if b.IsIssued {
		fmt.Printf("Книга %s уже выдана\n", b.Title)
		return
	}

	b.IsIssued = true
	b.ReaderID = &reader.ID
	fmt.Printf("Книга %s выдана читателю %s %s\n", b.Title, reader.FirstName, reader.LastName)

	reader.AssignBook(b)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так в библиотеке\n", b.Title)
		return
	}

	b.ReaderID = nil
	b.IsIssued = false
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}

type Library struct {
	Books   map[int]*Book
	Readers map[int]*Reader
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]*Book),
		Readers: make(map[int]*Reader),
	}
}

func (l *Library) AddBook(book *Book) {
	l.Books[book.ID] = book
}

func (l *Library) AddReader(reader *Reader) {
	l.Readers[reader.ID] = reader
}
