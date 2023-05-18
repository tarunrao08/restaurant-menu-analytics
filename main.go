package main

import (
	"strconv"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type MenuItem struct {
	FoodmenuID int
	Count      int
}

func main() {
	menuItems := make(map[int]int)
	duplicateItems := make(map[int]int)

	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Printf("Invalid log entry: %s", line)
			return
		}

		eaterID := parseID(fields[0])
		foodmenuID := parseID(fields[1])

		if eaterID == 0 || foodmenuID == 0 {
			log.Printf("Invalid ID in log entry: %s", line)
			return
		}

		if duplicateItems[eaterID] == foodmenuID {
			fmt.Println("Error: Duplicate foodmenu IDs found")
		    return
		} else {
			duplicateItems[eaterID] = foodmenuID
			menuItems[foodmenuID]++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Create a slice of menu items for sorting
	sortedItems := make([]MenuItem, 0, len(menuItems))
	for foodmenuID, eaterID := range menuItems {
		sortedItems = append(sortedItems, MenuItem{
			FoodmenuID: foodmenuID,
			Count:      eaterID,
		})
	}

	// Sort menu items by count in descending order
	sort.Slice(sortedItems, func(i, j int) bool {
		return sortedItems[i].Count > sortedItems[j].Count
	})

	// Print top 3 menu items
	fmt.Println("Top 3 menu items consumed:")
	for i, item := range sortedItems[:3] {
		fmt.Printf("%d. FoodmenuID: %d, Count: %d\n", i+1, item.FoodmenuID, item.Count)
	}
}

func parseID(str string) int {
	id, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("Failed to parse ID: %s", str)
		return 0
	}
	return id
}
