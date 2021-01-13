package LCS

func LCS(str1, str2 string) string {
	if len(str1) == 0 || len(str2) == 0 {
		return ""
	}
	dp := make([][]int, len(str1)+1)
	for i := range dp {
		dp[i] = make([]int, len(str2)+1)
	}

	var max, maxM int

	for m := 1; m < len(str1)+1; m++ {
		for n := 1; n < len(str2)+1; n++ {
			if str1[m-1] == str2[n-1] {
				dp[m][n] = dp[m-1][n-1] + 1
				if dp[m][n] > max {
					max = dp[m][n]
					maxM = m
				}
			}
		}
	}
	return str1[maxM-max : maxM]
}
