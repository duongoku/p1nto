package p1nto

import (
	"encoding/gob"
	"fmt"
	"os"
)

func SaveData() {
	fmt.Println("Saving data . . .")
	defer fmt.Println("Saved")

	tmp := make(map[string]player)

	for i := range players {
		tmp[i] = *players[i]
	}

	dataFile, err := os.OpenFile("players.gob", os.O_RDWR, 0644)
	defer dataFile.Close()

	if err!=nil {
		fmt.Println(err)
		return
	}

	dataEncoder := gob.NewEncoder(dataFile)
	err = dataEncoder.Encode(tmp)

	if err != nil {
 		fmt.Println(err)
 		return
 	}
}

func LoadData() {
	fmt.Println("Loading data . . .")
	defer fmt.Println("Loaded")

	tmp := make(map[string]player)

	dataFile, err := os.OpenFile("players.gob", os.O_RDWR, 0644)
	defer dataFile.Close()

	if err!=nil {
		fmt.Println(err)
		return
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&tmp)

	if err != nil {
 		fmt.Println(err)
 		return
 	}

	for i := range tmp {
		v := tmp[i]
		players[i] = &v
	}
}
