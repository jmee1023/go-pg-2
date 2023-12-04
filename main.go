package main

import "fmt"

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessage := []string{}
	for _, message := range messages {
		formattedMessage = append(formattedMessage, formatter(message))
	}
	return formattedMessage
}

func addSignature(message string) string {
	return message + " Kind Regard."
}

func addGreeting(message string) string {
	return "Hello " + message
}

func test(messages []string, formatter func(string) string) {
	defer fmt.Println("====================================")
	formattedMessages := getFormattedMessages(messages, formatter)
	if len(formattedMessages) != len(messages) {
		fmt.Println("The number of messages returned is incorrect.")
		return
	}
	for i, message := range messages {
		formatted := formattedMessages[i]
		fmt.Printf(" * %s -> %s\n", message, formatted)
	}
}

func main() {
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addSignature)
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addGreeting)
}
