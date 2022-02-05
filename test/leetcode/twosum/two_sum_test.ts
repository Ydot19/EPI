import chai from "https://cdn.skypack.dev/chai@4.3.4?dts";
import { TwoSum } from "../../../leetcode/twosum/two_sum.ts"

Deno.test("Two Sums Implementation", async (t)=>{
    // arrange
    const impl: TwoSum = new TwoSum();
    const expect = chai.expect;
    type testCase = {
        description: string;
        nums: Array<number>;
        target: number;
        expected: Array<number>;
    }
    const testCases: Array<testCase> = new Array<testCase>(
        {
            description: "Base Case A",
            nums: [2,7,11,15],
            target: 9,
            expected: [0,1],
        },
        {
            description: "Base Case B",
            nums: [3,2,4],
            target: 6,
            expected: [1,2],
        },
        {
            description: "Base Case C",
            nums: [3,3],
            target: 6,
            expected: [0,1],
        },
    )

    for (const tc of testCases) {
        await t.step(tc.description, ()=>{
            // act
            const actual = impl.Optimized(tc.nums, tc.target)
            // assert
            expect(actual).to.have.deep.members(tc.expected)
        })
    }
})
