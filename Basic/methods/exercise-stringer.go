package main

import "fmt"

// Definição do tipo IPAddr
type IPAddr [4]byte

// Implementação do método String() para IPAddr
func (ip IPAddr) String() string {
	// Retorna o endereço IP no formato "X.X.X.X"
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	// Um map com alguns IPs
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	// Exibe os nomes e seus respectivos IPs formatados
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
