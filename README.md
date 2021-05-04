# Gnomad Webscrapper

## Purpose:
___

The purpose of this repository is to create a cli that can be used to interact with the gnomad graphQL api 
* at the current moment the script can give you the start and stop site of each gene, he gene full name, the chromosome number that the gene is on, and the omim_id for the gene

## Structure of the project:
___
**Outer directories**:

* ***.github***: This directory has two github actions. One for the unit test and one for the linter

* ***bin***: This directory will have the compiled executable

* ***data***: Contains test sets for running the program
    * *test_data.txt*: this is a test input file as a text file 
    * *test_data.xlsx*: this is a test input file as a excel file

* ***src***:
    * *go.mod* & *go.sum*: Two files that keep track of dependencies
    * *main.go*: This is the main file that will run in the program
    * *parse_input.go*: This file reads through the input file and creates a slice of all the gene ids
    * *parse_input_test.go*: This is the unit test file for the parse_input.go file
    * *parser.go*: This is the file that will hand parser through the provided commandline arguments
    * *parser_test.go*: This is the unit test file for the parser.go file
    * *requests.go*: This is the file that handles actually gather all the information from the api. It returns a slice of structs that has the data
    * *requests_test.go*: This is the unit test for the requests.go file 

* ***test_data:***: This directory has two files to use with the unit testing for the program
    * *Book3.xlsx*: This is a test input excel file
    * *test_data.txt*: This is a test input text file used for the unit testing

**Other Files**:
* *.gitignore*: pretty self explanatory
* *LICENSE*: This is under the GNU license
* *Makefile*: This is the make file to automate unit testing, the build process, and removing the previous build
* *README.md*: This is also self explanatory

