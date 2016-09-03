package service

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/zyfdegh/hiupdate/server/entity"
	"github.com/zyfdegh/hiupdate/server/orm"
	"github.com/zyfdegh/hiupdate/server/util"
)

// FilePersonsList is path to persons.list
const FilePersonsList = "./persons.list"

// GetAllPersons read persons.list file and return persons
func GetAllPersons() (persons []entity.Person, err error) {
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

// GetForgotPersons returns those who forgot writing record
func GetForgotPersons(date string) (respForgotPersons *entity.RespForgotPersons, err error) {
	var forgotPersons []string
	respForgotPersons = &entity.RespForgotPersons{}
	if len(date) != 8 {
		log.Printf("invalid date: %s", date)
		respForgotPersons.Success = false
		respForgotPersons.ErrNo = http.StatusBadRequest
		respForgotPersons.Errmsg = "invalid date"
		return
	}

	allPersons, err := GetAllPersons()
	if err != nil {
		log.Printf("get all persons error: %v", err)
		respForgotPersons.Success = false
		respForgotPersons.ErrNo = http.StatusInternalServerError
		respForgotPersons.Errmsg = "get all persons error"
		return
	}

	for _, person := range allPersons {
		id := util.GenerateID(person.Name, date)
		record, _ := orm.QueryRecord(id)
		if record == nil || len(record.Content.Done) == 0 || len(record.Content.Todo) == 0 {
			forgotPersons = append(forgotPersons, person.Name)
		}
	}
	respForgotPersons.Success = true
	respForgotPersons.Data = forgotPersons
	log.Println(forgotPersons)
	return
}
