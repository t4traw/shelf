# shelf - CLI String Storage

Shelfは簡単にcliで文字列を保存・取得できるツールです。キーを指定して標準出力で文字列を保存しておき、キーで文字列を取り出して標準出力に出力します。

[English README](README.md)

## Usage

### Storing Strings

shelfは標準出力から入力を取ります。たとえば改行で区切られたmy_books.txtがあった場合、以下のように文字列を保存できます。

```sh
$ cat my_books.text | shelf set books
```

`cat my_books.text`の出力を取り、それを『books』というキーで保存します。

### Retrieving Strings

保存した文字列を取り出すのも簡単です。設定したキーを使用します。

```sh
$ shelf get books
```

これにより、 booksというキーに保存された文字列を標準出力で取り出します。

### Listing Keys

listコマンドで保存されているすべてのキーのリストを表示できます。

```sh
$ shelf list
```
