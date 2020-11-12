package p678

//There is another simple O(n) solution with O(1) space complexity, not very intuitive like the greedy approach, but it's nice to know about it.
//We can rephrase the problem by listing all the valid cases. There are 3 valid cases:
//
//1- There are more open parenthesis but we have enough '*' so we can balance the parenthesis with ')'
//2- There are more close parenthesis but we have enough '*' so we can balance the parenthesis with '('
//3- There are as many '(' than ')' so all parenthesis are balanced, we can ignore the extra '*'
//
//Algorithm: You can parse the String twice, once from left to right by replacing all '*' by '(' and once from right to left by replacing all '*' by ')'.
//For each of the 2 loops, if there's an iteration where you end up with a negative count (SUM['('] - SUM[')'] < 0) then you know the parenthesis were not
//balanced. You can return false. After these 2 checks (2 loops), you know the string is balanced because you've satisfied all the 3 valid cases mentioned above.
//Voila!

// valid:
// (())
// (*())

// invalid:
// )()(
// ())
// (())((())()()(*)(*()(())())())()()((()())((()))(*

func CheckValidString(s string) bool {
	const leftParanth = '('
	const rightParanth = ')'
	const star = '*'

	leftBalance := 0
	for _, rl := range s {
		if rl == leftParanth || rl == star {
			leftBalance++
		} else {
			leftBalance--
		}
		// balance has to be positive
		// since we moving left to right and our first paran is left one
		// we either keep going with the same paran, e.g. (( or close on the valid sequence e.g. () for the string/sequence
		// to be valid
		// if the sequence is being closed (()) (once the balance of 0 met), the new sequence has to begin (of course
		// with the same left paran)
		// or it is the end of the sting

		if leftBalance < 0 {
			return false
		}
	}

	if leftBalance == 0 {
		return true
	}

	rightBalance := 0
	for i:=len(s)-1; i >= 0; i--  {
		if s[i] == rightParanth ||  s[i] == star {
			rightBalance++
		} else {
			rightBalance--
		}
		if rightBalance < 0 {
			return false
		}
	}
	return true
}
