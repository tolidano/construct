package construct

// See https://github.com/parsiya/Hacking-with-Go/blob/master/content/04.4.md
certCheck := &ssh.CertChecker{
	IsHostAuthority: hostAuthCallback(),
	IsRevoked:       certCallback(s),
	HostKeyFallback: hostCallback(s),
}

