import unittest
from dataclasses import dataclass
from typing import Callable, List
from leetcode.numberofStepsToReduceToZero.number_of_steps import NumberOfSteps


@dataclass(frozen=True)
class TestSpecs:
    description: str
    num: int
    expected: int


def number_of_steps_test(func: Callable):
    class NumberOfStepsTest(unittest.TestCase):

        def setUp(self) -> None:
            self.func = func
            self.test_cases: List[TestSpecs] = [
                TestSpecs(
                    description="Base Case A",
                    num=14,
                    expected=6,
                ),
                TestSpecs(
                    description="Base Case B",
                    num=8,
                    expected=4,
                ),
                TestSpecs(
                    description="Base Case C",
                    num=123,
                    expected=12,
                ),
                TestSpecs(
                    description="Base Case D",
                    num=9311,
                    expected=21,
                ),
                TestSpecs(
                    description="Base Case E",
                    num=723119,
                    expected=29,
                ),
            ]

        def testImpl(self):
            for tc in self.test_cases:
                with self.subTest(tc.description):
                    actual = func(tc.num)
                    self.assertEqual(tc.expected, actual)

    return NumberOfStepsTest


class TestSolutionA(number_of_steps_test(NumberOfSteps.solution_a)):
    pass


class TestOptimized(number_of_steps_test(NumberOfSteps.optimized)):
    pass


if __name__ == "__main__":
    unittest.main()
