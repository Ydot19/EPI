import {DecodedXORedArray} from "../../../leetcode/decodedXORedArray/decoded_xored_array.ts";
import { assertEquals } from "https://deno.land/std@0.119.0/testing/asserts.ts";


Deno.test("Decoded XORed Array Implementation", async (t)=>{
    // arrange
    const impl: DecodedXORedArray = new DecodedXORedArray()
    type testSpec = {
        description: string;
        encoded: Array<number>;
        first: number;
        expected: Array<number>;
    }

    const testCases: Array<testSpec> = new Array<testSpec>(
        {
            description: "Base Case A",
            encoded: [1, 2, 3],
            first: 1,
            expected: [1, 0, 2, 1]
        },
        {
            description: "Base Case B",
            encoded: [6, 2, 7, 3],
            first: 4,
            expected: [4, 2, 0, 7, 4]
        }
    )

    for (const tc of testCases) {
        await t.step(tc.description, () => {
            // act
            const actual = impl.Optimized(tc.encoded, tc.first)
            // assert
            assertEquals(actual, tc.expected)
        })
    }
})
