package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

const (
	githubUser     = "" //username của khứa anh em muốn stalk vô đây
	discordWebhook = "" //webhook của anh em vô đây
	checkInterval  = 10 * time.Minute
	stateFile      = "/app/state/last_id.txt" // File lưu dấu nghiệp lực
)

var secretPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)(api_key|secret|password|passwd|token|auth_key|private_key)\s*[:=]\s*["'][^"']+["']`),
	regexp.MustCompile(`(?i)\.env`),
	regexp.MustCompile(`-----BEGIN (RSA|OPENSSH) PRIVATE KEY-----`),
}

type Event struct {
	ID      string                `json:"id"`
	Type    string                `json:"type"`
	Repo    struct{ Name string } `json:"repo"`
	Payload struct {
		Commits []struct {
			URL string `json:"url"`
		} `json:"commits"`
	} `json:"payload"`
}

func main() {
	fmt.Printf("🚀 [Trụ Trì Thư] Hộ Pháp khởi động, đang lật sổ rình khứa %s...\n", githubUser)

	for {
		// 1. Đọc ID cuối cùng từ sổ
		lastID := loadLastID()

		events, err := fetchEvents()
		if err == nil && len(events) > 0 {
			if lastID == "" {
				// Lần đầu chạy, ghi nhận mốc rồi thôi
				saveLastID(events[0].ID)
			} else if events[0].ID != lastID {
				// 2. Truy vết những nghiệp lực đã bỏ lỡ
				var missedEvents []Event
				for _, e := range events {
					if e.ID == lastID {
						break
					}
					missedEvents = append(missedEvents, e)
				}

				// 3. Xử lý từ cũ đến mới (đảo ngược mảng)
				for i := len(missedEvents) - 1; i >= 0; i-- {
					processEvent(missedEvents[i])
				}

				// 4. Cập nhật sổ sau khi xử lý xong hết
				saveLastID(events[0].ID)
			}
		}
		time.Sleep(checkInterval)
	}
}

func loadLastID() string {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return ""
	}
	return string(bytes.TrimSpace(data))
}

func saveLastID(id string) {
	_ = os.MkdirAll("/app/state", 0755)
	_ = os.WriteFile(stateFile, []byte(id), 0644)
}

func fetchEvents() ([]Event, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events/public", githubUser))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var events []Event
	json.NewDecoder(resp.Body).Decode(&events)
	return events, nil
}

func processEvent(e Event) {
	msg := fmt.Sprintf("📢 **[Ò Í E] Nghiệp lực mới!**\nKhứa vừa làm gì đó ở: `%s` (Type: `%s`)", e.Repo.Name, e.Type)
	if e.Type == "PushEvent" {
		isLeaking := false
		for _, commit := range e.Payload.Commits {
			if scanCommit(commit.URL) {
				isLeaking = true
				break
			}
		}
		if isLeaking {
			msg += "\n🚨 **CẢNH BÁO:** Phát hiện mùi LEAK SECRETS! Trụ trì ơi vô Hall of Shame gấp!"
		}
	}
	sendDiscord(msg)
}

func scanCommit(commitURL string) bool {
	resp, err := http.Get(commitURL + ".diff")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	for _, pattern := range secretPatterns {
		if pattern.Match(content) {
			return true
		}
	}
	return false
}

func sendDiscord(content string) {
	payload := map[string]string{"content": content}
	data, _ := json.Marshal(payload)
	_, _ = http.Post(discordWebhook, "application/json", bytes.NewBuffer(data))
}
