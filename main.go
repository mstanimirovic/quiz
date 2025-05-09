package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"
	"time"
)

type Question struct {
	Text          string
	Answers       []string
	CorrectAnswer string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("koriscenje: go run main.go <putanja fajla>\n\tili ./quiz <putanja fajla>")
		os.Exit(1)
	}

	fileContent := GetFileContent(os.Args[1])
	questions := ParseFileContent(fileContent)
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	var limitQuestions int

	fmt.Printf("Broj pitanja: ")
	fmt.Scan(&limitQuestions)

	startTime := time.Now()
	correctAnswers := AskTheQuestions(questions, limitQuestions)

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	secs := int(timeTaken.Seconds())
	mins := secs / 60
	secs %= 60

	fmt.Printf("Rezultat: %v/%v\nVreme: %02d:%02d\n", correctAnswers, limitQuestions, mins, secs)
}

// returns the number of right answers
func AskTheQuestions(questions []Question, limit int) uint {
	var input string
	var answer string
	var correctCount uint

	for idx, question := range questions {
		if idx == limit {
			break
		}
		fmt.Printf("%v. %v\n\n", idx+1, question.Text)

		rand.Shuffle(len(question.Answers), func(i, j int) {
			question.Answers[i], question.Answers[j] = question.Answers[j], question.Answers[i]
		})
		for idx, ans := range question.Answers {
			fmt.Printf("%c) %v\n", 'A'+idx, ans)
		}

		answer = ""
		for answer == "" {
			fmt.Print("\n> ")
			fmt.Scan(&input)

			if input[0] >= 'A' && input[0] <= 'D' {
				answer = question.Answers[int(input[0]-'A')]
			} else if input[0] >= 'a' && input[0] <= 'd' {
				answer = question.Answers[int(input[0]-'a')]
			} else if input[0] >= '1' && input[0] <= '4' {
				answer = question.Answers[int(input[0]-'1')]
			} else {
				fmt.Println("Unesi broj (1-4) ili slovo (a-d)")
			}
		}

		if answer != question.CorrectAnswer {
			fmt.Printf("Netacno!\nTacan odgovor: %s\n", question.CorrectAnswer)
		} else {
			fmt.Println("Tacno!")
			correctCount++
		}

		fmt.Println()
	}
	return correctCount
}

func ParseFileContent(content []string) []Question {
	var questions []Question
	var length int = 0
	var answerIdx uint
	var flagQuestion bool = true

	for _, line := range content {
		if line == "" {
			flagQuestion = true
			continue
		}

		if flagQuestion {
			questions = append(questions, Question{Text: Trim(line)})
			answerIdx = 0
			length++
			flagQuestion = false
			continue
		}

		if line[0] == ' ' || line[0] == '\t' {
			questions[length-1].CorrectAnswer = Trim(line)
		}
		questions[length-1].Answers = append(questions[length-1].Answers, Trim(line))
		answerIdx++
	}

	return questions
}

func Trim(str string) string {
	return strings.Trim(str, "\n ")
}

func GetFileContent(path string) []string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(bytes), "\n")
}
