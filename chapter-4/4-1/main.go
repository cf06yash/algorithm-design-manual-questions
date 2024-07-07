package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	rating float32
	name   string
	sex    string
}

func CreateNewPlayer(rating float32, name string, sex string) *Player {
	return &Player{
		rating: rating,
		name:   name,
		sex:    sex,
	}
}

func quickSort(players *[]*Player, l, h int) {
	if l < h {
		p := partition(players, l, h)
		quickSort(players, l, p-1)
		quickSort(players, p+1, h)
	}
}

func compare(p1, p2 *Player) int {
	if p1.rating < p2.rating {
		return -1
	}
	return 1
}

func partition(players *[]*Player, l, h int) int {
	var firsthigh int = l
	p := h
	i := l
	for i = l; i < h; i++ {
		if compare((*players)[i], (*players)[p]) < 0 {
			temp := (*players)[i]
			(*players)[i] = (*players)[firsthigh]
			(*players)[firsthigh] = temp
			firsthigh++
		}
	}
	temp := (*players)[p]
	(*players)[p] = (*players)[firsthigh]
	(*players)[firsthigh] = temp
	return firsthigh
}

func main() {
	/*
		The Grinch is given the job of partitioning 2n players into two teams of n
		players each. Each player has a numerical rating that measures how good he or
		she is at the game. The Grinch seeks to divide the players as unfairly as possible,
		so as to create the biggest possible talent imbalance between the teams. Show
		how the Grinch can do the job in O(n log n) time
	*/
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack", "Kathy", "Leo", "Mona", "Nate", "Olivia", "Paul", "Quinn", "Rose", "Steve", "Tina"}
	sexOptions := []string{"NA", "M", "F"}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	players := make([]*Player, 20)

	for i := 0; i < 20; i++ {
		rating := rng.Float32() * 100
		name := names[rng.Intn(len(names))]
		sex := sexOptions[rng.Intn(len(sexOptions))]
		players[i] = CreateNewPlayer(rating, name, sex)
	}

	fmt.Println("Generated Players:")
	for i, player := range players {
		fmt.Printf("Player %d: Name=%s, Rating=%.2f, Sex=%s\n", i+1, player.name, player.rating, player.sex)
	}

	quickSort(&players, 0, len(players)-1)
	team1 := make([]*Player, 10)
	team2 := make([]*Player, 10)

	for i := 0; i < 20; i++ {
		if i < 10 {
			team1[i] = players[i]
		} else {
			team2[i-10] = players[i]
		}
	}

	fmt.Println("Team1 Players:")
	for i, player := range team1 {
		fmt.Printf("Player %d: Name=%s, Rating=%.2f, Sex=%s\n", i+1, player.name, player.rating, player.sex)
	}
	fmt.Println("Team2 Players:")
	for i, player := range team2 {
		fmt.Printf("Player %d: Name=%s, Rating=%.2f, Sex=%s\n", i+1, player.name, player.rating, player.sex)
	}
}
