package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Document struct {
	ID      int
	Owner   string
	Content string
}

type User struct {
	Username string
	Role     string
}

type AuditLog struct {
	Time     time.Time
	Username string
	Role     string
	DocID    int
	Result   string
}

var logs []AuditLog
var attemptCount int

const maxActions = 5

func main() {
	documents := []Document{
		{ID: 1, Owner: "userA", Content: "A's secret"},
		{ID: 2, Owner: "userB", Content: "B's secret"},
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Secure In-Memory Access System ===")

	user := readUser(reader)
	actions := 0

	for {
		if actions >= maxActions {
			fmt.Println("\nSession expired. Please login again.")
			user = readUser(reader)
			actions = 0
		}

		fmt.Println("\nChoose an option:")
		fmt.Println("1. View documents")
		fmt.Println("2. Access document")
		fmt.Println("3. View audit logs (admin only)")
		fmt.Println("4. Exit")
		fmt.Println("5. Switch user")

		fmt.Print("Choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {

		case "1":
			printDocuments(documents)
			actions++

		case "2":
			accessDocument(reader, user, documents)
			actions++

		case "3":
			if user.Role == "admin" {
				printAuditLogs()
			} else {
				fmt.Println(`{"status":"denied","reason":"admin only"}`)
			}
			actions++

		case "4":
			fmt.Println("\nTotal access attempts:", attemptCount)
			fmt.Println("Exiting system.")
			return

		case "5":
			fmt.Println("\nSwitching user...")
			user = readUser(reader)
			actions = 0

		default:
			fmt.Println("Invalid choice")
		}
	}
}

// ---------- Helper Functions ----------

func readUser(reader *bufio.Reader) User {
	fmt.Print("\nUsername: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Role (user/admin): ")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(role)

	fmt.Println("Session started for", username, "| Role:", role)
	return User{Username: username, Role: role}
}

func printDocuments(docs []Document) {
	fmt.Println("\nDocuments:")
	for _, d := range docs {
		fmt.Printf("ID:%d | Owner:%s\n", d.ID, d.Owner)
	}
}

func accessDocument(reader *bufio.Reader, user User, docs []Document) {
	fmt.Print("Enter document ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	attemptCount++

	for _, d := range docs {
		if d.ID == id {
			if user.Role == "admin" {
				logAccess(user, id, "granted-admin")
				fmt.Println(`{"status":"granted","content":"` + d.Content + `"}`)
				return
			}

			if d.Owner == user.Username {
				logAccess(user, id, "granted-owner")
				fmt.Println(`{"status":"granted","content":"` + d.Content + `"}`)
				return
			}

			logAccess(user, id, "denied")
			fmt.Println(`{"status":"denied","reason":"not owner"}`)
			return
		}
	}

	logAccess(user, id, "not-found")
	fmt.Println(`{"status":"error","reason":"document not found"}`)
}

func logAccess(user User, docID int, result string) {
	logs = append(logs, AuditLog{
		Time:     time.Now(),
		Username: user.Username,
		Role:     user.Role,
		DocID:    docID,
		Result:   result,
	})
}

func printAuditLogs() {
	fmt.Println("\nAudit Logs:")
	for _, l := range logs {
		fmt.Printf("[%s] %s (%s) -> Doc %d : %s\n",
			l.Time.Format(time.RFC822),
			l.Username,
			l.Role,
			l.DocID,
			l.Result,
		)
	}
}
