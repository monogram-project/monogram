# Get Started with the Monogram command-line tool

## Install from pre-built binaries

The basic approach is as follows:

- On GitHub, go to at [the releases page](https://github.com/monogram-project/monogram-go/releases).
- Download the pre-built archive that matches your operating-system and architecture.
- Unpack the archive and move the binaries to somewhere on your $PATH.

### Installing on Linux or MacOS

You can use the following command on Linux or MacOS to download the latest
release. Just replace DEST with the directory where you'd like to put
`monogram`:
```sh
curl --proto '=https' --tlsv1.2 -sSf "http://github.com/monogram-project/monogram-go/install.sh" | bash -s -- --to DEST
```

For example, to install the monogram binaries to `~/.local/bin`:

```sh
# Create the folder if necessary.
mkdir -p ~/.local/bin

# download and extract just to ~/.local/bin
curl --proto '=https' --tlsv1.2 -sSf "http://github.com/monogram-project/monogram-go/install.sh" | bash -s -- --to ~/.local/bin

# add `~/.local/bin` to the paths that your shell searches for executables
# this line should be added to your shell's initialization file,
# e.g. `~/.bashrc` or `~/.zshrc`
export PATH="$PATH:$HOME/.local/bin"

# monogram should now be executable.
monogram --help
```

### Installing from pre-built binaries on Windows

We don't have a neat script as yet.

- Download [the latest zip](https://github.com/monogram-project/monogram-go/releases/download/latest/monogram-windows.zip)
- Unpack the monogram binaries and move them to a folder on your PATH e.g. 'C:\Program Files'.

## Installing from source

To install the Monogram CLI from source you will need a working Go
installation. After that simply enter this command:

```sh
go install github.com/monogram-project/monogram-go/cmd/monogram@latest
```

This command will compile the CLI and install the monogram executable in your
$GOBIN directory. Make sure $GOBIN is in your system's PATH so that you can
invoke monogram directly from the terminal.

## Running directly using docker or podman

If you have docker or podman installed you can use the binary without
any further installation.

```sh
docker run --rm -i sfkleach/monogram:latest [OPTIONS] < STDIN > STDOUT
```

Or for podman:

```sh
podman run --rm -i docker.io/sfkleach/monogram:latest [OPTIONS] < STDIN > STDOUT
```

### Restrictions

Because the `monogram` binary is jailed, you will need to map some of the filing
system if you want to use the `--input` and `--output` options. (Of course you
can avoid this by redirecting standard input/output.)

```sh
podman run --rm -i -v "$(pwd):/mnt" docker.io/sfkleach/monogram:latest --input myfile.mg
```

In the same vein, if you want to use the `--test` web page option then you will
need to map the port appropriately:

```sh
podman run --rm -i -p8080:8080 docker.io/sfkleach/monogram:latest --test
```

