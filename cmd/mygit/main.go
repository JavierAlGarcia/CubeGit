package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stderr, "Logs from your program will appear here!\n")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "init":
		for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "Error creando directorio: %s\n", err)
			}
		}

		headFileContents := []byte("ref: refs/heads/main\n")
		if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error escribiendo archivo: %s\n", err)
		}

		fmt.Println("Directorio git inicializado")

	case "cat-file":
		hash_sha1 := os.Args[3]
		ruta_fichero := fmt.Sprintf(".git/objects/%s/%s", hash_sha1[0:2], hash_sha1[2:])
		fichero, error := os.Open(ruta_fichero)
		if error != nil {
			fmt.Fprintf(os.Stderr, "Error abriendo archivo: %s\n", error)
			os.Exit(1)
		}
		defer fichero.Close()

	default:
		fmt.Fprintf(os.Stderr, "Comando desconocido %s\n", command)
		os.Exit(1)
	}
}
