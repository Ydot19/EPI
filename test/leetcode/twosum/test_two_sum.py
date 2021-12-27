from dataclasses import dataclass

from leetcode.twosum.two_sum import TwoSum
from typing import Callable, List
import unittest


@dataclass(frozen=True)
class TestSpecs:
    description: str
    nums: List[int]
    target: int
    expected: List[int]


def two_sum_tests(func: Callable):
    class TwoSumTestCases(unittest.TestCase):

        def setUp(self) -> None:
            self.func = func
            self.test_cases: List[TestSpecs] = [
                TestSpecs(
                    description="Base Case A",
                    nums=[2, 7, 11, 15],
                    target=9,
                    expected=[0, 1]
                ),
                TestSpecs(
                    description="Base Case B",
                    nums=[3, 2, 4],
                    target=6,
                    expected=[1, 2]
                ),
                TestSpecs(
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

    return TwoSumTestCases


class TestOptimized(two_sum_tests(TwoSum().a)):
    pass


if __name__ == "__main__":
    unittest.main()
