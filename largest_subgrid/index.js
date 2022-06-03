function findMaxMatrixSize(arr, K)
{
 
    let i, j;
 
    // N size of rows and M size of cols
    let n = arr.length;
    let m = arr[0].length;
 
    // To store the prefix sum of matrix
    let sum=[];
    for(i =0;i<n+1;i++){
       sum[i] = [];
       for(j =0;j<m+1;j++){
           sum[i][j] = 0;
       }
    }
    console.log(sum)
    // Create Prefix Sum
    for ( i = 0; i <= n; i++) {
 
        // Traverse each rows
        for (j = 0; j <= m; j++) {
 
            if (i == 0 || j == 0) {
                sum[i][j] = 0;
                continue;
            }
 
            // Update the prefix sum
            // till index i x j
            sum[i][j] = arr[i - 1][j - 1] + sum[i - 1][j]
                        + sum[i][j - 1] - sum[i - 1][j - 1];
        }
    }
 
    // To store the maximum size of
    // matrix with sum <= K
    let ans = 0;
 
    // Traverse the sum matrix
    for (i = 1; i <= n; i++) {
 
        for (j = 1; j <= m; j++) {
 
            // Index out of bound
            if (i + ans - 1 > n || j + ans - 1 > m)
                break;
 
            let mid, lo = ans;
 
            // Maximum possible size
            // of matrix
            let hi = Math.min(n - i + 1, m - j + 1);
 
            // Binary Search
            while (lo < hi) {
 
                // Find middle index
                mid = Math.floor((hi + lo + 1) / 2);
 
                // Check whether sum <= K
                // or not
                // If Yes check for other
                // half of the search
                if (sum[i + mid - 1][j + mid - 1]
                        + sum[i - 1][j - 1]
                        - sum[i + mid - 1][j - 1]
                        - sum[i - 1][j + mid - 1]
                    <= K) {
                    lo = mid;
                }
 
                // Else check it in first
                // half
                else {
                    hi = mid - 1;
                }
            }
 
            // Update the maximum size matrix
            ans = Math.max(ans, lo);
        }
    }
 
    // Print the final answer
    console.log(ans)
}
 
// Driver Code
 
    let arr = [[ 1, 1, 3, 2, 4, 3, 2 ],
            [ 1, 1, 3, 2, 4, 3, 2 ],
            [ 1, 1, 3, 2, 4, 3, 2 ]];
 
    // Given target sum
    let K = 4;
 
    // Function Call
    findMaxMatrixSize(arr, K);