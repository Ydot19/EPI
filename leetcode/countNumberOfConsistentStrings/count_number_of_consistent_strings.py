"""
You are given a string allowed consisting of distinct characters and an array of strings words. A string is consistent if all characters in the string appear in the string allowed.

Return the number of consistent strings in the array words.



Example 1:

Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
Output: 2
Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.
Example 2:

Input: allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
Output: 7
Explanation: All strings are consistent.
Example 3:

Input: allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
Output: 4
Explanation: Strings "cc", "acd", "ac", and "d" are consistent.
"""

from typing import List


class CountNumberOfConsistentLetters:
    @staticmethod
    def solution_a(allowed: str, words: List[str]) -> int:
        allowed_set = set(allowed)
        count = 0
        for word in words:
            unq = set(word)
            count += 1 if unq.issubset(allowed_set) else 0

        return count

    @staticmethod
    def solution_b(allowed: str, words: List[str]) -> int:
        allowed_set = set(allowed)
        count = 0
        for word in words:
            add = 1
            for letter in word:
                if letter not in allowed_set:
                    add = 0
                    break

            count += add

        return count

    @staticmethod
    def optimized(allowed: str, words: List[str]) -> int:
        allowed_characters = 0
        for char in allowed:
            val = ord(char) - ord('a')  # get the character value
            # append to bitwise expression
            allowed_characters |= 1 << val   # shift left based on character length

        count = 0
        for word in words:
            add = 1
            for char in word:
                if (allowed_characters & (1 << (ord(char) - ord('a')))) == 0:
                    add = 0
                    break

            count += add
        return count
