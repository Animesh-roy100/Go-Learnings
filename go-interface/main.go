package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type AccountNotifier interface {
	NotifyAccountCreated(context.Context, Account) error
}

type SimpleAccountNotifier struct{}

func (n SimpleAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("new account created", "username", account.Username, "email", account.Email)
	return nil
}

type BetterAccountNotifier struct{}

func (n BetterAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("new account created by better notifier", "username", account.Username, "email", account.Email)
	return nil
}

type Account struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

func (h *AccountHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error("failed to decode the response body", "err", err)
		return
	}
	if err := h.AccountNotifier.NotifyAccountCreated(r.Context(), account); err != nil {
		slog.Error("failed to notify account created", "err", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

// func notifyAccountCreated(account Account) error {
// 	time.Sleep(time.Millisecond * 500)
// 	slog.Info("new account created", "username", account.Username, "email", account.Email)
// 	return nil
// }

func main() {
	mux := http.NewServeMux()

	accountHandler := &AccountHandler{
		AccountNotifier: BetterAccountNotifier{},
	}
	mux.HandleFunc("POST /account", accountHandler.handleCreateAccount)

	http.ListenAndServe(":9090", mux)
}
