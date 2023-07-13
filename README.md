# shelf - CLI String Storage

Shelf is a command-line tool that enables quick storage and retrieval of strings. It's built with simplicity in mind, keeping your data right at your fingertips. Just set, get, and list.

## Install

```sh
$ brew tap t4traw/shelf
$ brew install shelf
```

## Usage

### Storing Strings

Easily store your strings under a specified key. Here's how you do it:

```sh
$ shelf set books --source=my_books.text
```

This command will store the contents of my_books.text under the key books.

### Retrieving Strings

Getting your strings back is just as easy. Use the key you've set:

```sh
$ shelf get books
```

This will retrieve the strings stored under the key books.

### Listing Keys

Want to see what you have stored? Use the list command:

```sh
$ shelf list
```

This will list all the keys you have stored in Shelf.

Shelf is about simplicity and speed. Save time and keep your data organized right in your command line with Shelf.
