package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Football Rules Guide</title>
	<script src="https://cdn.tailwindcss.com"></script>
	<style>
		[aria-expanded="true"] ~ .answer {
			display: block;
		}
	</style>
</head>
<body class="bg-gray-100 text-gray-900">

	<div class="max-w-2xl mx-auto p-6">
		<h1 class="text-3xl font-bold text-center mb-6">American Football Rules Guide 🏈</h1>

		<div class="space-y-4">
			<!-- Question 1 -->
			<div class="bg-white shadow-md rounded-lg">
				<button aria-expanded="false" class="w-full text-left p-4 font-semibold text-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
					onclick="toggleAccordion(this)">
					What is American football?
				</button>
				<div class="answer px-4 pb-4 hidden">
					<p>American football is a team sport played on a rectangular field with goalposts at each end. Two teams of 11 players compete to score points by advancing the ball into the opposing team's end zone.</p>
				</div>
			</div>

			<!-- Question 2 -->
			<div class="bg-white shadow-md rounded-lg">
				<button aria-expanded="false" class="w-full text-left p-4 font-semibold text-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
					onclick="toggleAccordion(this)">
					Objectives of the game
				</button>
				<div class="answer px-4 pb-4 hidden">
					<p>The objective of American football is to score more points than the opposing team by the end of the game. This is done by moving the ball down the field and into the opposing team's end zone, either by running or passing the ball.</p>
				</div>
			</div>

			<!-- Question 3 -->
			<div class="bg-white shadow-md rounded-lg">
				<button aria-expanded="false" class="w-full text-left p-4 font-semibold text-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
					onclick="toggleAccordion(this)">
					Basic rules
				</button>
				<div class="answer px-4 pb-4 hidden">
					<p>1. Each team has four downs (attempts) to advance the ball 10 yards.</p>
					<p>2. If the team fails to advance 10 yards, possession goes to the opposing team.</p>
					<p>3. The game is divided into four quarters of 15 minutes each.</p>
					<p>4. Players can pass the ball forward only once during a play, and it must be from behind the line of scrimmage.</p>
				</div>
			</div>

			<!-- Question 4 -->
			<div class="bg-white shadow-md rounded-lg">
				<button aria-expanded="false" class="w-full text-left p-4 font-semibold text-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
					onclick="toggleAccordion(this)">
					Scoring system
				</button>
				<div class="answer px-4 pb-4 hidden">
					<p>1. Touchdown: 6 points - Scored by carrying the ball into the opponent's end zone or catching it there.</p>
					<p>2. Extra Point: 1 point - Kicked through the goalposts after a touchdown.</p>
					<p>3. Two-Point Conversion: 2 points - An alternative to an extra point by running or passing the ball into the end zone.</p>
					<p>4. Field Goal: 3 points - Kicked through the goalposts during play.</p>
					<p>5. Safety: 2 points - Awarded to the defense when the offense is tackled in their own end zone.</p>
				</div>
			</div>
		</div>
	</div>

	<script>
		function toggleAccordion(button) {
			// Toggle the aria-expanded attribute
			const expanded = button.getAttribute('aria-expanded') === 'true';
			button.setAttribute('aria-expanded', !expanded);

			// Find the answer div and toggle its visibility
			const answer = button.nextElementSibling;
			if (expanded) {
				answer.classList.add('hidden');
			} else {
				answer.classList.remove('hidden');
			}
		}
	</script>

</body>
</html>
	`)
}

func main() {
	// Serve the web page
	http.HandleFunc("/", mainHandler)

	// Start the server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
