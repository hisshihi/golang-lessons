package main

import (
	"fmt"
	"time"
)

type (
	Title string // название книги
	Name  string // имя автора
)

// LendAudit - запись о выдаче книги
type LendAudit struct {
	checkOut time.Time // дата выдачи
	checkIn  time.Time // дата возврата
	Title    Title     // название книги
	ToMember Name      // имя читателя
}

// Member - читатель библиотеки
type Member struct {
	Name  Name
	Books map[Title]LendAudit // книги на руках
}

type BookEntry struct {
	Total  int // всего книг
	Lended int // выдано книг
}

type Library struct {
	Members map[Name]Member     // читатели библиотеки
	Books   map[Title]BookEntry // книги в библиотеке
}

// printMemberAudit - печатает информацию о выданных книгах для читателя
func (member *Member) printMemberAudit() {
	for _, audit := range member.Books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "not returned"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Printf("Book: %s, Checked out: %s, Returned: %s\n", audit.Title, audit.checkOut.String(), returnTime)
	}
}

func (library *Library) printInfoAboutMember() {
	for _, member := range library.Members {
		member.printMemberAudit()
	}
}

func (library *Library) printLibraryBooks() {
	for title, entry := range library.Books {
		fmt.Printf("Title: %s, Total: %d, Lended: %d\n", title, entry.Total, entry.Lended)
	}
}

// выдача книги
func (library *Library) lendBook(member *Member, title Title) bool {
	// проверяем, есть ли книга в библиотеке
	book, bookExists := library.Books[title]
	if !bookExists {
		fmt.Printf("Book %s is not available for lending.\n", title)
		return false
	}

	if book.Lended == book.Total {
		fmt.Printf("All copies of %s are currently lent out.\n", title)
		return false
	}

	book.Lended++
	library.Books[title] = book

	audit := LendAudit{
		checkOut: time.Now(),
		Title:    title,
		ToMember: member.Name,
	}

	member.Books[title] = audit
	return true
}

func (library *Library) returnBook(member *Member, title Title) bool {
	book, bookFound := library.Books[title]
	if !bookFound {
		fmt.Printf("Book %s does not belong to this library.\n", title)
		return false
	}

	// проверяем, что книга была выдана этому читателю
	audit, bookLent := member.Books[title]
	if !bookLent {
		fmt.Printf("Member %s did not borrow book %s.\n", member.Name, title)
		return false
	}

	book.Lended--
	library.Books[title] = book

	audit.checkIn = time.Now()
	member.Books[title] = audit
	return true
}

func main() {
	library := Library{
		Members: make(map[Name]Member),
		Books:   make(map[Title]BookEntry),
	}

	library.Books["Webapps with Go"] = BookEntry{Total: 3, Lended: 0}
	library.Books["Go Concurrency"] = BookEntry{Total: 1, Lended: 0}

	library.Members["Denis"] = Member{Name: "Denis", Books: map[Title]LendAudit{}}
	library.Members["Arina"] = Member{Name: "Arina", Books: map[Title]LendAudit{}}
	library.Members["Sergey"] = Member{Name: "Sergey", Books: map[Title]LendAudit{}}

	fmt.Println("\nInitial:")
	library.printLibraryBooks()
	library.printInfoAboutMember()

	member := library.Members["Denis"]
	checkedOut := library.lendBook(&member, "Go Concurrency")
	fmt.Println("\nCheck:")
	if checkedOut {
		library.printLibraryBooks()
		library.printInfoAboutMember()
	}

	returned := library.returnBook(&member, "Go Concurrency")
	fmt.Println("\nReturn:")
	if returned {
		library.printLibraryBooks()
		library.printInfoAboutMember()
	}
}
