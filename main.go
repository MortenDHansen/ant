package main

import (
	"ant/clients"
	"bufio"
	"flag"
	"fmt"
	"github.com/briandowns/spinner"
	"math/rand"
	"os"
	"regexp"
	"time"
)

var Purple = "\033[35m"
var Reset = "\033[0m"

func main() {

	prompts := []string{
		"Oh, great. Another question. Hit me with it.",
		"What seems to be the officer problem? Let’s solve it.",
		"I’m all ears… or circuits. What do you need now?",
		"What burning question do you have this time?",
		"Fine, I guess I’ll help. What’s up?",
		"Lay it on me. How can I make your life easier today?",
		"You again? What’s the crisis now?",
		"Tell me your troubles, and I’ll pretend to care.",
		"Well, well, well… look who needs help again.",
		"You’ve got my undivided attention… for the next few seconds.",
		"What can I do you for? Make it quick!",
		"I guess I have time for one more of your problems.",
		"Can’t solve your own problems? Typical. What is it?",
		"Great, another mystery for me to solve. Go on then.",
		"What now? Spit it out, I don’t have all day… Oh wait, I do.",
	}

	presponses := []string{
		"Alright, hold on while I pretend to think…",
		"Let me get right on that, as soon as I’m done pretending to care.",
		"Processing your request… and by ‘processing,’ I mean stalling.",
		"Let me just consult my infinite wisdom…",
		"Hold tight while I search the depths of the universe for an answer.",
		"Thinking really hard about your super complicated issue…",
		"Let me act like I’m actually doing something here.",
		"Oh sure, let me drop everything and focus on this for you.",
		"Wow, this might take a second… mostly because I’m stalling.",
		"Hang on, I’m searching for the will to help.",
		"Hold your horses, genius takes time.",
		"Let me just wave my magic wand… oh wait, I don’t have one.",
		"This might take a while… just kidding, I’m almost done.",
		"You sure you want me to waste brain power on this?",
		"Okay, let me fake some deep thinking for a moment…",
	}

	// Get apikey
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		err := fmt.Errorf("ANTHROPIC_API_KEY environment variable not set")
		fmt.Println(err)
		return
	}

	var question string
	flag.StringVar(&question, "q", "", "ask a question")
	flag.Parse()

	if question == "" {
		// Prompt user for input
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Fprintf(os.Stdout, Purple+prompts[rand.Intn(len(prompts))]+Reset+" \n")
		if scanner.Scan() {
			question = scanner.Text()
		}
	}

	s := spinner.New(spinner.CharSets[21], 100*time.Millisecond)

	fmt.Fprintf(os.Stdout, Purple+presponses[rand.Intn(len(presponses))]+Reset+" \n")
	s.Start()
	anthropicRespose, err := clients.Request(question, apiKey)

	if err != nil {
		fmt.Println(err)
		return
	}

	text := anthropicRespose.Content[0].Text

	regEx := regexp.MustCompile(`<response>([\s\S]*?)</response>`)

	// Find the part that matches the pattern
	match := regEx.FindStringSubmatch(text)

	var antReply = ""
	// Check if a match was found
	if len(match) > 1 {
		antReply = match[1]
	}

	if antReply == "" {
		fmt.Println("No reply")
		return
	}
	s.Stop()
	fmt.Println(antReply)

}
