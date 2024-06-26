package elements

import "time"

//Book: Represents a book in the library.
//Attributes: ID, Title, Author, YearPublished, IsAvailable.

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
	IsAvailable   bool
	// Quantity      int
}

//Member: Represents a library member.
//Attributes: ID, Name, Email.

type Member struct {
	ID    int
	Name  string
	Email string
}

//Loan: Represents a loan record for books borrowed by members.
//Attributes: ID, MemberID, BookID, LoanDate, ReturnDate.

type Loan struct {
	ID         int
	MemberID   int
	BookID     int
	LoanDate   time.Time
	ReturnDate int
}
