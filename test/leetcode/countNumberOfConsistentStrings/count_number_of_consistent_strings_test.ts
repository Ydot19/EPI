import {CountNumberOfConsistentLetters} from "../../../leetcode/countNumberOfConsistentStrings/count_number_of_consistent_strings.ts";
import { assertEquals } from "https://deno.land/std@0.119.0/testing/asserts.ts";

// arrange
const impl: CountNumberOfConsistentLetters = new CountNumberOfConsistentLetters()
type testSpec = {
    description: string;
    allowed: string;
    words: Array<string>;
    expected: number;
}

const testCases: Array<testSpec> = [
    {
        description: "Base Case 1",
        allowed:     "ab",
        words:       ["ad", "bd", "aaab", "baa", "badab"],
        expected:    2,
    },
    {
        description: "Base Case 2",
        allowed:     "abc",
        words:       ["a", "b", "c", "ab", "ac", "bc", "abc", "abcd"],
        expected:    7,
    },
    {
        description: "Base Case 3",
        allowed:     "cad",
        words:       ["cc", "acd", "b", "ba", "bac", "bad", "ac", "d"],
        expected:    4,
    }
]

Deno.test("Test Count Number of Consistent Letters Solution A", async (t)=>{
    for (const tc of testCases){
        await t.step(tc.description, ()=>{
            // act
            const actual = impl.SolutionA(tc.allowed, tc.words)
            // assert
            assertEquals(tc.expected, actual)
        })
    }
})

Deno.test("Test Count Number of Consistent Letters Solution B", async (t)=>{
    for (const tc of testCases){
        await t.step(tc.description, ()=>{
            // act
            const actual = impl.SolutionB(tc.allowed, tc.words)
            // assert
            assertEquals(tc.expected, actual)
        })
    }
})

Deno.test("Test Count Number of Consistent Letters Solution C", async (t)=>{
    for (const tc of testCases){
        await t.step(tc.description, ()=>{
            // act
            const actual = impl.Optimized(tc.allowed, tc.words)
            // assert
            assertEquals(tc.expected, actual)
        })
    }
})
