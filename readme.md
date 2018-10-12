## Online Test Implementation
This is my collections of online test question about basic programming test from various sources 

## How to Run
```bash
# clone this project
git clone https://github.com/jojoarianto/online-test-1.git
cd online-test-1

# run 
go run q3.go
```

## Q4

There is a well-known puzzle called Word Search that involves looking for words in a grid of letters. The words are given in a list and can appear in the grid horizontally, vertically, or diagonally in any direction. In this task, you should implement a solver for word search. You will be given grids and a word to search for, and you have to find how many times that word comes out in the grid. Words that are spelled the same backwards and forwards, also known as palindromes, will not be given, so you don’t need to worry about words that match in the exact same spot in two different directions.

Input:

The first line is the number of test cases T. Each test case will have two numbers N and M, each on their own line given in that order. Following that is N lines of M lowercase letters each representing the grid of letters. Lastly, a word W is given that you must look for.

Output:

For each test case, output one line of the form “Case C: X” (without the quotes), where C is the case number (starting from 1), and X is how many times the word W appeared in the grid.

Constraints:

```json
1 ≤ T ≤ 100
1 ≤ N ≤ 100
1 ≤ M ≤ 100
1 ≤ length(W) ≤ 100
```

Sample Input:

```json
3
3
4
catt
aata
tatc
cat
5
5
gogog
ooooo
godog
ooooo
gogog
dog
2
8
bananana
kalibrrr
nana
```

Sample Output:
```json
Case 1: 4
Case 2: 8
Case 3: 4
```


## Q5

You are presented with a map of a kingdom. Empty land on this map is depicted as ‘.’ (without the quotes), and mountains are depicted by ‘#’. This kingdom has factions whose armies are represented by lowercase letters in the map. Two armies of the same letter belong to the same faction.

Armies can go up, down, left, and right, but cannot travel through mountains. A region is an area enclosed by mountains. From this it can be deduced that armies cannot travel between different regions. A region is said to be controlled by a faction if it’s the only faction with an army in that region. If there are at least two armies from different factions in a region, then that region is said to be contested. Your task is to determine how many regions each faction controls and how many contested regions there are.

Input:

The first line is the number of test cases T. Each test case will have two numbers N and M, each on their own line given in that order. Following that is N lines of M characters that represent the map.

Output:

For each test case, output one line of the form “Case C:” (without the quotes), where C is the case number (starting from 1). On the following lines, output the factions that appear in the grid in alphabetical order if the faction controls at least one region and how many regions that letter controls. Factions that don’t control any regions should not be in the output. After this, print one last line of the form “contested K” where K is the number of regions that contain at least two distinct letters.

See the sample output below for details.

Constraints:

```json
1 ≤ T ≤ 100
1 ≤ N ≤ 100
1 ≤ M ≤ 100
```

Sample Input:
```json
2
2
2
.#
#a
12
15
###########....
#.......c.###..
####......#.#..
.#.########.#..
##...#..b#..#..
#.a.#...#...###
####.#.#d###..#
......#...e#xx#
.#....#########
.#.x..#..#.....
.######.c#.....
......####.....
```
Sample Output:

Case 1:
```json
a 1
contested 0
Case 2:
a 1
b 1
c 2
x 2
contested 1
```

## Q6
This code should print the numbers from 1 to 100 including both ends, one per line, except it prints "Fizz" when the number is divisible by 3, "Buzz", when the number is divisible by 5, and "FizzBuzz" when the number is divisible by 3 and 5. Something is wrong with it. Find and fix it by modifying one line!

Select your language below:

```php
<?php
for ($i=1; $i<100; $i++) {
    if ($i % 15 == 0) {
        printf("FizzBuzz\n");
    } else if ($i % 5 == 0) {
        printf("Buzz\n");
    } else if ($i % 3 == 0) {
        printf("Fizz\n");
    } else{
        printf("%d\n", $i);
    }
}
?>
```

## Q7

The binary search algorithm allows you to find the position of a specific value in a sorted array of data very quickly. It works by checking the middle of that array. If the middle is smaller than the value we're looking for, then we can be sure that the value we're looking for must be in the larger half of the array. Otherwise it's in the smaller half. We can repeatedly apply that idea, with each repetition cutting away half of the items we're looking through. Below is an implementation of binary search, but something is wrong with it. Find it and fix it by modifying one line!

Input:

The input consists of the sorted array of length 10,000, followed by 10,000 queries. Each integer is given in its own line (there are 20,000 lines in total).

It is guaranteed that there are no duplicates in the array.

Output:

For every query value, output one line of output containing a single integer: the index that matches the query value, or -1 if the value is not in the array.

Select your language below:

```php
<?php
    function binary_search($array, $value, $low, $high) {
        if ($high < $low) {
            return -1;
        } else {
            $mid = ($low + $high)/2;
            if ($array[$mid] > $value) {
                return binary_search($array, $value, $low, $mid);
            } else if ($array[$mid] < $value) {
                return binary_search($array, $value, $mid+1, $high);
            } else {
                return $mid;
            }
        }
    }
    $array = array();
    for ($i=0; $i<10000; $i++) {
        fscanf(STDIN, "%d", $array[$i]);
    }
    for ($i=0; $i<10000; $i++) {
        fscanf(STDIN, "%d", $value);
        $answer = binary_search($array, $value, 0, 9999);
        printf("%d\n", $answer);
    }
?>
```

## Q8

This code prints all the permutations of the string lexicographically. Something is wrong with it. Find and fix it by modifying or adding one line!

Input:

The input consists of a single line containing a string of lowercase characters with no spaces in between. Its length is at most 7 characters, and its characters are sorted lexicographically.

Output:

All permutations of the string printed one in each line, listed lexicographically.

Sample Input:

```json
abc
```

Sample Output:
```json
abc
acb
bac
bca
cab
cba
```
Select your language below:

C/C++
Java
Javascript
PHP
Python
Ruby

```php
<?php
function permutations() {
    global $running;
    global $characters;
    global $bitmask;
    if (count($running) == count($characters)) {
        printf("%s\n", implode($running));
    } else {
        for ($i=0; $i<count($characters); $i++) {
            if ( (($bitmask>>$i)&1) == 0 ) {
                array_push($running, $characters[$i]);
                $bitmask |= (1<<$i);
                permutations();
                array_pop($running);
            }
        }
    }
}

fscanf(STDIN, '%s', $raw_input);
$characters = str_split($raw_input);
$running = array();
$bitmask = 0;
permutations();
?>
```

## T1

This is a demo task.

Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].

## T2

A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

Write a function:

function solution($N);

that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5. Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..2,147,483,647].

## T3

An array A consisting of N integers is given. Rotation of the array means that each element is shifted right by one index, and the last element of the array is moved to the first place. For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7] (elements are shifted right by one index and 6 is moved to the first place).

The goal is to rotate array A K times; that is, each element of A will be shifted to the right K times.

Write a function:

class Solution { public int[] solution(int[] A, int K); }

that, given an array A consisting of N integers and an integer K, returns the array A rotated K times.

For example, given

    A = [3, 8, 9, 7, 6]
    K = 3
the function should return [9, 7, 6, 3, 8]. Three rotations were made:

    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]
For another example, given

    A = [0, 0, 0]
    K = 1
the function should return [0, 0, 0]

Given

    A = [1, 2, 3, 4]
    K = 4
the function should return [1, 2, 3, 4]

Assume that:

N and K are integers within the range [0..100];
each element of array A is an integer within the range [−1,000..1,000].
In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.