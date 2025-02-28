.PHONY: backend frontend

backend:
	cd ./backend && go run main.go

frontend:
	cd ./frontend && npm run dev