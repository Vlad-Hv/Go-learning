# CLI Expense Tracker

a small CLI programm to manage your expences

## features
- add expence
- show all expences
- show amount total

## how to run
```bash
go run ./cmd/app
```
## structure section
```text
CLI-Expense-Tracker/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── expense/
│   └── ...
├── README.md
├── .gitignore
└── go.mod
```

`cmd/app` - application entry point
`internal/expense` — expense model and business
`internal/ui` - ui system and getting info
`internal/storage` - storage model and business
`internal/validate` - validate all inputs
