# Speed comparison between programming languages

## Introduction

This comparison is based on the time it takes to run a simple program that counts the number of prime numbers between 1 and 1000000. The programs are written in different programming languages and run on the same machine. The goal is to compare the speed of the languages and see which one is the fastest.

## Languages

The following programming languages are included in the comparison:

- Go
- Python
- Javascript (node)

## environment

- OS: MacOS Sonoma 14.5
- CPU: 2.2 GHz 6-Core Intel Core i7
- Memory: 32 GB 3200 MHz DDR4
- Radeon Pro 555X 4 GB

## Results

all the programs were run 3 times and the average time was calculated.

### Print numbers between 1 and 1000000

#### First Run

| Language   | Result 1 Time (s) | Result 2 Time (s) | Result 3 Time (s) | Average Time (s) |
| ---------- | ----------------- | ----------------- | ----------------- | ---------------- |
| Go         | 3.587603267       | 3.23135254        | 4.284089136       | 3.701014981      |
| Javascript | 7.421             | 7.899             | 8.331             | 7.883666667      |
| Python     | 0.483777          | 0.468314          | 0.476542          | 0.476211         |

### Count numbers between 1 and 1000000

#### First Run

| Language   | Result 1 Time (ms) | Result 2 time (ms) | Result 3 time (ms) | Average Time (ms) |
| ---------- | ------------------ | ------------------ | ------------------ | ----------------- |
| Go         | 0.573              | 0.562              | 0.527              | 0.554             |
| Javascript | 2.621              | 2.681              | 2.86               | 2.720666667       |
| Python     | 93.486             | 85.66              | 89.243             | 89.463            |

## Disclaimer

This comparison is not meant to be a definitive answer to which programming language is the fastest. The results may vary depending on the machine, the compiler, and the code itself. The goal is to provide a general idea of the speed of the languages and to help you choose the right one for your project.
