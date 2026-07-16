# SSH Website

This is a different version of my website that you can ssh into!

## Setup

Pretty simple, just clone the repo

```bash
git clone https://github.com/kashsuks/ssh-kashyap.git
cd ssh-kashyap
```

Then clean and build the package

```bash
go mod tidy
go build -o ssh-kashyap-server .
```

Generate a host key (once)

```bash
ssh-keygen -t ed25519 -f host_key -N ""
```

Run the project!

```bash
./ssh-kashyap-server
```

### Expected Output

You should see

```bash
twice ssh server listening on 0.0.0.0:2222
```

And if you were to try it from another terminal

```bash
ssh visitor@your.nest.hackclub.app -p 2222
```

Any username and password is accepted since there isn't an auth system and meant for public access
