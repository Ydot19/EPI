import unittest
from dataclasses import dataclass
from typing import Callable, List

from leetcode.countNumberOfConsistentStrings.count_number_of_consistent_strings import CountNumberOfConsistentLetters


@dataclass()
class TestSpecs:
    description: str
    allowed: str
    words: List[str]
    expected: int


def count_number_of_consistent_strings_tests(func: Callable):
    class CountNumberOfConsistentStringTest(unittest.TestCase):
        def setUp(self) -> None:
            self.func = func
            self.test_cases: List[TestSpecs] = [
                TestSpecs(
                    description="Base Case 1",
                    allowed="ab",
                    words=["ad", "bd", "aaab", "baa", "badab"],
                    expected=2
                ),
                TestSpecs(
                    description="Base Case 2",
                    allowed="abc",
                    words=["a", "b", "c", "ab", "ac", "bc", "abc", "abcd"],
                    expected=7,
                ),
                TestSpecs(
                    description="Base Case 3",
                    allowed="cad",
                    words=["cc", "acd", "b", "ba", "bac", "bad", "ac", "d"],
                    expected=4,
                )
            ]

        def testImpl(self):
            for tc in self.test_cases:
                with self.subTest(msg=tc.description):
                    actual = self.func(tc.allowed, tc.words)
                    self.assertEqual(tc.expected, actual)

    return CountNumberOfConsistentStringTest


class TestSolutionA(count_number_of_consistent_strings_tests(CountNumberOfConsistentLetters.solution_a)):
    pass


class TestSolutionB(count_number_of_consistent_strings_tests(CountNumberOfConsistentLetters.solution_b)):
    pass


class TestOptimized(count_number_of_consistent_strings_tests(CountNumberOfConsistentLetters.optimized)):
    pass


if __name__ == "__main__":
    unittest.main()
