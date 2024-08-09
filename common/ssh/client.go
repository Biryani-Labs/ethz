package ssh

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	sshMutex sync.Mutex
)

const (
	maxRetries        = 5
	retryDelay        = 3 * time.Second
	connectionTimeout = 30 * time.Second
)

func GetSSHConnection(ctx context.Context, user, host, port, password string) (*ssh.Client, error) {
	sshMutex.Lock()
	defer sshMutex.Unlock()

	for retry := 1; retry <= maxRetries; retry++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			sshClient, err := connect(ctx, user, host, port, password)
			if err == nil {
				return sshClient, nil
			}
			log.Printf("SSH connection failed (Attempt %d/%d): %v", retry, maxRetries, err)
			if retry < maxRetries {
				time.Sleep(retryDelay)
			}
		}
	}

	return nil, fmt.Errorf("failed to establish an SSH connection after %d attempts", maxRetries)
}

func connect(ctx context.Context, user, host, port, password string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         connectionTimeout,
	}

	addr := net.JoinHostPort(host, port)

	var sshConn net.Conn
	var err error

	dialCtx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	sshConn, err = (&net.Dialer{}).DialContext(dialCtx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	c, chans, reqs, err := ssh.NewClientConn(sshConn, addr, config)
	if err != nil {
		sshConn.Close()
		return nil, fmt.Errorf("failed to create client connection: %w", err)
	}

	return ssh.NewClient(c, chans, reqs), nil
}
