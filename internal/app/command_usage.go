package app

import "fmt"

func CommandUsage() {
	fmt.Println("\nUsage: [command] [input filepath] [output filepath]")
	fmt.Println("Commands:\n	- encrypt : encrypt the file at the input filepath.\n	- decrypt : decrypt the file at the input filepath.\n	- help : display usage information.")
	fmt.Println("Example: gocrypt encrypt file.txt file.txt\n.gcrypt will be appended to the filename.")
}
