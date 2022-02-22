import {
    PrimeNumberOfSetsInBinaryRepresentation
} from "../../../leetcode/primeNumberOfSetsInBinaryRepresentation/prime_number_of_sets_in_binary_representation.ts";
import {assertEquals} from "https://deno.land/std@0.119.0/testing/asserts.ts";


Deno.test("Test Prime Number Of Sets In Binary Representation Implementation", async(t)=>{
    // arrange
    const Impl = new PrimeNumberOfSetsInBinaryRepresentation()
    type testCase = {
        description: string
        left:        number
        right:       number
        expected:    number
    }

    const testCases: Array<testCase> = [
        {
            description: "Case 1",
            left:        6,
            right:       10,
            expected:    4,
        },
        {
            description: "Case 2",
            left:        10,
            right:       15,
            expected:    5,
        },
        {
            description: "Case 3",
            left:        14,
            right:       97,
            expected:    54,
        },
        {
            description: "Case 4",
            left:        131,
            right:       926,
            expected:    416,
        },
    ]

    const testFuncs = [
        Impl.solutionA,
        Impl.optimized
    ]

    for (const func of testFuncs) {
        for (const tc of testCases){
            const description = `${func.name}: ${tc.description}`
            await t.step(description, ()=>{
                const actual = func(tc.left, tc.right)
                assertEquals(actual, tc.expected)
            })
        }
    }
})
