# shelf - CLI String Storage

Shelf is a command-line tool that enables quick storage and retrieval of strings. It's built with simplicity in mind, keeping your data right at your fingertips. Just set, get, and list.

[日本語版READMEはこちら](README_ja.md)

## Install

```sh
$ brew install t4traw/shelf/shelf
```

## Usage

### Storing Strings

Shelf takes its input from the standard output. For instance, if you have a text file called my_books.txt with line-separated entries, you can store the strings like so:

```sh
$ cat my_books.text | shelf set books
```

This command takes the output from cat my_books.text and stores it under the key books.

### Retrieving Strings

Retrieving your stored strings is just as straightforward. You simply use the key you've set:

```sh
$ shelf get books
```

This command retrieves the strings stored under the key books and prints them to the standard output.

### Listing Keys

You can display a list of all stored keys using the list command:

```sh
$ shelf list
```

This will show you all the keys that have stored strings.
