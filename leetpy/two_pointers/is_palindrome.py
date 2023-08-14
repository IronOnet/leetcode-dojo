"""
A phrase is a paindrome if, after converting all uppercase letters into lowercase 
letters and removing all non-alphanumeric characters, it reads the same forward 
and backward. Alphanumeric characters include letters and numbers

Example 1: 

Input: s = "A man, a plan, a canal: Panama" 
Output: true 
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2: 

Input: s = "race a car" 
Output: false 
Explanation: "raceacar" is not valid palindrome"""

def is_valid_palindrome(s : str) -> bool: 
    return s == s[::-1]

def is_valid_palindrome_2(s: str) -> bool: 
    left, right = 0, len(s) - 1 
    for char in s: 
        if s[left] != s[right]: 
            return False 
        left += 1 
        right -= 1 
    return True 