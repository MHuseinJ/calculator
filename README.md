# Arithmetic Calculator With Golang

## Description
This application is a arithmetic calculator developed in the Go programming language. It allows users to perform basic arithmetic operations such as addition, subtraction, multiplication, and division on two numbers, and many more.

## Features
- add `n`: Add two numbers together.
- subtract `n`: Subtract one number from another.
- multiply `n`: Multiply two numbers.
- divide `n`: Divide one number by another.
- cancel: Cancel all the calculation and give back sum to zero
- exit: exit the app, clear all the history of command
- abs: giving absolute value of the last value
- neg: giving negative value of the last value
- sqrt: giving square root of last value
- sqr: giving power 2 of last value
- cubert: giving cube root of last value
- cube: giving power 3 of last value
- repeat `n`: repeat last n command

## Prerequisite
- Golang: 1.19.5

## How to Use The Calculator
### 1. Clone this repository to your local machine.
- `git clone git@github.com:MHuseinJ/calculator.git` or `git clone https://github.com/MHuseinJ/calculator.git`
- if you haven't installed golang please do the installment
- if you have please process to step 3

### 2. Installing Go
- Download Go from the official site: https://go.dev/dl/.
- Run the installer for your OS.
- Verify the installation with `$ go version` in your terminal or command prompt.

### 3. Testing App
    $ go test ./...
### 4. Building App
    $ cd cmd
    $ go build .
### 5. Running The App
    $ ./cmd
## Brief About Simple Calculator Apps
our apps consist of 3 packages,
- calculator engine
    - providing features in calculator such as cancel and repeat
    - providing interface to run the math engine for math inquiries
- math engine
    - providing math feature such as add, subtract, multiplu etc
- util service
    - providing general methods that can be used in other package
        - contains
        - readInputFloat
        - readInputString
## Things to Improve
- add more method to math engine to accommodate more features in calculator
- adding native method for sqrt and cubert instead of using math
- adding implementation of reading util method for other input that command line
- unified readInputFloat and readInputString
- move exit method to calculator engine so can be modified how to exit