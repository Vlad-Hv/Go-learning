# Go Backend Developer Roadmap

## Purpose

This roadmap is the main learning guide for becoming a junior-ready Go Backend Developer.

The roadmap is treated as a stable route. Topics should not be skipped. Large blocks may be split into smaller blocks if needed, but the core content must stay.

The main goal is not to finish quickly, but to build a strong foundation in Go, backend development, professional English, and practical project work.

---

## Learning Goal

Become a junior-ready Go Backend Developer by October.

By the end of this roadmap, the student should be able to:

* write Go programs without copying code blindly;
* understand Go syntax and Go-specific thinking;
* work with functions, errors, arrays, slices, maps, structs, pointers, methods, interfaces, goroutines, channels, context;
* build CLI applications;
* build REST APIs;
* work with PostgreSQL;
* use migrations;
* implement authentication and authorization;
* write tests;
* use Docker and Docker Compose;
* structure backend projects properly;
* explain technical decisions in English;
* prepare a GitHub portfolio project.

---

## Learning Rules

Each block follows this structure:

1. Theory first.
2. Tasks only when requested.
3. Three tasks with increasing difficulty.
4. One mini-project.
5. Code review.
6. Professional English vocabulary:

   * nouns;
   * verbs;
   * useful developer phrases;
   * short realistic developer report.
7. Vocabulary check.
8. Final evaluation.
9. Move to the next block only after everything is completed.

No block is skipped.

If a topic is too large, it can be split into smaller blocks.

---

## Weekly Rules

At the end of each week:

1. Complete a weekly project.
2. Do a Go review.
3. Do an English technical interview.
4. Make a mistake review.
5. Repeat weak words and weak Go concepts.
6. Summarize mastered Go topics.
7. Summarize mastered English words.
8. Update progress.

Starting from Week 2, weekly reviews must include:

* Go topics really learned, not just touched;
* English words really learned, not just listed;
* common mistakes;
* fixed mistakes;
* next week focus.

---

# Week 1 — Go Basics

Status: Completed.

## Block 1 — Go Program Structure

Topics:

* `package`
* `import`
* `func main`
* `go run`
* `go build`
* `go fmt`
* `go version`
* comments
* naming conventions

English vocabulary:

* package
* entry point
* import
* standard library
* compile
* run
* formatting
* naming convention

Practice:

* rewrite first Go program cleanly;
* create `profile.go`;
* run program with `go run`;
* build binary with `go build`.

Main goal:

Understand the basic structure of every Go program.

---

## Block 2 — Variables, Constants, Types

Topics:

* `var`
* `:=`
* `const`
* `string`
* `int`
* `float64`
* `bool`
* zero value
* type inference
* basic type conversion

English vocabulary:

* variable
* constant
* type
* explicit type
* inferred type
* zero value
* type conversion

Practice:

* age calculator;
* temperature converter;
* simple currency converter;
* BMI calculator.

Main goal:

Understand how Go stores data and how variable declarations work.

---

## Block 3 — Input and Control Flow

Topics:

* `fmt.Scanln`
* `err`
* `if`
* `else if`
* `else`
* comparison operators
* logical operators
* early return

English vocabulary:

* condition
* input
* invalid input
* return
* validate
* compare

Practice:

* password check;
* login check;
* age access check.

Mini-project:

Console Login / Registration System.

Main goal:

Understand user input, validation, conditions, and early return.

---

## Block 4 — Switch

Topics:

* `switch`
* `case`
* `default`
* grouped cases
* switch without expression
* `fallthrough`
* when to use `switch` vs `if`

English vocabulary:

* switch
* case
* default
* option
* command
* branch
* expression
* choose
* select
* match
* execute
* handle
* exit

Practice:

* CLI menu;
* weekday/weekend detector;
* age category detector.

Mini-project:

CLI Developer Assistant.

Main goal:

Understand how to choose between multiple program branches cleanly.

---

## Block 5 — For Loops

Topics:

* classic `for`
* Go-style while loop
* infinite loop
* `break`
* `continue`
* counters
* iterations

English vocabulary:

* loop
* iteration
* counter
* condition
* infinite loop
* statement
* iterate
* increment
* decrement
* skip
* terminate
* repeat

Practice:

* counter;
* multiplication table;
* break and continue task.

Mini-project:

Training Session Tracker.

Main goal:

Understand repetition and loop control.

---

## Block 6 — Arrays and Range

Topics:

* array
* fixed length
* index
* value
* `len`
* `range`
* `_`
* index/value pair

English vocabulary:

* array
* element
* index
* value
* length
* collection
* position
* total
* average
* maximum
* access
* iterate
* store
* retrieve
* calculate
* compare
* process

Practice:

* print array values;
* print index and value;
* calculate array sum.

Mini-project:

Student Grade Analyzer.

Main goal:

Understand fixed-size collections and iteration with `range`.

---

## Block 7 — Type Conversion and Week 1 Review

Topics:

* `int` to `float64`
* `float64` to `int`
* integer division
* floating-point division
* truncation
* preserving original variables
* basic calculations

English vocabulary:

* type
* type conversion
* integer
* floating-point number
* percentage
* discount
* calculation
* result
* expression
* convert
* cast
* truncate
* preserve
* assign
* evaluate

Practice:

* integer vs float division;
* discount calculator;
* student statistics.

Mini-project:

Shopping Receipt.

Main goal:

Understand why Go requires explicit type conversion and how to calculate averages correctly.

---

## Week 1 Final Project — CLI Student Management System

Status: Completed.

Features:

* store names of 5 students;
* store ages of 5 students;
* store final grades of 5 students;
* validate input;
* show all students;
* calculate average grade;
* find best student;
* show adult students;
* menu with `switch`;
* repeat menu with loop;
* use arrays;
* use `range`;
* use type conversion;
* use early return.

Main result:

The student built the first real CLI project using only Week 1 knowledge.

---

# Week 2 — Functions, Errors, Collections

Status: Current.

Week 2 is more important and deeper than Week 1. Functions and errors should not be rushed.

## Block 1 — Functions Basics

Topics:

* function declaration;
* calling functions;
* parameters;
* arguments;
* return values;
* multiple parameters;
* scope;
* local variables;
* one function = one responsibility.

English vocabulary:

* function
* parameter
* argument
* return value
* scope
* declaration
* call
* reusable code
* helper function

Practice:

* `calculateDiscount(price, percent)`
* `isAdult(age)`
* `add(a, b)`
* `formatUser(name, age)`

Mini-project:

Calculator Functions.

Main goal:

Stop writing everything inside `main()` and learn how to split code by meaning.

---

## Block 2 — Multiple Returns and Errors

Topics:

* multiple return values;
* `value, err`;
* `error`;
* `nil`;
* `errors.New`;
* `fmt.Errorf`;
* early return with errors;
* error messages;
* validation functions.

English vocabulary:

* error handling
* return an error
* wrap an error
* nil
* early return
* invalid input
* validation
* failure
* success

Practice:

* `divide(a, b)` with division by zero error;
* `validateEmail(email)`;
* `validateAge(age)`;
* `registerUser(name, email, age)`.

Mini-project:

Registration Validator.

Main goal:

Understand Go-style error handling.

---

## Block 3 — Functions with Arrays and Slices

Topics:

* passing arrays to functions;
* why arrays are limited;
* slice basics;
* `append`;
* `len`;
* `cap`;
* `range` over slices;
* difference between array and slice;
* why slices are used more often in real Go.

English vocabulary:

* array
* slice
* length
* capacity
* append
* collection
* dynamic size
* pass data
* process collection

Practice:

* create list of programming languages;
* calculate sum through function;
* find maximum through function;
* delete element from slice;
* find duplicates.

Mini-project:

To-Do List CLI.

Main goal:

Understand why slices are the main collection type in Go.

---

## Block 4 — Maps

Topics:

* `map`
* key-value pairs
* lookup
* `delete`
* comma ok idiom
* map iteration
* map zero value
* checking if key exists

English vocabulary:

* map
* dictionary
* key
* value
* key-value pair
* lookup
* delete
* exists
* missing key

Practice:

* username to age map;
* word counter;
* vocabulary dictionary;
* check if user exists;
* delete user.

Mini-project:

Vocabulary Trainer v1.

Main goal:

Understand dictionary-like data storage in Go.

---

## Week 2 Final Project — CLI Vocabulary Trainer

Features:

* add English term;
* add Russian translation;
* show all terms;
* quiz mode;
* delete term;
* validate input;
* use functions;
* use errors;
* use slices and/or maps;
* use professional English vocabulary.

Main goal:

Build a useful CLI application connected to the student's English learning.

---

# Week 3 — Structs, Pointers, Methods

## Block 1 — Structs

Topics:

* `struct`
* fields
* struct literal
* nested struct
* anonymous struct
* composition basics
* entity/model thinking

English vocabulary:

* struct
* field
* literal
* nested
* composition
* entity
* model

Practice:

* `User`
* `Product`
* `Order`
* `Book`
* `Course`

Main goal:

Learn how to model real-world entities in Go.

---

## Block 2 — Pointers

Topics:

* pointer
* address
* dereference
* `&`
* `*`
* value semantics
* pass by value
* mutation
* reference-like behavior

English vocabulary:

* pointer
* address
* dereference
* pass by value
* mutate
* memory
* reference-like behavior

Practice:

* `changeName(name string)`
* `changeUser(user *User)`
* `increaseBalance(account *Account)`
* `applyDiscount(product *Product)`

Main goal:

Understand how Go passes data and how to modify existing values.

---

## Block 3 — Methods

Topics:

* method
* receiver
* value receiver
* pointer receiver
* behavior
* state
* encapsulation basics

English vocabulary:

* method
* receiver
* value receiver
* pointer receiver
* behavior
* state
* encapsulation

Practice:

* `User.ChangeEmail()`
* `Product.ApplyDiscount()`
* `Account.Deposit()`
* `Account.Withdraw()`

Mini-project:

Bank Account CLI.

Main goal:

Connect data and behavior using methods.

---

# Week 4 — Packages, Modules, Project Structure

## Block 1 — Modules

Topics:

* `go mod init`
* `go.mod`
* module path
* dependency
* `go get`
* `go mod tidy`

Practice:

* create module;
* initialize `go.mod`;
* run project as module.

---

## Block 2 — Packages

Topics:

* package naming;
* exported names;
* unexported names;
* visibility;
* internal package;
* cmd folder.

Practice:

* split CLI app into packages:

  * `cmd/app`
  * `internal/input`
  * `internal/vocabulary`

---

## Block 3 — Project Layout

Topics:

* `cmd`
* `internal`
* `pkg`
* README
* `.gitignore`
* project structure without overengineering

Weekly result:

Project should run with:

```bash
go run ./cmd/app
```

Main goal:

Start writing projects like real Go applications.

---

# Week 5 — Testing and Idiomatic Go

## Block 1 — Unit Tests

Topics:

* `testing`
* `_test.go`
* `TestXxx`
* `go test`

Practice:

* `TestValidateEmail`
* `TestWithdraw`
* `TestApplyDiscount`
* `TestWordCounter`

---

## Block 2 — Table-Driven Tests

Topics:

* test table;
* `t.Run`;
* edge cases;
* expected vs actual result.

Practice:

* rewrite tests in table-driven style.

---

## Block 3 — Idiomatic Go

Topics:

* `gofmt`
* clear names
* short variable names where appropriate
* small functions
* error handling style
* comments
* readability
* maintainability

Weekly project:

Refactor one previous project.

Main goal:

Write Go code that looks like Go, not translated Python/C++.

---

# Week 6 — Interfaces and Composition

## Block 1 — Interfaces

Topics:

* interface;
* implicit implementation;
* behavior;
* contract;
* abstraction;
* small interfaces.

Practice:

```go
type Storage interface {
    Save(word string, translation string) error
    Get(word string) (string, error)
}
```

---

## Block 2 — Composition over Inheritance

Topics:

* embedding;
* composition;
* struct embedding;
* interface embedding;
* reuse.

Practice:

* `User + Profile`
* `Logger interface`
* `Storage interface`

---

## Block 3 — Dependency Inversion Basics

Topics:

* handler uses service;
* service uses repository interface;
* constructor injection;
* layers.

Weekly project:

Rewrite CLI Vocabulary Trainer using `Storage` interface.

Main goal:

Prepare for backend architecture.

---

# Week 7 — Concurrency and Context

## Block 1 — Goroutines

Topics:

* goroutine;
* `go` keyword;
* concurrent execution;
* `sync.WaitGroup`;
* race condition.

Practice:

* run 5 tasks concurrently;
* process fake jobs;
* use WaitGroup.

---

## Block 2 — Channels

Topics:

* channel;
* send;
* receive;
* buffered channel;
* unbuffered channel;
* close;
* range over channel;
* deadlock.

Practice:

* worker pool;
* jobs channel;
* results channel.

---

## Block 3 — Context

Topics:

* `context.Context`;
* `WithCancel`;
* `WithTimeout`;
* deadline;
* cancellation;
* request lifecycle.

Mini-project:

Concurrent URL Checker.

Main goal:

Understand one of Go's strongest features: concurrency.

---

# Week 8 — HTTP and REST API

## Block 1 — HTTP Basics

Topics:

* request;
* response;
* method;
* status code;
* header;
* body;
* JSON;
* endpoint.

Practice:

* `GET /health`
* `GET /hello`
* `POST /echo`

---

## Block 2 — `net/http`

Topics:

* `http.HandleFunc`;
* `http.Server`;
* handler;
* `ResponseWriter`;
* `Request`;
* JSON encode/decode.

Practice:

In-memory Notes API:

* `GET /notes`
* `POST /notes`
* `GET /notes/{id}`
* `DELETE /notes/{id}`

---

## Block 3 — Router and Middleware

Topics:

* chi or gin;
* middleware;
* logging middleware;
* recovery middleware;
* request ID;
* route parameters.

Weekly project:

Notes REST API with in-memory storage.

Main goal:

Build the first real backend API.

---

# Week 9 — Clean Backend Structure

## Block 1 — Layers

Topics:

* handler layer;
* service layer;
* repository layer;
* model/entity;
* DTO;
* business logic.

Practice:

Split Notes API into:

* handler;
* service;
* repository;
* model.

---

## Block 2 — Config and Logging

Topics:

* environment variables;
* config struct;
* `slog`;
* log levels.

Practice:

* `PORT` from env;
* `LOG_LEVEL` from env;
* structured logs.

---

## Block 3 — Validation and Error Responses

Topics:

* request validation;
* error response;
* status codes;
* bad request;
* not found;
* internal server error.

Weekly result:

Notes API should look like a small real backend project.

---

# Week 10 — SQL and PostgreSQL

## Block 1 — SQL Basics

Topics:

* `SELECT`
* `INSERT`
* `UPDATE`
* `DELETE`
* `WHERE`
* `ORDER BY`
* `LIMIT`
* `JOIN`

Practice:

* users table;
* notes table;
* basic CRUD queries.

---

## Block 2 — PostgreSQL with Go

Topics:

* `database/sql` or `pgx`;
* connection pool;
* `Query`;
* `QueryRow`;
* `Exec`;
* `Scan`.

Practice:

Connect Notes API to PostgreSQL.

---

## Block 3 — Migrations

Topics:

* migration;
* schema;
* up migration;
* down migration;
* rollback;
* constraint;
* index;
* goose or golang-migrate.

Weekly result:

Notes API with PostgreSQL and migrations.

---

# Week 11 — Authentication and Security Basics

## Block 1 — Password Hashing

Topics:

* bcrypt;
* hash;
* salt;
* compare password;
* credentials.

Practice:

* `POST /register`
* `POST /login`

---

## Block 2 — JWT

Topics:

* JWT;
* claims;
* access token;
* expiration;
* Bearer token.

Practice:

* JWT login;
* auth middleware;
* protected routes.

---

## Block 3 — Ownership

Topics:

* user owns resource;
* authorization check;
* forbidden;
* unauthorized;
* permission.

Weekly result:

User can only access their own notes/tasks.

---

# Week 12 — Docker and Deployment Basics

## Block 1 — Docker

Topics:

* Dockerfile;
* image;
* container;
* build;
* run;
* ports.

---

## Block 2 — Docker Compose

Topics:

* `docker-compose.yml`;
* app service;
* postgres service;
* environment variables;
* volumes;
* networks.

Practice:

Run app + postgres using Docker Compose.

---

## Block 3 — Graceful Shutdown and Healthcheck

Topics:

* `http.Server`;
* `Shutdown`;
* context timeout;
* `/health`;
* signal.

Weekly result:

Project runs with:

```bash
docker compose up --build
```

---

# Week 13 — Final Project

## Task Manager API

Features:

* register;
* login;
* JWT;
* create project;
* list projects;
* create task;
* update task status;
* delete task;
* PostgreSQL;
* migrations;
* Docker Compose;
* tests;
* README.

Project structure:

```text
task-manager-api/
  cmd/api/main.go
  internal/
    config/
    handler/
    service/
    repository/
    model/
    auth/
  migrations/
  docker-compose.yml
  Dockerfile
  README.md
  go.mod
```

Main goal:

Build a portfolio-ready backend project.

---

# Week 14 — Interview and Portfolio

Topics:

* Go syntax;
* slices/maps;
* pointers;
* structs/methods;
* interfaces;
* errors;
* goroutines/channels;
* context;
* HTTP;
* REST;
* SQL;
* PostgreSQL;
* Docker;
* JWT;
* testing.

English interview questions:

* What is a goroutine?
* What is the difference between an array and a slice?
* How does Go handle errors?
* What is an interface in Go?
* What is middleware?
* What is the difference between authentication and authorization?
* How do you structure a Go backend project?

Portfolio polish:

* README;
* setup instructions;
* API examples;
* `.env.example`;
* screenshots or curl examples;
* clean commits.

---

# Rule for Future Changes

This roadmap is stable.

Allowed:

* split a big block into smaller blocks;
* add extra practice;
* add deep-dive lessons;
* add review tasks;
* repeat weak topics;
* slow down when needed.

Not allowed:

* skip core topics;
* remove backend topics;
* move forward without tasks and review;
* ignore English vocabulary;
* finish a week without a project.

---

# Current Progress

Completed:

* Week 1 — Go Basics
* Week 1 Final Project — CLI Student Management System

Current:

* Week 2 — Block 1: Functions

Next:

* Functions tasks
* Functions mini-project
* Functions vocabulary
* Week 2 Errors
