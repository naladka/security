package main

import (
	"strings"
	//"fmt"
	"os"

	"github.com/ivan-bogach/utils"
)

// Allowed - ...
type Allowed struct {
	Members []*Member
}

func (a Allowed) IDExist(id string) bool {
	ids := GetIDs()
	if utils.StringInSlice(id, ids) {
		return true
	} else {
		return false
	}
}

func GetIDs() []string {
	var a []string
	sl, _ := utils.ReadFileLines("collegues.csv")
	for _, s := range(sl) {
		sl := strings.Split(s,",")
		a = append(a, sl[1])
	}
	return a
}

//func (a Allowed) AppendMember(id, name, status, image string) {
//	m := NewMember(id, name, status, image)
//	a.Members = append(a.Members, m)
//}

// GetAllowed
func GetAllowed () Allowed {
	m := []*Member{}
	sl, _ := utils.ReadFileLines("collegues.csv")
	for _, s := range(sl) {
		sl := strings.Split(s,",")
		if len(sl) < 4 {
			continue
		}
		mem := NewMember(sl[0], sl[1], sl[3], "/static/images/" + sl[2])
		m = append(m, mem)
	}
	a := Allowed{Members: m}
	return a	
}

func AddMember(id, name, status, image string){
	text := id + "," + name + "," + status + "," + image + "\n"
	f, err := os.OpenFile("collegues.csv", os.O_APPEND|os.O_WRONLY, 0600)
	
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}


func ChangeMember(id, name, status, image string){
	sl, _ := utils.ReadFileLines("collegues.csv")
	var slice []string
	for _, s := range(sl) {
		if len(s) < 4 {
			continue
		}

		if sl[1] == id {
			s = id + "," + name + "," + status + "," + image + "\n"
		}
		slice = append(slice, s + "\n")
	}

	err := os.Remove("collegues.csv")
	if err != nil {
		panic(err)
	}

    file, err := os.Create("collegues.csv")
	if err != nil {
		panic(err)
	}
    defer file.Close()



	for _, s := range(slice){
		f, err := os.OpenFile("collegues.csv", os.O_APPEND|os.O_WRONLY, 0600)
		if _, err = f.WriteString(s); err != nil {
			panic(err)
		}
	}
	
}

func (a Allowed) GetName(id string) string {
	var name string
	sl, _ := utils.ReadFileLines("collegues.csv")
	for _, s := range(sl) {
		sl := strings.Split(s,",")
		if sl[1] == id {
			name = sl[0]
		}
	}
	return name
}

func (a Allowed) GetImage(id string) string {
	var image string
	sl, _ := utils.ReadFileLines("collegues.csv")
	for _, s := range(sl) {
		sl := strings.Split(s,",")
		if sl[1] == id {
			image = sl[2]
		}
	}
	return image
}

func (a Allowed) GetPosition(id string) string {
	var position string
	sl, _ := utils.ReadFileLines("collegues.csv")
	for _, s := range(sl) {
		sl := strings.Split(s,",")
		if sl[1] == id {
			position = sl[3]
		}
	}
	return position
}

// GetMembers ...
func GetMembers() map[string]string {
	s, _ := utils.ReadFileLines("collegues.csv")
	m := utils.StringToMap(strings.Join(s, "\n"), "\n", ",")
	return m
}

// Member - ...
type Member struct {
	Id   string
	Name string
	Status string
	Image string
}

func NewMember(id, name, status, image string) *Member {
	return &Member{Id: id, Name: name, Status: status, Image: image}
}





