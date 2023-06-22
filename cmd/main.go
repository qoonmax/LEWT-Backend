package main

import (
	text "LEWT_Backend"
	"LEWT_Backend/keyboard"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var data *text.Data

func init() {
	data = text.NewText()
}

func startClient() {
	appPath := "/Applications/LEWT Client.app"
	err := exec.Command("open", appPath).Start()
	if err != nil {
		panic(fmt.Sprintf("Ошибка при запуске приложения: %v", err))
	}

	waitCh := make(chan os.Signal, 1)
	signal.Notify(waitCh, os.Interrupt, syscall.SIGTERM)

	<-waitCh

	err = exec.Command("pkill", "-f", "LEWT Client").Run()
	if err != nil {
		panic(fmt.Sprintf("Ошибка при завершении приложения: %v", err))
	}
}

func main() {
	go startClient()
	go keyboard.Listen(data)
	handleFunc()
}

func handleFunc() {
	http.HandleFunc("/", getText)
	http.ListenAndServe(":3333", nil)
}

func getText(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен!", http.StatusMethodNotAllowed)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(jsonData)
	return
}
