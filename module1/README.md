## Module 1 - Algorithmic Analysis

### Tasks
1) Determine time and space complexity of code snippets.
2) Create simple profiler that will:
   * contains implementations of linear and binary search
   * created large sorted array of integers
   * runs both algorithms on the array and measures time taken for each
   * shows results demonstrating the difference in performance between linear and binary search.
   * 



### Results

Profiler Results
```
asad@obsidian:~/Projects/Personal/codingCurriculum/module1$ go run main.go 
Array Size      Linear Search Time   Binary Search Time   Speed Factor (Binary is X times faster)
1000            1.791µs              398ns                4.50                
100000          157.103µs            234ns                671.38              
10000000        16.89489ms           683ns                24736.30            
1000000000      1.919233681s         1.839µs              1043628.97 
```

We see that as the size of the array to search incrases the binary search becomes a biggger and bigger multiple of the speed of the linear search.