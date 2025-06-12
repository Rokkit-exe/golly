package ui

import (
	"fmt"

	"github.com/Rokkit-exe/golly/models"
	"github.com/charmbracelet/glamour"
	"github.com/kyokomi/emoji/v2"
)

type UI struct {
	FullResponse string
	Query        string
	Renderer     *glamour.TermRenderer
}

func (u *UI) PrintAI(streamCh <-chan *models.ChatResponseChunk, errCh <-chan error) {
	for chunk := range streamCh {
		if chunk.Message.Content != "" {
			u.FullResponse += chunk.Message.Content
			// Use glamour for pretty printing
			u.FullResponse = emoji.Sprint(u.FullResponse) // Convert emojis
			rendered, err := u.Renderer.Render(u.FullResponse)
			if err != nil {
				fmt.Println("Error rendering message:", err)
				continue
			}
			fmt.Print("\033[H\033[2J") // Optional: clears screen for live update
			fmt.Print(rendered)
		}
	}

	if err, ok := <-errCh; ok {
		fmt.Printf("Error: %v\n", err)
	}
}

func (u *UI) PrintUser(query string) {
	u.Query += "\n\n" + emoji.Sprint(":user: ") + query + "\n"
	rendered, err := u.Renderer.Render(u.FullResponse)
	if err != nil {
		fmt.Println("Error rendering user message: ", err)
		return
	}
	fmt.Print("\033[H\033[2J") // Optional: clears screen for live update
	fmt.Print(rendered)
}

func (u *UI) Scan() (string, bool) {
	message, err := u.Renderer.Render("## Type your message (or 'exit' to quit): ")
	if err != nil {
		fmt.Println("Error rendering user message: ", err)
		return "", false
	}
	fmt.Print(message)
	u.Query = ""
	u.FullResponse = ""
	fmt.Scanf("%s", u.Query)
	if u.Query == "exit" {
		return "", false
	}
	return u.Query, true
}

func (u *UI) Clear() {
	u.FullResponse = ""
	u.Query = ""
	fmt.Print("\033[H\033[2J") // Clears the screen
}

func (u *UI) PrintEndOfMessage() {
	u.Renderer.Render("---")
}
