package parse

import (
	"bufio"
	"fmt"
	"go-rest-api-logs/internal/interfaces"
	"go-rest-api-logs/internal/models"
	"io/ioutil"
	"os"
	"strings"
)

//SquidParseCreateFolder Crea la estructura de carpetas el del directorio del proyecto
func SquidParseCreateFolder() {
	var resultEmpresas = interfaces.GetAllEmpresas()

	for _, index := range resultEmpresas {
		os.Mkdir("./uploads/" + index.Identificador, 0755)

		squidParseReadFile(index.Identificador)
	}
}

func squidParseReadFile(idempresa string) {
	files, err := ioutil.ReadDir("./uploads/" + idempresa)
	if err != nil {
		fmt.Println("error el directorio el fichero " + err.Error())
	}

	for _, f := range files {
		fmt.Println(f.Name())

		file, err := os.Open("./uploads/" + idempresa + "/" + f.Name())
		if err != nil {
			fmt.Println("error el leer el fichero " + err.Error())
		}
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			parse := strings.Fields(scanner.Text())
			//fmt.Println(parse[3])

			interfaces.CreateCode(&models.Codes{Code: parse[3]})
		}

		file.Close()
		// eliminando el fichero
		err = os.Remove("./uploads/" + idempresa + "/" + f.Name())
		if err != nil {
			fmt.Println("No se pudo eliminar el fichero " + idempresa + "/" + f.Name())
		}
	}

	/*scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	file.Close()*/
}
