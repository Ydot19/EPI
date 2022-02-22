import unittest
from dataclasses import dataclass
from typing import Callable, List

from leetcode.primeNumberOfSetsInBinaryRepresentation.prime_number_of_sets_in_binary_representation import \
    PrimeNumberOfSetsInBinaryRepresentationInterface, Impl


@dataclass
class TestCase:
    description: str
    left: int
    right: int
    expected: int


Impl: PrimeNumberOfSetsInBinaryRepresentationInterface = Impl()


def parameterized_prime_number_of_sets_in_binary_representation_test(func: Callable):
    class PrimeNumberOfSetsInBinaryRepresentationTest(unittest.TestCase):
        def setUp(self) -> None:
            self.func = func
            self.testCases: List[TestCase] = [
                TestCase(
                    description="Case 1",
                    left=6,
                    right=10,
                    expected=4,
                ),
                TestCase(
                    description="Case 2",
                    left=10,
                    right=15,
                    expected=5,
                ),
                TestCase(
                    description="Case 3",
                    left=14,
                    right=97,
                    expected=54,
                ),
                TestCase(
                    description="Case 4",
                    left=131,
                    right=926,
                    expected=416,
                ),
            ]

        def testImpl(self):
            for tc in self.testCases:
                with self.subTest(tc.description):
                    actual = self.func(tc.left, tc.right)
                    self.assertEqual(tc.expected, actual)

    return PrimeNumberOfSetsInBinaryRepresentationTest


class TestOptimized(parameterized_prime_number_of_sets_in_binary_representation_test(Impl.optimized)):
    pass


class TestSolutionA(parameterized_prime_number_of_sets_in_binary_representation_test(Impl.solution_a)):
    pass
