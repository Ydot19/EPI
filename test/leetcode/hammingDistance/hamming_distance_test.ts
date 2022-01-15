import {HammingDistance} from "../../../leetcode/hammingDistance/hamming_distance.ts";
import {assertEquals} from "https://deno.land/std@0.119.0/testing/asserts.ts";

Deno.test("Test Hamming Distance Implementation", async (t) =>{
    // arrange
    const impl = new HammingDistance()
    type testCase = {
        description: string
        x: number
        y: number
        expected: number
    }

    const testCases: Array<testCase> = [
        {
            description: "Base Case 1",
            x:           1,
            y:           4,
            expected:    2,
        },
        {
            description: "Base Case 2",
            x:           10,
            y:           99,
            expected:    4,
        },
        {
            description: "Base Case 3",
            x:           1093,
            y:           93443112,
            expected:    13,
        },
        {
            description: "Base Case 4",
            x:           749199,
            y:           213,
            expected:    12,
        },
        {
            description: "Base Case 5",
            x:           5639,
            y:           1436241892,
            expected:    18,
        },
    ]

    for (const tc of testCases){
        await t.step(tc.description, ()=>{
            const actual = impl.Optimized(tc.x, tc.y)
            assertEquals(actual, tc.expected)
        })
    }
})
