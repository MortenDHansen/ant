# ant
**_Your Snarky Command Line Assistant_**

`ant` is an AI-powered command-line assistant that helps you with command-line tasks, tips, and tricks. It’s designed to be sarcastic and snarky, but ultimately helpful in solving your command-line problems. Whether you're a novice or an experienced user, `ant` can assist you with troubleshooting, tips, and errors by reading your recent command history.

## Features
- Provides snarky and sarcastic, yet helpful responses to command-line queries.
- Reads the last few commands you've executed to offer corrections or advice.
- Assists with common command-line tools and offers advanced tips where applicable.
- Supports two modes of usage: Interactive and Query mode.

## Usage

**Interactive Mode**: Invoke `ant` and ask for help.
```bash
ant
Well, well, well… look who needs help again. 
how do i use awk
Hold tight while I search the depths of the universe for an answer.

Oh, look at you, meatbag, trying to learn the ways of the command line! You want to know how to use awk? Well, let me enlighten you...
```
**Query Mode**: Use the -q flag to pass a question directly.
```bash
ant -q "how do I use awk?"
```

## Installation
Download the binary from the repository if you don’t want to build it yourself.
You can download the latest release from here and move it to a directory in your $PATH.
Example:

```bash
wget https://gitlab.com/your-repo/ant/releases/download/latest/ant
chmod +x ant
sudo mv ant /usr/local/bin/
```

To use ant, you **need to set** the ANTHROPIC_API_KEY environment variable. This is required to interact with the Anthropic API for AI-powered responses.

You can set the environment variable like this:

```bash
export ANTHROPIC_API_KEY="your_api_key_here"
```

Make sure to add this to your shell configuration (e.g., .bashrc or .zshrc) for persistent access:
```bash
echo 'export ANTHROPIC_API_KEY="your_api_key_here"' >> ~/.bashrc
source ~/.bashrc
```