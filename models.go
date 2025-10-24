package main

import (
	"errors"
	"fmt"
)

type Book struct {
	ID       int
	Year     int
	Title    string
	Author   string
	IsIssued bool
	ReaderId *int
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

type Library struct {
	Books   []*Book
	Readers []*Reader

	lastBookID   int
	lastReaderID int
}

func (b *Book) IssueBook(readerID int) error {
	if b.IsIssued {
		return errors.New("книга уже используется")
	}
	b.IsIssued = true
	b.ReaderId = &readerID
	return nil
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}
	b.IsIssued = false
	b.ReaderId = nil
	return nil
}

func (b Book) String() string {
	status := "не используется."
	if b.IsIssued {
		status = fmt.Sprintf("используется читателем %d", *b.ReaderId)
	}
	return fmt.Sprintf("ID: %d, %s (%s %d), книга %s", b.ID, b.Title, b.Author, b.Year, status)
}

func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r Reader) String() string {
	status := "активен"
	if !r.IsActive {
		status = "неактивен"
	}
	return fmt.Sprintf("ID: %d, %s %s, статус: %s", r.ID, r.FirstName, r.LastName, status)
}

func (l *Library) AddReader(firstName, lastName string) (*Reader, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("имя и фамилия читателя не могут быть пустыми")
	}
	
	l.lastReaderID++
	newReader := &Reader{
		ID:        l.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	l.Readers = append(l.Readers, newReader)
	return newReader, nil
}

func (l *Library) AddBook(year int, title, author string) (*Book, error) {
	if year <= 0 {
		return nil, errors.New("год издания должен быть положительным числом")
	}
	if title == "" || author == "" {
		return nil, errors.New("название и автор книги не могут быть пустыми")
	}

	for _, book := range l.Books {
		if book.Title == title && book.Author == author {
			return nil, fmt.Errorf("книга '%s' автора '%s' уже существует в библиотеке", title, author)
		}
	}
	
	l.lastBookID++
	newBook := &Book{
		ID:       l.lastBookID,
		Year:     year,
		Title:    title,
		Author:   author,
		IsIssued: false,
	}
	l.Books = append(l.Books, newBook)
	return newBook, nil
}

func (l *Library) FindBookById(id int) (*Book, error) {
	for _, book := range l.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (l *Library) FindReaderById(id int) (*Reader, error) {
	for _, reader := range l.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

func (l *Library) IssueBookToReader(bookId, readerId int) error {
	book, err := l.FindBookById(bookId)
	if err != nil {
		return err
	}
	
	reader, err := l.FindReaderById(readerId)
	if err != nil {
		return err
	}
	
	if !reader.IsActive {
		return errors.New("нельзя выдать книгу неактивному читателю")
	}
	
	return book.IssueBook(readerId)
}

func (l *Library) ReturnBook(bookId int) error {
	book, err := l.FindBookById(bookId)
	if err != nil {
		return err
	}
	return book.ReturnBook()
}

func (l *Library) GetBooksCount() int {
	return len(l.Books)
}

func (l *Library) GetReadersCount() int {
	return len(l.Readers)
}