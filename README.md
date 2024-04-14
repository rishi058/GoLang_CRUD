# Golang Book API

This is a simple Golang project that demonstrates a CRUD API for managing books. It uses the Fiber web framework and PostgreSQL as the database.

## Features

- Create a new book (POST /api/create_books)
- Get a book by ID (GET /api/get_books/:id)
- Get all books (GET /api/books)
- Delete a book (DELETE /api/delete_book/:id)

## Prerequisites

- Golang (version 1.16 or higher)
- PostgreSQL (version 12 or higher)

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/your-username/golang-book-api.git
```

2. Navigate to the project directory:

```bash
cd golang-book-api
```

3. Set up the database:

   - Create a new PostgreSQL database.
   - Update the database connection details in the `config/config.go` file.

4. Install the dependencies:

```bash
go mod download
```

5. Run the application:

```bash
go run main.go
```

The server will start running at `http://localhost:3000`.

## Endpoints

### Create a Book

**Request**:
```
POST /api/create_books
{
    "title": "Book Title",
    "author": "Book Author",
    "publisher": "Book Publisher"
}
```

**Response**:
```
{
    "id": 1,
    "title": "Book Title",
    "author": "Book Author",
    "publisher": "Book Publisher"
}
```

### Get a Book by ID

**Request**:
```
GET /api/get_books/1
```

**Response**:
```
{
    "id": 1,
    "title": "Book Title",
    "author": "Book Author",
    "publisher": "Book Publisher"
}
```

### Get All Books

**Request**:
```
GET /api/books
```

**Response**:
```
[
    {
        "id": 1,
        "title": "Book Title 1",
        "author": "Book Author 1",
        "publisher": "Book Publisher 1"
    },
    {
        "id": 2,
        "title": "Book Title 2",
        "author": "Book Author 2",
        "publisher": "Book Publisher 2"
    },
    {
        "id": 3,
        "title": "Book Title 3",
        "author": "Book Author 3",
        "publisher": "Book Publisher 3"
    }
]
```

### Delete a Book

**Request**:
```
DELETE /api/delete_book/1
```

**Response**:
```
{
    "message": "Book deleted successfully"
}
```

## Data Model

The `Book` struct represents the structure of a book in the application:

```go
type Book struct {
    ID        int    `json:"id"`
    Author    string `json:"author"`
    Title     string `json:"title"`
    Publisher string `json:"publisher"`
}
```

## Contributing

If you find any issues or have suggestions for improvements, feel free to create a new issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
