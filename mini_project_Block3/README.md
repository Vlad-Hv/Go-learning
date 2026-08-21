# CLI Elevator Controller

a small CLI programm to control the elevator

## Features
- Choose a target floor
- Open door
- Close door
- Print status report

## How to run

terminal:

```bash
go run "./cmd/app"
```

## Project Structure

``` text
/cmd/app - application entry point

/internal/elevator - elevator state and business logic

/internal/ui - user input and output

/internal/validation - input validation and invariant checks
```

## Architecture

The application is built around a single elevator state.
The elevator package owns the elevator stat.
The ui package handles console input and output, while validation checks user input, business rules, invariants.
The main package coordinates the application flow.