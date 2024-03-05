
# Vocabulary Processor

![License](https://img.shields.io/badge/license-MIT-blue.svg)

Vocabulary Processor is a program that automates the processing of vocabulary lists in the .svc format. It generates English articles for users to review and study. The program provides features to extract desired columns from a file, filter words and their meanings, generate new English articles, and create a new .svc file with the extracted vocabulary. Additionally, it allows users to randomly select a specified number of English words from the extracted vocabulary and display them.

## Features

- Automatically extracts desired columns from a file and filters words and their meanings.
- Generates new English articles and saves them as a new .svc file.
- Randomly selects a specified number of English words from the extracted vocabulary and displays them.

## Installation
1. Ensure you have Go installed on your system. If not, you can download and install it from the official Go website: [https://golang.org/](https://golang.org/)

2. Clone this repository:
```shell
git@github.com:Ellioben/wordwander.git
```

3. Navigate to the project directory.

```shell
cd vocabulary-processor
```

4. Place the input .svc file in the `input` folder within the project directory.

5. Compile the source code to generate the binary file.

```shell
go build -o wordwander main.go
```

## Usage

1. Run the compiled binary file.

```shell
./main
```

2. Follow the program prompts to perform the desired operations, such as generating English articles or displaying random words.

3. View the generated English articles and other result files in the `output` folder within the project directory.

## Example

Here's an example of using Vocabulary Processor to generate an English article:

```shell
./vocab-processor 
```
