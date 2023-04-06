# Go REST Api

> This is a simple REST API written in Go that follows the Model View Controller (MVC) pattern. The
> API utilizes only the Go standard library.

## Getting Started

To get started, follow these steps:

**Clone the repository:**

```bash
git clone https://github.com/nico-mayer/go-api.git
```

**Change to the project directory:**

```bash
cd go-api
```

**Run the application:**

```bash
go run .
```

> The application will start listening on port `8080`. You can test the API by making requests to the
> following endpoints:

-   `GET /people` : Returns all people.
-   `POST /people/create` : Creates a new Person
    -   Body format:
    ```json
    {
    	"id": "1",
    	"firstName": "John",
    	"lastName": "Doe",
    	"age": 30
    }
    ```

## Usage

You can use this API as a starting point for your own Go projects.
