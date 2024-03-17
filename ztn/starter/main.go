package main

import (
	"fmt"
	"github.com/kr/pretty"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}

		fmt.Println(member.name, ":", title, "", audit.checkOut.String(), "through", returnTime)
	}
}

func printLibraryMembersAudit(library *Library) {
	fmt.Println("------------Members Audit------------")
	for _, member := range library.members {
		printMemberAudit(&member)
	}
	fmt.Println("------------Members Audit------------")
}

func printLibraryBooksAudit(library *Library) {
	fmt.Println("------------Books Audit------------")
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println("------------Books Audit------------")
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	if book.lended == book.lended {
		fmt.Println("No more books available to lend")
		return false
	}

	book.lended += 1

	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]

	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()

	member.books[title] = audit

	return true
}

func main() {

	var library Library

	library = Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}

	library.books["The Little Go Book"] = BookEntry{
		total:  3,
		lended: 0,
	}

	library.books["Let's learn go"] = BookEntry{
		total:  2,
		lended: 0,
	}

	library.books["Go Bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.members["Jason"] = Member{"Jason", make(map[Title]LendAudit)}
	library.members["Billy"] = Member{"Billy", make(map[Title]LendAudit)}
	library.members["Susana"] = Member{"Susana", make(map[Title]LendAudit)}

	fmt.Printf("%+v\n", pretty.Sprint(library))

	printLibraryBooksAudit(&library)
	printLibraryMembersAudit(&library)
}
