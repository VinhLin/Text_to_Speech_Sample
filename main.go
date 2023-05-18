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
	fmt.Println("Convert Text to Speech. Please wait.")
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

// convert_text_to_speech_wav convert text to speech (wav, 8bit(mono), 8000Hz)
// func convert_text_to_speech_wav(text string, lineNum int) error {
// 	// Input MP3 file
// 	mp3FilePath := "input.mp3"

// 	// Output WAV file
// 	wavFilePath := "output.wav"

// 	// Read MP3 file content
// 	// mp3File, err := os.Open(mp3FilePath)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// defer mp3File.Close()

// 	// Read the mp3 file into memory
// 	fileBytes, err := os.ReadFile(mp3FilePath)
// 	if err != nil {
// 		panic("reading my-file.mp3 failed: " + err.Error())
// 	}

// 	// Convert the pure bytes into a reader object that can be used with the mp3 decoder
// 	fileBytesReader := bytes.NewReader(fileBytes)

// 	mp3Decoder := mp3.NewDecoder(fileBytesReader)

// 	// Create WAV file
// 	wavFile, err := os.Create(wavFilePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer wavFile.Close()

// 	// WAV file header
// 	sampleRate := 8000
// 	bitsPerSample := 8
// 	numChannels := 1

// 	header := minimp3.WaveHeader{
// 		SampleRate:    uint32(sampleRate),
// 		BitsPerSample: uint16(bitsPerSample),
// 		NumChannels:   uint16(numChannels),
// 		AudioFormat:   1, // PCM
// 		ByteRate:      uint32(sampleRate * bitsPerSample * numChannels / 8),
// 		BlockAlign:    uint16(bitsPerSample * numChannels / 8),
// 		Subchunk2ID:   [4]byte{'d', 'a', 't', 'a'},
// 		Subchunk2Size: 0, // Update later
// 	}

// 	// Write WAV file header
// 	err = binary.Write(wavFile, binary.LittleEndian, header)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Convert MP3 frames to WAV frames
// 	frameBuffer := make([]byte, 4096) // Adjust buffer size as needed
// 	for {
// 		_, samples, err := mp3Decoder.DecodeNext(frameBuffer)
// 		if err != nil {
// 			break
// 		}
// 		wavFile.Write(samples)
// 	}

// 	// Update WAV file size in the header
// 	wavFileSize, _ := wavFile.Seek(0, os.SEEK_CUR)
// 	wavFileSize -= 8  // Exclude RIFF header
// 	wavFileSize -= 16 // Exclude fmt chunk
// 	header.Subchunk2Size = uint32(wavFileSize)
// 	wavFile.Seek(0, os.SEEK_SET)
// 	binary.Write(wavFile, binary.LittleEndian, header)

// 	fmt.Println("Conversion completed.")

// }
