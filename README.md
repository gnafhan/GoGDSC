<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://fingers-site-production.s3.eu-central-1.amazonaws.com/uploads/images/72Risoyey04MAd2A4ZQUZE5nl0otT4KY39ah2izD.webp" alt="Project logo"></a>
</p>

<h3 align="center">Golang Todo List GDSC</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/gnafhan/GoGDSC.svg)](https://github.com/gnafhan/GoGDSC/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/gnafhan/GoGDSC.svg)](https://github.com/gnafhan/GoGDSC/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Api Documentation](#api)
- [Usage](#usage)
- [Built Using](#built_using)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)
- [License](#license)

## 🧐 About <a name = "about"></a>

Welcome to our project! 🎉 This is a straightforward yet powerful Todo List application, meticulously crafted using the robust and efficient Golang language. 🚀

Our project harnesses the power of Firebase, a versatile and scalable NoSQL cloud database, to store our data. This ensures that your todo items are safely stored and readily accessible, anytime and anywhere! 💾

But, what about security, you ask? Fear not! We’ve got it covered. Our application uses JWT (JSON Web Tokens) for authentication. This means that your data is not just stored; it’s stored securely. 🔒

And that’s not all! Our application is built on the Gin Framework, a web framework written in Golang. It’s known for its speed and efficiency, just like a well-mixed gin and tonic! 🍸

So, why wait? Dive in and start organizing your life with our Todo List application. Happy tasking! 😊

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

- Golang installed on your machine.
- Firebase account.
- `.env` file with your JWT secret key (`SECRET_KEY`).
- `secret.json` file with your Firebase configuration.

```
Give examples
```

### Installing Go Modules

To install Go modules for this project, navigate to the project directory and run:

```
go mod init
```

and

```
go mod tidy
```


## 🔧 API Documentation <a name = "api"></a>

###Temporary Array Storage Endpoints

- ```GET /```- Retrieve all todo items from temporary array storage.
- ```GET /:id```- Retrieve a todo item by ID from temporary array storage.
- ```POST /```- Add a todo item to temporary array storage.
- ```PUT /:id```- Update a todo item by ID in temporary array storage.
- ```DELETE /:id```- Delete a todo item by ID from temporary array storage.

###Firebase Endpoints

- ```POST /register```- Register a user in Firebase and get a JWT token.
- ```POST /login```- Login a user in Firebase and get a JWT token.
- ```GET /firebase``` - Retrieve all todo items for a user from Firebase (requires token).
- ```GET /firebase/:id```- Retrieve a todo item by ID for a user from Firebase (requires token).
- ```POST /firebase```  Add a todo item for a user to Firebase (requires token).
- ```PUT /firebase/:id``` - Update a todo item by ID for a user in Firebase (requires token).
- ```DELETE /firebase/:id```- Delete a todo item by ID for a user from Firebase (requires token).

###Authorization
- Request header with ```Authorization``` key

## 🎈 Usage <a name="usage"></a>

After installing the necessary prerequisites and setting up the environment variables, you can run the application using:
```go run init.go```

## ⛏️ Built Using <a name = "built_using"></a>

- [Firebase](https://firebase.google.com/) - Database
- [Golang](https://go.dev/) - Programming Language
- [Gin](https://gin-gonic.com/) - Server Framework


## ✍️ Authors <a name = "authors"></a>

- [@kgnafhan](https://github.com/gnafhan) - All work


## ⚖️ License <a name = "license"></a>
This project is licensed under the MIT License.
