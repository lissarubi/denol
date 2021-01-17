# Denol Language

Denol Language is a ~~seriously~~ programming language. You write Denol code, and this code is parsed to PHP code.

![Denol Logo](assets/logo.png)

# Install

## Windows

get `denol.exe` with this method:

```
git clone https://github.com/edersonferreira/denol
cd denol/builds
```

now, move `denol.exe` to a certain place (binaries directory). I know the path to this place? No! But I fell you know Kappa

## Linux

### Using DEB

Run these commands:

```
cd /tmp
git clone https://github.com/edersonferreira/denol
cd denol
sudo dpkg -i builds/denol.deb
```

### Using Golang

- [Install Golang](https://golang.org/)
- Run `go install github.com/edersonferreira/denol`

# Using Denol

Denol has a simple syntax, and a extended syntax to PHP.

To compile Denol Code to PHP code, only use `denol` binary, and the file(s) to compile, like this:

```
$ denol helloworld.denol
```

## Hello World

The Denol Hello World is:

> hello.denol

```
ola("Hello World")
```

To execute, run:

```
denol hello.denol
```

This command will parse Denol code to PHP code, and save it in a directory named `build`

And now, you can execute the PHP code with this command:

```
php build/hello.php
```

# Denol keywords

| Denol        |        PHP |
| :----------- | ---------: |
| `!import`    |   `import` |
| `amain`      |  `finally` |
| `cancelar`   |    `break` |
| `day`        |       `do` |
| `deno`       |     `exit` |
| `esporro`    | `function` |
| `mandabala`  | `continue` |
| `naosalvou`  |     `else` |
| `ola`        |     `echo` |
| `outrosalve` |  `else if` |
| `padraozao`  |  `default` |
| `pamonhosa`  |      `for` |
| `pediu`      |  `require` |
| `pega`       |    `catch` |
| `printa`     |    `print` |
| `salve`      |       `if` |
| `sepa`       |     `case` |
| `tenta`      |      `try` |
| `teste`      |   `switch` |
| `veio`       |   `return` |
| `zoeira`     |    `while` |

# Code Examples

Assign 10 to `a` and log `a`
```
a = 10
ola(a)
```

Assign 20 to `b` and if `b` is 20 log a "b is 20"

```
b = 20

salve (b == 20){
  ola ("b is 20")
}
```

Create a `alo` function, execute alo, assing 10 to `a` and log `a`

```
esporro alo(){
  ola("alo")
}
alo()

a = 10

ola(a)
```

# Denol VS PHP

- Denol is funnier than PHP
- Denol has "Deno" in the name, but PHP not.
- Denol has a Deninho as logo, PHP not.
- And `amain` ops, Finally, Denol is Denol, and PHP not is Denol.
