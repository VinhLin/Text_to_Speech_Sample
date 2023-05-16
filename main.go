package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"

	htgotts "github.com/hegedustibor/htgo-tts"
	handlers "github.com/hegedustibor/htgo-tts/handlers"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

func main() {
	fmt.Println("This tool help read text file and convert it to speech (Vietnamese).")
	fmt.Println("Please write your content in List_Text.txt")
	fmt.Println("Maximum for one line is 150 character.")
	convert_text_file("List_Text.txt")
}

// convert_text_to_speech convert text to speech (mp3)
func convert_text_to_speech(text string, lineNum int) error {
	speech := htgotts.Speech{Folder: "audio", Language: voices.Vietnamese, Handler: &handlers.MPlayer{}}

	charCount := utf8.RuneCountInString(text)
	// fmt.Println("Number of characters:", charCount)
	if charCount > 150 {
		return fmt.Errorf("Overload character (Maximum %v)", 150)
	}

	name_audio := fmt.Sprintf("Test_audio_%v", lineNum)
	fileName, err := speech.CreateSpeechFile(text, name_audio)
	if err != nil {
		fmt.Println(err)
	}
	speech.PlaySpeechFile(fileName)

	return nil
}

// convert_text_file read content on file text and convert text to speech
func convert_text_file(file_text string) {
	// Open the file
	file, err := os.Open(file_text)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line until the desired line is reached
	lineNumber := 1
	for scanner.Scan() {
		// Get the content of the first line
		line := scanner.Text()
		fmt.Println(line)
		fmt.Println("Convert Text to Speech. Please wait.")
		err := convert_text_to_speech(line, lineNumber)
		if err != nil {
			fmt.Println(err)
			break
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
