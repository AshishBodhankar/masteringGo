package main

import (
	"fmt"
	"strings"
	//"unicode"
	"errors"
)

func countDistinctWords(messages []string) int {

	frequency := map[string]bool{}
	concat := []string{}

	for _, message := range messages {
		lowerSlice := strings.Fields(strings.ToLower(message))

		concat = append(concat, lowerSlice...)
	}

	for _, word := range concat {
		frequency[word] = true
	}

	return len(frequency)
}

/*--------------------------------------------------------------------------------------------------------------------*/
func getNames(length int) []string {
	return []string{
		"Grant", "Eduardo", "Peter", "Matthew", "Matthew", "Matthew", "Peter", "Peter", "Henry", "Parker",
		"Parker", "Parker", "Collin", "Hayden", "George", "Bradley", "Mitchell", "Devon", "Ricardo", "Shawn",
		"Taylor", "Nicolas", "Gregory", "Francisco", "Liam", "Kaleb", "Preston", "Erik", "Alexis", "Owen",
		"Omar", "Diego", "Dustin", "Corey", "Fernando", "Clayton", "Carter", "Ivan", "Jaden", "Javier",
		"Alec", "Johnathan", "Scott", "Manuel", "Cristian", "Alan", "Raymond", "Brett", "Max", "Drew",
		"Andres", "Gage", "Mario", "Dawson", "Dillon", "Cesar", "Wesley", "Levi", "Jakob", "Chandler",
		"Martin", "Malik", "Edgar", "Sergio", "Trenton", "Josiah", "Nolan", "Marco", "Drew", "Peyton",
		"Harrison", "Drew", "Hector", "Micah", "Roberto", "Drew", "Brady", "Erick", "Conner", "Jonah",
		"Casey", "Jayden", "Edwin", "Emmanuel", "Andre", "Phillip", "Brayden", "Landon", "Giovanni", "Bailey",
		"Ronald", "Braden", "Damian", "Donovan", "Ruben", "Frank", "Gerardo", "Pedro", "Andy", "Chance",
		"Abraham", "Calvin", "Trey", "Cade", "Donald", "Derrick", "Payton", "Darius", "Enrique", "Keith",
		"Raul", "Jaylen", "Troy", "Jonathon", "Cory", "Marc", "Eli", "Skyler", "Rafael", "Trent",
		"Griffin", "Colby", "Johnny", "Chad", "Armando", "Kobe", "Caden", "Marcos", "Cooper", "Elias",
		"Brenden", "Israel", "Avery", "Zane", "Zane", "Zane", "Zane", "Dante", "Josue", "Zackary",
		"Allen", "Philip", "Mathew", "Dennis", "Leonardo", "Ashton", "Philip", "Philip", "Philip", "Julio",
		"Miles", "Damien", "Ty", "Gustavo", "Drake", "Jaime", "Simon", "Jerry", "Curtis", "Kameron",
		"Lance", "Brock", "Bryson", "Alberto", "Dominick", "Jimmy", "Kaden", "Douglas", "Gary", "Brennan",
		"Zachery", "Randy", "Louis", "Larry", "Nickolas", "Albert", "Tony", "Fabian", "Keegan", "Saul",
		"Danny", "Tucker", "Myles", "Damon", "Arturo", "Corbin", "Deandre", "Ricky", "Kristopher", "Lane",
		"Pablo", "Darren", "Jarrett", "Zion", "Alfredo", "Micheal", "Angelo", "Carl", "Oliver", "Kyler",
		"Tommy", "Walter", "Dallas", "Jace", "Quinn", "Theodore", "Grayson", "Lorenzo", "Joe", "Arthur",
		"Bryant", "Roman", "Brent", "Russell", "Ramon", "Lawrence", "Moises", "Aiden", "Quentin", "Jay",
		"Tyrese", "Tristen", "Emanuel", "Salvador", "Terry", "Morgan", "Jeffery", "Esteban", "Tyson", "Braxton",
		"Branden", "Marvin", "Brody", "Craig", "Ismael", "Rodney", "Isiah", "Marshall", "Maurice", "Ernesto",
		"Emilio", "Brendon", "Kody", "Eddie", "Malachi", "Abel", "Keaton", "Jon", "Shaun", "Skylar",
		"Ezekiel", "Nikolas", "Santiago", "Kendall", "Axel", "Camden", "Trevon", "Bobby", "Conor", "Jamal",
		"Lukas", "Malcolm", "Zackery", "Jayson", "Javon", "Roger", "Reginald", "Zachariah", "Desmond", "Felix",
		"Johnathon", "Dean", "Quinton", "Ali", "Davis", "Gerald", "Rodrigo", "Demetrius", "Billy", "Rene",
		"Reece", "Kelvin", "Leo", "Justice", "Chris", "Guillermo", "Matthew", "Matthew", "Matthew", "Kevon",
		"Steve", "Frederick", "Clay", "Weston", "Dorian", "Hugo", "Roy", "Orlando", "Terrance", "😊",
		"Kai", "Khalil", "Khalil", "Khalil", "Graham", "Noel", "Willie", "Nathanael", "Terrell",
	}[:length]
}

func getNameCounts(names []string) map[rune]map[string]int {
	result := make(map[rune]map[string]int)
	for _, name := range names {
		firstChar := name[0]
		mm, ok := result[rune(firstChar)]
		if !ok {
			mm = make(map[string]int)
			result[rune(firstChar)] = mm
		}
		mm[name]++
	}
	return result
}

/*--------------------------------------------------------------------------------------------------------------------*/

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, User := range messagedUsers {
		_, ok := validUsers[User]
		if ok {
			validUsers[User]++
		}
	}
}

/*--------------------------------------------------------------------------------------------------------------------*/

type Key struct {
	Path, Country string
}

var hits = map[Key]int{}

func add(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]

	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}

/*
map keys may be of any type that is comparable.
The language spec defines this precisely, but in short, comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types.

Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.
*/

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	user_, ok := users[name]

	if !ok {
		return false, errors.New("not found")
	} else if !user_.scheduledForDeletion {
		return false, nil
	} else {
		delete(users, name)
		return true, nil
	}
}

type user struct {
	name                 string
	phoneNumber          int
	scheduledForDeletion bool
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {

	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	//mapp := make(map[string]user)
	mapp := map[string]user{}

	for i := 0; i < len(names); i++ {
		mapp[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]}
	}
	return mapp, nil
}

func main() {
	names := []string{"Ashish", "Vaishu", "Aie"}
	numbers := []int{5514448052, 5716397308, 9948157504}

	fmt.Println(getUserMap(names, numbers))

	names_ := map[string]int{}
	missingNames := []string{}

	if _, ok := names_["Denna"]; !ok {
		missingNames = append(missingNames, "Denna")
	}
	fmt.Println(missingNames)
	_names := getNames(10)
	fmt.Println(getNameCounts(_names))

	test_ := []string{"Could you give me a number crunch real quick?", "Looks like we have a 32.33% (repeating of course) percentage of survival."}
	fmt.Println(countDistinctWords(test_))
}
