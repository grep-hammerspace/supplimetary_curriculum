## Hashtable

- Write  a Hashtable implementation using an array and implement a chaining strategy to handle has collisions
- Use the HashTable to count the frequency of each word in a given text, output should look like a set of {"word: count}

## Solutions
HashTable is defined in ./hastable/HashTable.go.

Frequency counter method is defined in main.go

#### Sample Output
```
asad@basalt:~/Projects/personal/go/supplimetary_curriculum$ go run ./module4/
banana,{banana 3 3649609552}
cherry,{cherry 4 1232791672}
apple,{apple 5 280767167}
Key 'apple' has value 5
Key 'banana' has value 3
Removing by key 'apple' and 'cherry'
banana,{banana 3 3649609552}
---------------------------------------------
Use HashTable to count words in 'Hello Hello,Hello - Fox - Fox Fox Fence Fence Hill.'
Fence,{Fence 2 1648146358}
Hill,{Hill 1 3086529418}
Hello,{Hello 3 4116459851}
Fox,{Fox 3 427698926}
```