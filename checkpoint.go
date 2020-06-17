package main

import (
	"strings"
	"fmt"

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

func (a Allowed) AppendMember(id, name, status, image string) {
	m := NewMember(id, name, status, image)
	a.Members = append(a.Members, m)
}

// GetAllowed
func GetAllowed () Allowed {
	var a Allowed
	sl, _ := utils.ReadFileLines("collegues.csv")
	for _, s := range(sl) {
		sl := strings.Split(s,",")
		a.AppendMember(sl[0], sl[1], sl[3], sl[2])
	}
	fmt.Println(len(a.Members))
	return a	
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
	id   string
	name string
	status string
	image string
}


func NewMember(id, name, status, image string) *Member {
	return &Member{id: id, name: name, status: status, image: image}
}





