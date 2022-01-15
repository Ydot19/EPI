import unittest
from dataclasses import dataclass
from typing import Callable, List

from leetcode.hammingDistance.hamming_distance import HammingDistance


@dataclass
class TestCase:
    description: str
    x: int
    y: int
    expected: int


def hamming_distance_test(func: Callable):
    class HammingDistanceTest(unittest.TestCase):
        def setUp(self) -> None:
            self.func = func
            self.testCases: List[TestCase] = [
                TestCase(
                    description="Base Case 1",
                    x=1,
                    y=4,
                    expected=2,
                ),
                TestCase(
                    description="Base Case 2",
                    x=10,
                    y=99,
                    expected=4,
                ),
                TestCase(
                    description="Base Case 3",
                    x=1093,
                    y=93443112,
                    expected=13,
                ),
                TestCase(
                    description="Base Case 4",
                    x=749199,
                    y=213,
                    expected=12,
                ),
                TestCase(
                    description="Base Case 5",
                    x=5639,
                    y=1436241892,
                    expected=18,
                ),
            ]

        def testImpl(self):
            for tc in self.testCases:
                with self.subTest(msg=tc.description):
                    actual = self.func(tc.x, tc.y)
                    self.assertEqual(tc.expected, actual)

    return HammingDistanceTest


class TestOptimized(hamming_distance_test(HammingDistance.optimized)):
    pass


if __name__ == "__main__":
    unittest.main()
