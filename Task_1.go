package main

import (
	"errors"
	"fmt"
)

type Battery struct {
	Charge int
}

type MusicS struct {
	Name           string
	IsTurnedOn     bool
	IsMusicPlaying bool
	PlayingMusic   string
	Battery        Battery
	Musics         []string
}

func main() {
	var music int

	station := createMusicStation()

	for {
		option, err := writeMainMenu()

		if err != nil {
			fmt.Println(err)
			continue
		}

		if option == 6 {
			break
		}

		switch option {
		case 1:
			err = station.TurnStationOn()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Station turned on succesfully")

		case 2:
			err = station.TurnStationOff()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Station turned off Sucessfully")

		case 3:
			err = station.TurnOnMusic(music)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Music turned on sucesully!\nPlaying music:", station.Musics[music])
			music, err = station.ModifyMusicMenu(music)

			if err != nil {
				fmt.Println(err)
				continue
			}

		case 4:
			err = station.Battery.ChargeBattery()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Battery charged 10 percents succesfully")

		case 5:
			station.PrintStatus(music)
		}
	}
	fmt.Println("Good Bye :)")
}

func createMusicStation() MusicS {
	return MusicS{
		Name: "JBL 3",
		Battery: Battery{
			Charge: 100,
		},
		Musics: []string{"Mockingbird  Eminem", "Whight night  Valera Saltykov", "Beat It.  MJ"},
	}
}

func (station *MusicS) ModifyMusicMenu(music int) (int, error) {
	if station == nil {
		return music, errors.New("station is nil")
	}
	for {
		option, err := writeMusicMenu()

		if err != nil {
			return music, err
			//continue
		}

		if option == 1 {
			err = station.PauseMusic()
			if err != nil {
				//fmt.Println(err)
				return music, err
			}
			fmt.Println("Music paused succesfully")
			break
		}

		switch option {
		case 2:
			music, err = station.ChangeMusic(music)
			if err != nil {
				//fmt.Println(err)
				//continue
				return music, err
			}

		default:
			fmt.Println("incorrect choose")
		}
	}

	return music, nil
}

func (station *MusicS) TurnStationOn() error {
	if station == nil {
		return errors.New("station is nil")
	}

	if station.IsTurnedOn {
		return errors.New("is already working")
	}

	if station.Battery.Charge < 10 {
		return errors.New("charge is not enough")
	}

	station.IsTurnedOn = true
	return nil
}

func (station *MusicS) TurnStationOff() error {
	if station == nil {
		return errors.New("station is nil")
	}

	if !station.IsTurnedOn {
		return errors.New("station already isnot working")
	}

	if station.IsMusicPlaying == true {
		station.IsMusicPlaying = false
	}

	station.IsTurnedOn = false
	return nil
}

func (battery *Battery) DisChargeStation() error {
	if battery == nil {
		return errors.New("station is nil")
	}

	battery.Charge -= 10
	return nil
}

func (station *MusicS) TurnOnMusic(music int) error {
	if station == nil {
		return errors.New("station is nil")
	}

	if station.IsMusicPlaying {
		return errors.New("music is already playing")
	}

	if !station.IsTurnedOn {
		return errors.New("turn on the station before")
	}

	if station.Battery.Charge <= 10 {
		return errors.New("please, charge your music station")
	}

	err := station.Battery.DisChargeStation()

	if err != nil {
		return fmt.Errorf("cannot turn on music:%w", err)
	}
	station.IsMusicPlaying = true
	station.PlayingMusic = station.Musics[music]
	return nil
}

func (station *MusicS) ChangeMusic(music int) (int, error) {
	if station == nil {
		return music, errors.New("station is nil")
	}

	if !station.IsMusicPlaying {
		return music, errors.New("music isnot playing")
	}

	if station.Battery.Charge <= 10 {
		return music, errors.New("please, charge your music station")
	}

	err := station.Battery.DisChargeStation()

	if err != nil {
		return music, fmt.Errorf("cannot change music:%w", err)
	}

	music++
	if music == 3 {
		music = 0
	}

	station.PlayingMusic = station.Musics[music]
	return music, nil
}

func (station *MusicS) PauseMusic() error {
	if station == nil {
		return errors.New("station is nil")
	}

	if !station.IsMusicPlaying {
		return errors.New("music already isnot working")
	}

	station.IsMusicPlaying = false
	return nil
}

func (battery *Battery) ChargeBattery() error {
	if battery == nil {
		return errors.New("batery is nil")
	}

	if battery.Charge >= 100 {
		return errors.New("batery must have 100 percent or less")
	}

	battery.Charge += 10
	return nil
}

func (station *MusicS) PrintStatus(music int) {
	fmt.Println("Name:", station.Name)
	fmt.Println("Is turned on:", station.IsTurnedOn)
	fmt.Println("Is music playing:", station.IsMusicPlaying)
	fmt.Println("Which music is playing/played last:", station.Musics[music])
	fmt.Println("Charge percent:", station.Battery.Charge)
}

func writeMainMenu() (int, error) {
	var option int

	fmt.Println("----Main Menu----\n1. Turn on\n2. Turn off\n3. Turn on the music\n4. Charge the battery\n5. Print status\n6. Exit")
	_, err := fmt.Scanln(&option)

	if err != nil {
		return 0, errors.New("invalid input")
	}

	return option, nil
}

func writeMusicMenu() (int, error) {
	var option int

	fmt.Println("\n---Music Menu---\n1. Pause\n2. Next music")
	_, err := fmt.Scanln(&option)

	if err != nil {
		return 0, errors.New("invalid input")
	}

	return option, nil
}

/*
1 turn on
2 Turn off
3 Turn on Music
4 next music
5 turn off music
6 pause musik
7chargebatery
8 writeStatus
*/
