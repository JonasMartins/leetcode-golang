package minimumrecolorstogetkconsecutiveblocks

// sliding window
func MinimumRecolors(blocks string, k int) int {
	begin, recolor := 0, 0
	res := k
	for i := 0; i < len(blocks); i += 1 {
		if blocks[i] == 'W' {
			recolor += 1
		}
		if (i-begin)+1 == k {
			if recolor < res {
				res = recolor
			}
			if blocks[begin] == 'W' {
				recolor -= 1
			}
			begin += 1
		}
	}
	return res
}
