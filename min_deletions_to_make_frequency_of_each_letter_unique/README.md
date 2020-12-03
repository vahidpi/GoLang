# Min Deletions to Make Frequency of Each Letter Unique

Given a string s consisting of n lowercase letters, you have to delete the minimum number of characters from s so that every letter in s appears a unique number of times. We only care about the occurrences of letters that appear at least once in result.

Example 1:

    Input: "eeeeffff"
    Output: 1
Explanation:
We can delete one occurrence of 'e' or one occurrence of 'f'. Then one letter will occur four times and the other three times.

Example 2:

    Input: "aabbffddeaee"
    Output: 6
Explanation:
For example, we can delete all occurrence of 'e' and 'f' and one occurrence of 'd' to obtain the word "aabbda".
Note that both 'e' and 'f' will occur zero times in the new word, but that's fine, since we only care about the letter that appear at least once.

Example 3:

    Input: "llll"
    Output: 0
Explanation:
There is no need to delete any character.

Example 4:

    Input: "example"
    Output: 4


