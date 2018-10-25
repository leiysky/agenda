package main

import (
	"os"

	"github.com/leiysky/agenda/commands"
)

func initConf() {
	if _, err := os.Stat("/var/agenda"); err != nil {
		if err := os.Mkdir("/var/agenda", 0); err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(-1)
		}
	}
	if _, err := os.Stat("/var/agenda/data.json"); err != nil {
		dataFile, err := os.Create("/var/agenda/data.json")
		os.Chmod("/var/agenda/data.json", 0666)
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(-1)
		}
		dataFile.Write([]byte(`{
			"collection": [],
			"session": {}
		}`))
		dataFile.Close()
		return
	}
	dataFile, err := os.OpenFile("/var/agenda/data.json", os.O_RDWR, 0)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(-1)
	}
	dataFile.Close()
}

func main() {
	initConf()
	cmd := commands.NewAgendaCommand()
	cmd.Execute()
}
