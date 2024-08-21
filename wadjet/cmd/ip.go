package cmd

import (
	"net/http"
)

// IP returns the server's public IP address using ipify.org.
func IP() (string, error) {
	// IPv4: https://api.ipify.org
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	buf := make([]byte, 1024)
	n4, err := res.Body.Read(buf)
	if err != nil {
		return "", err
	}

	ip4 := string(buf[:n4])

	// IPv6: https://api6.ipify.org if applicable
	res, err = http.Get("https://api6.ipify.org")
	if err != nil {
		return ip4, nil
	}
	defer res.Body.Close()

	n6, err := res.Body.Read(buf)
	if err != nil {
		return ip4, nil
	}

	return ip4 + "\n" + string(buf[:n6]), nil
}
