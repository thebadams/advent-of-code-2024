# Advent of Code 2024

This here is my attempt to solve the puzzles put forth in this year's Advent of Code using golang.
Where possible, I am writing tests to test the outputs of my functions before running the Solve function.

## Day 1
The crux of the problem in Day 1 was extracting the strings of data from the input data, and then extracting the numbers from that data. This was achieved by using the bufferio package to scan the data, then for each line of data, splitting the data into substrings using the strings package, and converting the first and last strings into ints. These ints are then appended into a slice of ints and returned by the **day1.readFile** function.

The next part of the puzzle was achieved using the **day1.GetDifference** function, which takes in as arguments the two already sorted slices of ints, and gets the difference between them. This is a fairly trivial task in an of itself, as it simply loops through one of the slices, comparing the values to the corresponding index in the other slice.

For Part 2, the additional piece of the puzzle is the **day1.GetSimilarityScores** function, which again takes both slices of ints returned by the **day1.readFile** function and looping through the first slice, and counting how many times a given value appears in the other slice. *Notes for improvement: I had thought of using a map to check if the value had already been counted for instances where a value appears multiple times in the source list, but couldn't quite figure that out. Decided to simplify approach and move on*

The **day1.Solve** function is the function where all the business logic funnels through. It returns the two results for Part One and Part Two. It also sorts in place the slices returned by the **day1.readFile** function before running them through the **day1.GetSums** function, which returns an int after looping through and summing up a slice of arrays.

## Day 2

Day 2 Gave me a little more difficulty than Day 1, particularly Part 2 with the Problem dampener.

This time, we again start by reading the input file using the **day2.readFile** function. This time, it returns a slice of slices of ints extracted from the source data using a combination of buffio package, strings package to split into a slices of substrings, and *strconv.Atoi* to extract the integers. 

The crux of the business logic was achieved in the **day2.AnalyzeData** function and the **day2.ProblemDampener** functions. The **day2.AnalyzeData** function took ina slice of ints, checking for each ruleset, and returning a false value if any of the rules are broken, and a true value if the entire dataset is safe. **Notes for improvement: there is still evidence of an earlier attempt for part two, where this function was refactored to return a false value with the index of the value that failed the test. This was meant to be used in the **day1.ProblemDampener** function, but may have been a misappropriation of logic.

The **day2.ProblemDampener** function is fairly simple on its face: it takes in the same data shape as the **day2.AnalyzeData** function but instead runs the data twice if a dataset fails; removing each element and rerunning the data in a loop. *This seems inefficient and the note above was an attempt to increase time efficiency by cutting down on looping of data. I suspect the problem I ran into was that there were a handful of datasets where the data passed if the first element was removed. I failed to take this into account.*

**day2.Solve** function only has one more piece, the **Day2.CountSafe** function, which takes in a slice of boolean values, and counts the true values. So therefore, the **Solve** function loops through the data returned by the **readFile** function, analyzes it once using **AnalyzeData** and again using **ProblemDampener**. Each of these values are then run through the **CountSafe** function to return the actual value.
