# Interview Project

This project implements a single endpoint API service to retrieve a number's index from a file based on the given endpoint.

## Setup

### Environment Variables
Create a .env file in the root directory of the project with the following variables:
```
INPUT_PATH=/path/to/numbers.txt
LOG_LEVEL=info
PORT=8080
```

* INPUT_PATH: Specifies the path to the text file containing numbers.
* LOG_LEVEL: Sets the logging level for the application (debug, info, error).
* PORT: Specifies the port on which the server will run.

### Installation

1. Clone the repository
```
git clone git@github.com:Bartosz-D3V/recruitment-task-go.git
cd recruitment-task-go
```

2. Install dependencies
```
go mod tidy
```

### Running the Application
```
go run main.go
```

By default, the application starts on localhost at the port specified in the .env file (PORT variable).

## API Endpoint
### Get Number Endpoint
* Endpoint: /endpoint/{number}
* Method: GET
* Description: Retrieves a number from the file specified in INPUT_PATH based on the {number} provided in the endpoint.

### Example
#### Request
```
curl --location --request GET 'localhost:3000/endpoint/1100'
```
#### Response
```json
{
  "foundIndex": 11
}
```

## Testing
```
make test
```

