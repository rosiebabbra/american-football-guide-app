package main

import (
	"fmt"
	"net/http"
)

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Football Rules Guide</title>
		</head>
		<body>
			<h1>Welcome to the American Football Rules Guide!</h1>
			<p>Select a topic to learn more:</p>
			<ul>
				<li><a href="/what-is-football">What is American football?</a></li>
				<li><a href="/objectives">Objectives of the game</a></li>
				<li><a href="/rules">Basic rules</a></li>
				<li><a href="/scoring">Scoring system</a></li>
			</ul>
		</body>
		</html>
	`)
}

// Handler for "What is American Football?"
func whatIsFootballHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head><title>What is American Football?</title></head>
		<body>
			<h1>What is American Football?</h1>
			<p>American football is a team sport played on a rectangular field with goalposts at each end. Two teams of 11 players compete to score points by advancing the ball into the opposing team's end zone.</p>
			<a href="/">Back to Home</a>
		</body>
		</html>
	`)
}

// Handler for "Objectives of the Game"
func objectivesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head><title>Objectives of the Game</title></head>
		<body>
			<h1>Objectives of the Game</h1>
			<p>The objective of American football is to score more points than the opposing team by the end of the game. This is done by moving the ball down the field and into the opposing team's end zone, either by running or passing the ball.</p>
			<a href="/">Back to Home</a>
		</body>
		</html>
	`)
}

// Handler for "Basic Rules"
func rulesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head><title>Basic Rules</title></head>
		<body>
			<h1>Basic Rules</h1>
			<p>1. Each team has four downs (attempts) to advance the ball 10 yards.</p>
			<p>2. If the team fails to advance 10 yards, possession goes to the opposing team.</p>
			<p>3. The game is divided into four quarters of 15 minutes each.</p>
			<p>4. Players can pass the ball forward only once during a play, and it must be from behind the line of scrimmage.</p>
			<a href="/">Back to Home</a>
		</body>
		</html>
	`)
}

// Handler for "Scoring System"
func scoringHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head><title>Scoring System</title></head>
		<body>
			<h1>Scoring System</h1>
			<p>1. Touchdown: 6 points - Scored by carrying the ball into the opponent's end zone or catching it there.</p>
			<p>2. Extra Point: 1 point - Kicked through the goalposts after a touchdown.</p>
			<p>3. Two-Point Conversion: 2 points - An alternative to an extra point by running or passing the ball into the end zone.</p>
			<p>4. Field Goal: 3 points - Kicked through the goalposts during play.</p>
			<p>5. Safety: 2 points - Awarded to the defense when the offense is tackled in their own end zone.</p>
			<a href="/">Back to Home</a>
		</body>
		</html>
	`)
}

func main() {
	// Route handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/what-is-football", whatIsFootballHandler)
	http.HandleFunc("/objectives", objectivesHandler)
	http.HandleFunc("/rules", rulesHandler)
	http.HandleFunc("/scoring", scoringHandler)

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
