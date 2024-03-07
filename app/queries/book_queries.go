package queries

import (
	"database/sql"

	"github.com/quocquann/book-api/app/models"
)

type BookQueries struct {
	*sql.DB
}

func (q *BookQueries) GetBooks() ([]models.Book, error) {
	books := []models.Book{}
	queryString := `SELECT isbn, title, summary, first_name as author, name as genre 
					FROM 
					Book JOIN Author 
					ON Book.author_id = Author.author_id 
					JOIN Genre 
					ON Book.genre_id = Genre.genre_id`
	rows, err := q.Query(queryString)
	if err != nil {
		return books, nil
	}

	defer rows.Close()

	for rows.Next() {
		book := models.Book{}
		if err := rows.Scan(&book.Isbn, &book.Title, &book.Summary, &book.Author, &book.Genre); err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (q *BookQueries) GetBookByIsbn(isbn string) (models.Book, error) {
	book := models.Book{}
	queryString := `SELECT isbn, title, summary, first_name as author, name as genre 
					FROM 
					Book JOIN Author 
					ON Book.author_id = Author.author_id 
					JOIN Genre 
					ON Book.genre_id = Genre.genre_id 
					WHERE isbn = ?`

	row := q.QueryRow(queryString, isbn)

	if err := row.Scan(&book.Isbn, &book.Title, &book.Summary, &book.Author, &book.Genre); err != nil {
		return book, err
	}
	return book, nil
}

func (q *BookQueries) GetBookByAuthor(authorId string) ([]models.Book, error) {
	books := []models.Book{}
	queryString := "SELECT COUNT(*) as count FROM Author WHERE author_id = ?"
	row := q.QueryRow(queryString, authorId)
	_ = row
	return books, nil
	//TODO: Not done yet
}
