package main

import "fmt"

func checkAccess(clearanceLevel int) func(string) string {
	if clearanceLevel >= 3 {
		return func(name string) string {
			return "Access Granted Level 3 Vault Open for " + name
		}
	}

	return func(name string) string {
		return "Access Denied Clearance too low for " + name
	}

}

func main() {
	activeUsers := map[string]int{"john": 3, "sam": 2, "alex": 10}

	currentUser := "alex"

	if level, ok := activeUsers[currentUser]; ok {
		securityGuard := checkAccess(level)

		fmt.Println(securityGuard(currentUser))
	} else {
		fmt.Printf("User '%s' not found in the database.\n", currentUser)
	}
}
