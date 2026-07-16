package main

import (
    "io"
    "log"
    "net"

    "github.com/gliberlabs/ssh"
)

const port = "2222"

// welcome screen
const welcome = `
⠀⠀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠀⠀⠀⠀⠀
⣠⣴⢟⣶⠂⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣴⣾⠀⠀⠀⠀⠀
⠀⠾⠋⠙⠂⢀⣠⣴⣶⣶⣶⣾⣿⣿⣟⣹⣾⠇⠀⠀⠀⠀⠀
⠀⠀⠀⢠⣶⣿⠿⠿⣿⣿⢿⡻⢿⡟⠿⠟⠁⠀⠀⠀⠀⠀⡰
⠀⠀⣰⡟⠁⠀⠀⣸⣿⣿⡿⠀⠀⠀⠀⠀⣀⡀⢀⡀⠀⠘⠁
⠀⢰⣿⡅⠀⠀⠀⣾⣿⣷⠁⠀⠀⠀⠀⢿⣿⣿⣮⣿⢢⠃⠀
⠀⢸⣿⣿⣄⠀⢠⣿⣿⠇⢠⣿⣿⣿⡄⠀⣿⣿⡏⠀⢸⠀⠀
⠀⠈⠛⠋⠀⠀⢸⣿⣿⠀⢾⣿⢩⣿⠇⠀⣿⣿⠏⠀⠸⡄⠀
⠀⠀⠀⠀⠀⠀⠸⣿⣿⡀⠘⣿⣾⠏⠁⣰⣿⡟⠀⠀⢀⣣⠀
⠀⠀⠀⠀⠀⠀⠀⠙⠿⣿⡶⠟⢿⣶⣾⡿⠋⠀⣯⠚⢹⡀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⡲⡴⠊⠁
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠁⠀⠀
`

func main() {
    server := &ssh.Server{
        Addr: ":" + port,
	Handler: func(s ssh.Session) {
	    io.WriteString(s, welcome)

	    buf := make([]byte, 1)
	    s.Read(buf)

	    io.WriteString(s, "r\ngoodbye.\r\n")
	    s.Close()
	},
	// accept any username/password -- this is public, read-only
	PasswordHandler: func(ctx ssh.Context, password string) bool {
	    return true
	},
	PublicKeyHandler: func(ctx ssh.Context, key ssh.PublicKey) bool {
	    return true
	},
    }

    // host key that is generated once and reused
    hostKeyOption := ssh.HostKeyFile("host_key")
    if err := hostKeyOption(server); err != nil {
        log.Fatalf("failed to load host key: %v", err)
    }

    log.Printf("ssh-kashyap server listening on %s", net.JoinHostPort("0.0.0.0", port))
    log.Fatal(server.ListenAndServe())
}
