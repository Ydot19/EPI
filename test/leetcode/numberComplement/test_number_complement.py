import unittest
from dataclasses import dataclass
from typing import Callable, List
from leetcode.numberComplement.number_complement import NumberComplement
impl = NumberComplement()


@dataclass
class TestCase:
    description: str
    num: int
    expected: int


def parameterized_number_complement_test(func: Callable):
    class NumberComplementTest(unittest.TestCase):
        def setUp(self) -> None:
            self.func = func
            self.testCases: List[TestCase] = [
                TestCase(
                    description="Case 1",
                    num=5,
                    expected=2,
                ),
                TestCase(
                    description="Case 2",
                    num=2147483646,
                    expected=1,
                ),
                TestCase(
                    description="Case 3",
                    num=319,
                    expected=192,
                )
            ]

        def testImpl(self):

            for tc in self.testCases:
                with self.subTest(tc.description):
                    actual = self.func(tc.num)
                    self.assertEqual(tc.expected, actual)

    return NumberComplementTest


class TestSolutionA(parameterized_number_complement_test(impl.solution_a)):
    pass
