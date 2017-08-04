package scp

import (
	"fmt"
	"github.com/kris-nova/kubicorn/logger"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
)

type SecureCopier struct {
	RemoteUser     string
	RemoteAddress  string
	RemotePort     string
	PrivateKeyPath string
}

func NewSecureCopier(remoteUser, remoteAddress, remotePort, privateKeyPath string) *SecureCopier {
	return &SecureCopier{
		RemoteUser: remoteUser,
		RemoteAddress: remoteAddress,
		RemotePort: remotePort,
		PrivateKeyPath: privateKeyPath,
	}
}

func (s *SecureCopier) ReadBytes(remotePath string) ([]byte, error) {
	pemBytes, err := ioutil.ReadFile(s.PrivateKeyPath)
	if err != nil {
		return nil, err
	}
	signer, err := GetSigner(pemBytes)
	if err != nil {
		return nil, err
	}
	auths := []ssh.AuthMethod{
		ssh.PublicKeys(signer),
	}

	sshConfig := &ssh.ClientConfig{
		User:            s.RemoteUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            auths,
	}
	sshConfig.SetDefaults()
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s",s.RemoteAddress, s.RemotePort), sshConfig)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	r, err := c.Open(remotePath)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *SecureCopier) Write(localPath, remotePath string) error {
	logger.Critical("Write not yet implemented!")
	return nil
}

func GetSigner(pemBytes []byte) (ssh.Signer, error) {
	signerwithoutpassphrase, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		logger.Warning(err.Error())
		fmt.Print("SSH Key Passphrase [none]: ")
		passPhrase, err := terminal.ReadPassword(0)
		if err != nil {
			return nil, err
		}
		signerwithpassphrase, err := ssh.ParsePrivateKeyWithPassphrase(pemBytes, passPhrase)
		if err != nil {
			return nil, err
		} else {
			return signerwithpassphrase, err
		}
	} else {
		return signerwithoutpassphrase, err
	}
}