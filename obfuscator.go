package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func encodeBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func splitAndReassemble(input string) string {
	var parts []string
	for i := 0; i < len(input); i += 10 {
		end := i + 10
		if end > len(input) {
			end = len(input)
		}
		part := input[i:end]
		parts = append(parts, fmt.Sprintf("\"%s\"", part))
	}
	reassembled := strings.Join(parts, "+")
	return reassembled
}

func addNOPs(input string) string {
	nopCommands := []string{
		"Write-Host 'NOP1'",
		"Write-Host 'NOP2'",
		"Start-Sleep -Milliseconds 10",
		"Write-Host 'NOP3'",
	}
	var output strings.Builder
	for i, char := range input {
		output.WriteRune(char)
		if i%5 == 0 && i < len(input)-1 {
			randomNOP := nopCommands[rand.Intn(len(nopCommands))]
			output.WriteString(";" + randomNOP)
		}
	}
	return output.String()
}

func layerEncoding(input string) string {
	encoded1 := encodeBase64(input)
	reversed1 := reverseString(encoded1)
	encoded2 := encodeBase64(reversed1)
	return encoded2
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func addFakeBranches(input string) string {
	branching := `
if ($false) {
    Write-Host 'This will never run'
} elseif ($true) {
    `
	branchingEnd := `
} else {
    Write-Host 'This will never run either'
}
`
	obfuscated := branching + input + branchingEnd
	return obfuscated
}

func randomizeNames(input string) string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomName := func() string {
		name := make([]byte, 8)
		for i := range name {
			name[i] = chars[rand.Intn(len(chars))]
		}
		return string(name)
	}

	replacements := map[string]string{
		"variableName": randomName(),
		"Write-Host":   randomName(),
		"DownloadFile": randomName(),
		"New-Object":   randomName(),
	}
	for key, value := range replacements {
		input = strings.ReplaceAll(input, key, value)
	}
	return input
}

func obfuscateScriptAdvanced(script string) string {
	encodedScript := encodeBase64(script)
	layeredScript := layerEncoding(encodedScript)
	splitScript := splitAndReassemble(layeredScript)
	withNops := addNOPs(splitScript)
	withBranches := addFakeBranches(withNops)
	finalObfuscatedScript := randomizeNames(withBranches)
	return finalObfuscatedScript
}

func main() {
	scriptPath := flag.String("script", "", "Path to the PowerShell script (.ps1) file")
	flag.Parse()

	if *scriptPath == "" {
		fmt.Println("Please provide a path to the PowerShell script using the -script flag.")
		os.Exit(1)
	}

	scriptContent, err := ioutil.ReadFile(*scriptPath)
	if err != nil {
		fmt.Printf("Error reading script file: %v\n", err)
		os.Exit(1)
	}

	obfuscatedScript := obfuscateScriptAdvanced(string(scriptContent))
	fmt.Println("Advanced Obfuscated PowerShell Script:")
	fmt.Println(obfuscatedScript)
}
