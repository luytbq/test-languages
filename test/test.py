def minimum_edit_distance(source, target):
    m = len(source)
    n = len(target)

    # Initialize the 2D matrix
    dp = [[0] * (n + 1) for _ in range(m + 1)]

    # Initialize the first row and column
    for i in range(m + 1):
        dp[i][0] = i
    for j in range(n + 1):
        dp[0][j] = j

    # Traverse through the matrix
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if source[i - 1] == target[j - 1]:
                dp[i][j] = dp[i - 1][j - 1]
            else:
                dp[i][j] = min(dp[i - 1][j - 1], dp[i][j - 1], dp[i - 1][j]) + 1

    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            print(dp[i][j])
            print(" ")
        print("\n")

    return dp[m][n]

# Test the minimum_edit_distance function
source_text = "kitten"
target_text = "sitting"
distance = minimum_edit_distance(source_text, target_text)
print("Minimum Edit Distance:", distance)
