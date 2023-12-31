# SimpleLayerOne

## Overview

SimpleLayerOne is a simple encryption example implemented in Go, providing an HTTP server for encrypting data. This project includes a Go API and a JavaScript client for testing the API.

## Repository Structure

- `main.go`: The main Go file to run the HTTP server for the encryption service.
- `/pkg/encrypt/layerone.go`: Contains the encryption logic using ChaCha20.

## Getting Started

### Prerequisites

- Go (Version 1.15 or later)
- Node.js
- Git

### Setting Up the Go API

1. **Clone the Repository**:
https://github.com/ShaneSCalder/SimpleLayerOne/


2. **Initialize and Download Dependencies**:
- go mod init SimpleLayerOne
- go mod tidy

  
3. **Run the Server**:
- go run main.go

  This starts the server on `http://localhost:8080/encrypt`.

### Setting Up the JavaScript Client

4. **Create a Project Folder**:
- copy the test files into a new project
- Create a folder for the client-side testing.
- Place `testapi.js` and `testdata.json` inside this folder.

5. **Install Axios**:
Navigate to your project folder and run:

npm install axios

6. **Running the Test Script**:
In the project folder, run:

- node testapi.js
  
## Usage

The Go API exposes an endpoint (`/encrypt`) for encrypting data:

- **URL**: `http://localhost:8080/encrypt`
- **Method**: `POST`
- **Request Body**: JSON object with a `data` field.

The JavaScript client (`testapi.js`) sends data from `testdata.json` to this endpoint, receives the encrypted data and keys, and stores them locally.

## Disclaimer

This project is intended for educational and testing purposes only. It should not be used as a production encryption service. The authors are not responsible for any misuse or damage caused by this software.
