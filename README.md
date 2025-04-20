Command line interface for google gemini.

This is a pet project of mine. Still an early version.

Assuming a linux environment: to install ensure first that a recent version of `go` is installed on your system, and then run the following

```bash
go install github.com/gianluca-peri/cligem@latest
```

also be sure to add the following lines to your `.bashrc` file:

```bash
export YOUR_API_KEY
export PATH="$PATH:$HOME/go/bin"
```

where of course you must substitute `YOUR_API_KEY` with the correct value. Right now you can get a Google Gemini API key for free (up to a limit).

After the setup you can simply run a chat by typing `cligem` in your terminal.