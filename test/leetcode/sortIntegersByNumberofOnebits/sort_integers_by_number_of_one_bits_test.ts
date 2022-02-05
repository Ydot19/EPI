import {SortIntegersByNumberOfOneBits} from "../../../leetcode/sortIntegersByNumberOfOneBits/sort_integers_by_number_of_one_bits.ts";
import chai from "https://cdn.skypack.dev/chai@4.3.4?dts";

Deno.test("Test Sort Integers By Number of One Bits", async (t)=>{
    // arrange
    const impl: SortIntegersByNumberOfOneBits = new SortIntegersByNumberOfOneBits();
    const expect = chai.expect;

    type testCase = {
        description: string;
        arr: Array<number>;
        expected: Array<number>;
    }

    const testCases: Array<testCase> = new Array<testCase>(
        {
            description: "Base Case 1",
            arr: [0,1,2,3,4,5,6,7,8],
            expected: [0,1,2,4,8,3,5,6,7]
        },
        {
            description: "Base Case 2",
            arr: [1024,512,256,128,64,32,16,8,4,2,1],
            expected: [1,2,4,8,16,32,64,128,256,512,1024]
        }
    )

    for (const tc of testCases) {
        await t.step(tc.description, ()=>{
            // act
            const actual = impl.SolutionA(tc.arr)
            // assert
            expect(actual).to.eql(tc.expected)
        })
    }
})
