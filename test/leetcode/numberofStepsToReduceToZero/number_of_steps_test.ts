import {NumberOfSteps} from "../../../leetcode/numberofStepsToReduceToZero/number_of_steps.ts";
import { assertEquals } from "https://deno.land/std@0.119.0/testing/asserts.ts";

// arrange
type testSpec = {
    description: string;
    num: number;
    expected: number;
}

const testCases: Array<testSpec> = [
    {
        description: "Base Case A",
        num:         14,
        expected:    6,
    },
    {
        description: "Base Case B",
        num:         8,
        expected:    4,
    },
    {
        description: "Base Case C",
        num:         123,
        expected:    12,
    },
    {
        description: "Base Case D",
        num:         9311,
        expected:    21,
    },
    {
        description: "Base Case E",
        num:         723119,
        expected:    29,
    }
]

const impl: NumberOfSteps = new NumberOfSteps()


Deno.test("Test NumberOfSteps SolutionA", async (t)=>{
    for (const tc of testCases){
        await t.step(tc.description, ()=>{
            // act
            const actual = impl.SolutionA(tc.num)
            // assert
            assertEquals(tc.expected, actual)
        })
    }
})


Deno.test("Test NumberOfSteps Optimized", async (t)=>{
    for (const tc of testCases){
        await t.step(tc.description, ()=>{
            // act
            const actual = impl.Optimized(tc.num)
            // assert
            assertEquals(tc.expected, actual)
        })
    }
})
