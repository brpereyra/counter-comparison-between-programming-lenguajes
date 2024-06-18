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

| Language   | Time (s)    |
| ---------- | ----------- |
| Go         | 3.587603267 |
| Javascript | 7.421       |
| Python     | 0.483777    |

#### Second Run

| Language   | Time (s)   |
| ---------- | ---------- |
| Go         | 3.23135254 |
| Javascript | 7.899      |
| Python     | 0.468314   |

#### Third Run

| Language   | Time (s)    |
| ---------- | ----------- |
| Go         | 4.284089136 |
| Javascript | 8.331       |
| Python     | 0.476542    |

### Count numbers between 1 and 1000000

#### First Run

| Language   | Time (ms) |
| ---------- | --------- |
| Go         | 0.573     |
| Javascript | 2.621     |
| Python     | 93.486    |

#### Second Run

| Language   | Time (ms) |
| ---------- | --------- |
| Go         | 0.562     |
| Javascript | 2.681     |
| Python     | 85.66     |

#### Third Run

| Language   | Time (ms)         |
| ---------- | ----------------- |
| Go         | 0.527             |
| Javascript | 2.86              |
| Python     | 89.24300000000001 |

## Disclaimer

This comparison is not meant to be a definitive answer to which programming language is the fastest. The results may vary depending on the machine, the compiler, and the code itself. The goal is to provide a general idea of the speed of the languages and to help you choose the right one for your project.
