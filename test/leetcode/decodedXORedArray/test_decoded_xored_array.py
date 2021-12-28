import unittest
from dataclasses import dataclass

from typing import Callable, List

from leetcode.decodedXORedArray.decoded_xored_array import DecodedXORedArray


@dataclass(frozen=True)
class TestSpecs:
    description: str
    encoded: List[int]
    first: int
    expected: List[int]


def decoded_xored_array_tests(func: Callable):
    class DecodedXORedArrayTest(unittest.TestCase):

        def setUp(self) -> None:
            self.func = func
            self.test_cases: List[TestSpecs] = [
                TestSpecs(
                    description="Base Case A",
                    encoded=[1, 2, 3],
                    first=1,
                    expected=[1, 0, 2, 1]
                ),
                TestSpecs(
                    description="Base Case B",
                    encoded=[6, 2, 7, 3],
                    first=4,
                    expected=[4, 2, 0, 7, 4]
                ),
            ]

        def testImpl(self):
            for tc in self.test_cases:
                with self.subTest(msg=tc.description):
                    actual = self.func(tc.encoded, tc.first)
                    self.assertListEqual(tc.expected, actual)

    return DecodedXORedArrayTest


class TestOptimized(decoded_xored_array_tests(DecodedXORedArray.optimized)):
    pass


if __name__ == "__main__":
    unittest.main()
