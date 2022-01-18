import unittest
from dataclasses import dataclass
from typing import Callable, List

from leetcode.countingBits.counting_bits import CountingBits


@dataclass
class TestCase:
    description: str
    n: int
    expected: List[int]


def counting_bits_test(func: Callable):
    class CountingBitsTest(unittest.TestCase):
        def setUp(self) -> None:
            self.func = func
            self.testCases: List[TestCase] = [
                TestCase(
                    description="Case A",
                    n=2,
                    expected=[0, 1, 1]
                ),
                TestCase(
                    description="Case B",
                    n=5,
                    expected=[0, 1, 1, 2, 1, 2]
                ),
                TestCase(
                    description="Case C",
                    n=98,
                    expected=[0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4,
                              4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5,
                              4, 5, 5, 6, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4,
                              4, 5, 4, 5, 5, 6, 2, 3, 3]
                ),
                TestCase(
                    description="Case D",
                    n=211,
                    expected=[0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4,
                              4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5,
                              4, 5, 5, 6, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4,
                              4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6,
                              4, 5, 5, 6, 5, 6, 6, 7, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4,
                              4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5,
                              4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4,
                              4, 5]
                )
            ]

        def testImpl(self):
            for tc in self.testCases:
                with self.subTest(tc.description):
                    actual = self.func(tc.n)
                    self.assertListEqual(tc.expected, actual)

    return CountingBitsTest


class TestCountingBits(counting_bits_test(CountingBits.solution_a)):
    pass


if __name__ == '__main__':
    unittest.main()
