module monorepo-template/api

go 1.20

require (
	github.com/go-chi/chi/v5 v5.0.8
	monorepo-template/pkg v0.0.0-00010101000000-000000000000
)

replace monorepo-template/pkg => ../pkg
