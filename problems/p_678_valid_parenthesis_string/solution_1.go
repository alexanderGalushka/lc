package p_678_valid_parenthesis_string

import "math"

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

const openParanth = '('
const closeParanth = ')'
const star = '*'

func checkValidString(s string) bool {
	if s == "" {
		return true
	}
	r := []rune(s)
	if r[0] == ')' {
		return false
	}

	openParanthCount := 0
	closeParanthCount := 0
	starCount := 0

	for _, r := range r {
		if r == openParanth {
			openParanthCount++
		}
		if r == closeParanth {
			closeParanthCount++
		}
		if r == star {
			starCount++
		}
	}

	// balanced
	if openParanthCount == closeParanthCount {
		return true
	}


	absParanthDiff := math.Abs(float64(openParanthCount - closeParanthCount))

	if int(absParanthDiff) < starCount {
		return true
	}

	return false

}
