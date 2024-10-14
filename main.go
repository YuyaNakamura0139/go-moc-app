package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Film struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
	Description   string `json:"description"`
}

func filmsHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://ghibliapi.vercel.app/films")
	if err != nil {
		http.Error(w, "Failed to fetch films", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch films", resp.StatusCode)
		return
	}

	// レスポンスボディから読み取った生のバイト列
	body, err := io.ReadAll(resp.Body)
	var films []Film
	// json形式のバイト列を構造体にデコード
	if err := json.Unmarshal(body, &films); err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(films)
}

func main() {
	fmt.Println("Starting the server!")

	// ルートとハンドラ関数を定義
	http.HandleFunc("/api/films", filmsHandler)

	// 8081番ポートでサーバーを起動
	http.ListenAndServe(":8081", nil)
}
