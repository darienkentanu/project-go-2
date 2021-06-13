package main

import (
	"io"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	const SSH_ADDRESS = "localhost:22"
	const SSH_USERNAME = "user"
	const SSH_PASSWORD = "password"

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(SSH_PASSWORD),
		},
	}

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if client != nil {
		defer client.Close()
	}
	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}

	// session, err := client.NewSession()
	// if session != nil {
	// 	defer session.Close()
	// }
	// if err != nil {
	// 	log.Fatal("failed to create session. " + err.Error())
	// }

	// session.Stdin = os.Stdin
	// session.Stdout = os.Stdout
	// session.Stderr = os.Stderr

	// err = session.Run("ls -l ~/")
	// if err != nil {
	// 	log.Fatal("command execution error. " + err.Error())
	// }

	// var stdout, stderr bytes.Buffer
	// session.Stdout = &stdout
	// session.Stderr = &stderr

	// stdin, err := session.StdinPipe()
	// if err != nil {
	// 	log.Fatal("Error getting stdin pipe. " + err.Error())
	// }

	// err = session.Start("/bin/bash")
	// if err != nil {
	// 	log.Fatal("Error starting bash. " + err.Error())
	// }

	// commands := []string{
	// 	"cd /where/is/the/path",
	// 	"cd src/myproject",
	// 	"ls",
	// 	"exit",
	// }
	// for _, cmd := range commands {
	// 	if _, err = fmt.Fprintln(stdin, cmd); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// err = session.Wait()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// outputErr := stderr.String()
	// fmt.Println("=====================ERROR")
	// fmt.Println(strings.TrimSpace(outputErr))

	// outputString := stdout.String()
	// fmt.Println("=================OUTPUT")
	// fmt.Println(strings.TrimSpace(outputString))

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Failed to create client sftp client. " + err.Error())
	}

	fDestination, err := sftpClient.Create("./test-file2.txt")
	if err != nil {
		log.Fatal("Failed to create destination file. " + err.Error())
	}

	fSource, err := os.Open("/home/acer/Documents/sumber/test-file2.txt")
	if err != nil {
		log.Fatal("Failed to read source file. " + err.Error())
	}

	_, err = io.Copy(fDestination, fSource)
	if err != nil {
		log.Fatal("Failed to copy source file into destination file. " + err.Error())
	}

	log.Println("file copied.")
}
