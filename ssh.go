package construct

import (
	"net"

	"golang.org/x/crypto/ssh"
)

type SshClient struct {
	Username string,
	password string,
	Server string,
	Port int,
	TerminalSession *ssh.Session,
}

var (
	certCheck := &ssh.CertChecker{
		IsHostAuthority: hostAuthCallback(),
		IsRevoked:       certCallback(s),
		HostKeyFallback: hostCallback(s),
	}
)

// See https://github.com/parsiya/Hacking-with-Go/blob/master/content/04.4.md
func New(server string, port int) (*SshClient) {
	return &SshClient{
		Server: server,
		Port: port,
	}
}

func (*SshClient s) connect() (error) {
	combinedAddress := net.JoinnHostPort(s.Server, s.Port)
	sshConn, err := ssh.Dial("tcp", t, config)
    if err != nil {
        return err
    }
    session, err := sshConn.NewSession()
    if err != nil {
        return err
    }
    session.Stdout = os.Stdout
    session.Stderr = os.Stderr
    input, err := session.StdinPipe()
    if err != nil {
    	return err
    }
    termModes := ssh.TerminalModes{
    	ssh.ECHO: 0, // Disable echo
    }
    err = session.RequestPty("vt220", 40, 80, termModes)
    if err != nil {
        return err
    }
    err = session.Shell()
    if err != nil {
        return err
    }
    s.TerminalSession = session
    return nil
}

func (*SshClient s) setUserAndPassword(username string, password string) {
	s.Username = username
	s.password = password
}

func (*SshClient s) configWithPassword() (*ssh.ClientConfig) {
	config := &ssh.ClientConfig{
        // Username
        User: s.username,
        // Each config must have one AuthMethod. In this case we use password
        Auth: []ssh.AuthMethod{
            ssh.Password(s.password),
        },
        // This callback function validates the server.
        // Danger! We are ignoring host info
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    return config
}
