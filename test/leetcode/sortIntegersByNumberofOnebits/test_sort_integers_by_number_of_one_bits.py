from dataclasses import dataclass

from leetcode.sortIntegersByNumberOfOneBits.sort_integers_by_number_of_one_bits import SortIntegersByNumberOfOneDigits
from typing import Callable, List
import unittest


@dataclass(frozen=True)
class TestCase:
    description: str
    arr: List[int]
    expected: List[int]


def sort_integers_by_number_of_one_bits_tests(func: Callable):
    class TestSortIntegersByNumberOfOneDigits(unittest.TestCase):

        def setUp(self) -> None:
            self.func = func
            self.test_cases: List[TestCase] = [
                TestCase(
                    description="Base Case 1",
                    arr=[0, 1, 2, 3, 4, 5, 6, 7, 8],
                    expected=[0, 1, 2, 4, 8, 3, 5, 6, 7]
                ),
                TestCase(
                    description="Base Case 2",
                    arr=[1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1],
                    expected=[1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024]
                )
            ]

        def testImpl(self):
            for tc in self.test_cases:
                with self.subTest(msg=tc.description):
                    actual = self.func(tc.arr)
                    self.assertListEqual(tc.expected, actual)

    return TestSortIntegersByNumberOfOneDigits


class TestSolutionA(sort_integers_by_number_of_one_bits_tests(SortIntegersByNumberOfOneDigits.solution_a)):
    pass


if __name__ == "__main__":
    unittest.main()
