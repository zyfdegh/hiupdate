package service

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/zyfdegh/hiupdate/server/entity"
)

// FilePersonsList is path to persons.list
const FilePersonsList = "./persons.list"

// GetAllPersons read persons.list file and return persons
func GetAllPersons() (persons []entity.Person, err error) {
	persons = make([]entity.Person, 20)
	lines, err := readPersonsList(FilePersonsList)
	if err != nil {
		log.Printf("read persons list error: %v", err)
		return
	}
	for _, line := range lines {
		arr := strings.Split(line, " ")
		if len(arr) != 2 {
			log.Println("parse line error")
			continue
		}
		person := entity.Person{}
		person.Name = arr[0]
		person.Team = arr[1]
		persons = append(persons, person)
	}
	return
}

// Read file
func readPersonsList(file string) (lines []string, err error) {
	_, err = os.Stat(file)
	if err != nil {
		log.Printf("stat file[%s] error: %v", file, err)
		return
	}

	f, err := os.Open(file)
	if err != nil {
		log.Printf("open file[%s] error: %v", file, err)
		return
	}

	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, e := rd.ReadString('\n')
		lines = append(lines, line)
		if e != nil {
			break
		}
	}
	return
}
