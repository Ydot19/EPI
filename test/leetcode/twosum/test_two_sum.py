from dataclasses import dataclass

from leetcode.twosum.two_sum import TwoSum
from typing import Callable, List
import unittest


@dataclass(frozen=True)
class TestCase:
    description: str
    nums: List[int]
    target: int
    expected: List[int]


def two_sum_tests(func: Callable):
    class TwoSumTest(unittest.TestCase):

        def setUp(self) -> None:
            self.func = func
            self.test_cases: List[TestCase] = [
                TestCase(
                    description="Base Case A",
                    nums=[2, 7, 11, 15],
                    target=9,
                    expected=[0, 1]
                ),
                TestCase(
                    description="Base Case B",
                    nums=[3, 2, 4],
                    target=6,
                    expected=[1, 2]
                ),
                TestCase(
                    description="Base Case C",
                    nums=[3, 3],
                    target=6,
                    expected=[0, 1]
                )
            ]

        def testImpl(self):
            for tc in self.test_cases:
                with self.subTest(msg=tc.description):
                    actual = self.func(tc.nums, tc.target)
                    self.assertCountEqual(tc.expected, actual)

    return TwoSumTest


class TestOptimized(two_sum_tests(TwoSum.optimized)):
    pass


if __name__ == "__main__":
    unittest.main()
