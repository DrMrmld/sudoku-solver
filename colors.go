package main

const gray string = "\033[0;90m"
const white string = ""
const reset string = "\033[0m"

var colors [9][9]string = [9][9]string{
	{white, white, white, gray, gray, gray, white, white, white},
	{white, white, white, gray, gray, gray, white, white, white},
	{white, white, white, gray, gray, gray, white, white, white},
	{gray, gray, gray, white, white, white, gray, gray, gray},
	{gray, gray, gray, white, white, white, gray, gray, gray},
	{gray, gray, gray, white, white, white, gray, gray, gray},
	{white, white, white, gray, gray, gray, white, white, white},
	{white, white, white, gray, gray, gray, white, white, white},
	{white, white, white, gray, gray, gray, white, white, white},
}
