package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"strings"
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
		/*
			You open the blob file located in the .git/objects directory based on the provided SHA-1 hash. DONE
			You read the contents of the blob file using the ioutil.ReadAll function after opening the file. DONE
			You need to decompress the data using Zlib, which you can do with the zlib.NewReader function. DONE
			After decompressing, you need to extract the actual content by finding the null byte (\0) in the decompressed data and reading everything that comes after it.
		*/
		hash_sha1 := os.Args[3]
		ruta_fichero := fmt.Sprintf(".git/objects/%s/%s", hash_sha1[0:2], hash_sha1[2:])
		fichero, error := os.Open(ruta_fichero)
		if error != nil {
			fmt.Fprintf(os.Stderr, "Error abriendo archivo: %s\n", error)
			os.Exit(1)
		}
		defer fichero.Close()

		zlib_reader, error := zlib.NewReader(fichero)
		if error != nil {
			fmt.Fprintf(os.Stderr, "Error descomprimiendo archivo: %s\n", error)
			os.Exit(1)
		}
		defer zlib_reader.Close()

		contenido_crudo, _ := io.ReadAll(zlib_reader)
		partes_contenido := strings.Split(string(contenido_crudo), "\x00")
		fmt.Print(partes_contenido[1])

	default:
		fmt.Fprintf(os.Stderr, "Comando desconocido %s\n", command)
		os.Exit(1)
	}
}
