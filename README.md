# Go Personal Library

This project is a RESTful API built using Go to manage personal book libraries. The API allows users to store information about books they are reading, including title, author, number of pages, publisher, and comments.

## Features

- Add, update, delete, and retrieve books.
- Store additional comments on books.
- Clean Architecture using repository and use-case layers.
- MongoDB for data persistence.
- Swagger documentation for easy API testing.

## Prerequisites

Before running the project, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or later)
- [MongoDB](https://www.mongodb.com/try/download/community) (running locally or on a cloud provider)
- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/) (optional for running MongoDB in a container)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/rfulgencio3/go-personal-library.git
cd go-personal-library


### 2. Set Up Environment Variables

Create a .env file in the root directory of the project with the following content:

SERVER_PORT=8080
MONGO_URI=mongodb://localhost:27017
MONGO_DATABASE=personal_library
MONGO_COLLECTION=books

### 3. Install Dependencies

Ensure that all dependencies are installed by running:
go mod tidy


### 4. Run MongoDB (Optional via Docker)
If you don't have MongoDB running locally, you can start a MongoDB instance using Docker:

docker run -d -p 27017:27017 --name mongo-library mongo

### 5. Running the API
Once everything is set up, you can start the API server:

go run cmd/main.go

The API will be running on http://localhost:8080. You can modify the port by updating the SERVER_PORT variable in the .env file.

### 6. Access the Swagger Documentation
Once the server is running, you can access the Swagger API documentation by visiting:

http://localhost:8080/swagger/index.html

This will allow you to test the API endpoints directly from your browser.

API Endpoints
Method	Endpoint	Description
POST	/books	Create a new book
GET	/books/{id}	Get a book by ID
PUT	/books/{id}	Update a book by ID
DELETE	/books/{id}	Delete a book by ID
GET	/books	Get all books
Contributing
Feel free to submit pull requests or issues to improve this project.

License
This project is licensed under the MIT License - see the LICENSE file for details.