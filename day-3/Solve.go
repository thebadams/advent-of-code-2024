package day3

func Solve(path string) (part1, part2 int) {
	data := readFile(path)
	uncorrupt := UncorruptData(data.String())
	ins1, ins2 := ReadInstructions(uncorrupt)
	part2 = RunInstructions(ins2)
	part1 = RunInstructions(ins1)
	return part1, part2
}
