import {assertEquals} from "https://deno.land/std@0.119.0/testing/asserts.ts";
import {NumberComplement} from "../../../leetcode/numberComplement/number_complement.ts"

const Impl: NumberComplement = new NumberComplement()

type testCase = {
    description: string
    num: number
    expected: number
}

const testCases: Array<testCase> = [
    {
        description:"Case 1",
        num:5,
        expected: 2,
    },
    {
        description: "Case 2",
        num:         2147483646,
        expected:    1,
    },
    {
        description: "Case 3",
        num:         319,
        expected:    192,
    }
]

Deno.test("Test Number Complement Implementation", async(t)=>{
    const testFuncs = [
        Impl.solutionA,
        Impl.solutionB,
        Impl.optimized
    ]

    for (const func of testFuncs) {
        for (const tc of testCases){
            const description = `${func.name}: ${tc.description}`
            await t.step(description, ()=>{
                // act
                const actual = func(tc.num)
                // assert
                assertEquals(tc.expected, actual)
            })
        }
    }
})
