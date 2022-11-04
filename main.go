package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

var home string
var dbPath string
var profiles []Profile

func main() {
	home, _ = os.UserHomeDir()
	dbPath = fmt.Sprintf("%s/%s", home, ".vs-code-profiles/profiles.data")
	Load()
	switch os.Args[1] {
	case "--create":
		profile := CreateProfile()
		fmt.Println("---")
		profile.Print()
		profiles = append(profiles, *profile)
		Save()
	case "--list":
		List()
	case "--update":
		profile := Find(os.Args[2])
		if profile != nil {
			profile.Update()
			fmt.Println("---")
			profile.Print()
		}
		Save()
	case "--delete":
		res := []Profile{}
		for _, profile := range profiles {
			if profile.Alias != os.Args[2] {
				res = append(res, profile)
			}
		}
		profiles = res
		Save()
	default:
		profile := Find(os.Args[1])
		if profile != nil {
			profile.Execute(os.Args[2])
		}
	}
}

type Profile struct {
	Name          string
	Alias         string
	UserDataDir   string
	ExtensionsDir string
}

func (profile *Profile) Print() {
	fmt.Printf("name: %s\nalias: %s\nuser data directory: %s\nextensions directory: %s\n", profile.Name, profile.Alias, profile.UserDataDir, profile.ExtensionsDir)
}

func (profile *Profile) Execute(path string) {
	command := fmt.Sprintf("code --user-data-dir %s --extensions-dir %s %s", profile.UserDataDir, profile.ExtensionsDir, path)
	exec.Command("bash", "-c", command).Run()
}

func CreateProfile() *Profile {
	profile := Profile{}
	fmt.Println("Enter new profile name:")
	fmt.Scanln(&profile.Name)
	fmt.Println("Enter new profile alias:")
	fmt.Scanln(&profile.Alias)
	var dirBuff string
	tempDir := fmt.Sprintf("%s/.vs-code-profiles/%s/data", home, profile.Alias)
	fmt.Printf("Enter new profile user date directory (default: %s):\n", tempDir)
	fmt.Scanln(&dirBuff)
	if dirBuff == "" {
		profile.UserDataDir = tempDir
	} else {
		profile.UserDataDir = dirBuff
	}
	tempDir = fmt.Sprintf("%s/.vs-code-profiles/%s/extensions", home, profile.Alias)
	fmt.Printf("Enter new profile extensions directory (default: %s):\n", tempDir)
	fmt.Scanln(&dirBuff)
	if dirBuff == "" {
		profile.ExtensionsDir = tempDir
	} else {
		profile.ExtensionsDir = dirBuff
	}

	return &profile
}

func Find(alias string) *Profile {
	for _, profile := range profiles {
		if profile.Alias == alias {
			return &profile
		}
	}
	fmt.Println("No profiles was found")
	return nil
}

func Load() {
	db, _ := os.Open(dbPath)
	var buff []byte
	fmt.Fscan(db, &buff)
	json.Unmarshal(buff, &profiles)
}

func Save() {
	j, _ := json.Marshal(profiles)
	db, _ := os.Create(dbPath)
	db.Write(j)
	db.Close()
}

func List() {
	for _, profile := range profiles {
		fmt.Println("---")
		profile.Print()
		fmt.Println("---")
	}
}

func (profile *Profile) Update() {
	buff := ""
	fmt.Printf("Enter new profile name (current %s):\n", profile.Name)
	fmt.Scanln(&buff)
	if buff != "" {
		profile.Name = buff
		buff = ""
	}
	fmt.Printf("Enter new profile alias (current %s):\n", profile.Alias)
	fmt.Scanln(&buff)
	if buff != "" {
		profile.Alias = buff
		buff = ""
	}
	fmt.Printf("Enter new profile user date directory (current: %s):\n", profile.UserDataDir)
	fmt.Scanln(&buff)
	if buff != "" {
		profile.UserDataDir = buff
		buff = ""
	}
	fmt.Printf("Enter new profile extensions directory (current: %s):\n", profile.ExtensionsDir)
	fmt.Scanln(&buff)
	if buff != "" {
		profile.ExtensionsDir = buff
		buff = ""
	}
}
