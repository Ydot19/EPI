/**
 * You are given a string allowed consisting of distinct characters and an array of strings words. A string is consistent if all characters in the string appear in the string allowed.
 *
 * Return the number of consistent strings in the array words.
 *
 *
 *
 * Example 1:
 *
 * Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
 * Output: 2
 * Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.
 * Example 2:
 *
 * Input: allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
 * Output: 7
 * Explanation: All strings are consistent.
 * Example 3:
 *
 * Input: allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
 * Output: 4
 * Explanation: Strings "cc", "acd", "ac", and "d" are consistent.
 */

export class CountNumberOfConsistentLetters {
    SolutionA(allowed: string, words: Array<string>): number {
        const allowedSet = new Set<string>();
        [...allowed].forEach(c => allowedSet.add(c))
        let count = 0
        for (const word of words) {
            let add = 1
            for (const char of word) {
                if (!allowedSet.has(char)) {
                    add = 0
                    break
                }
            }
            count += add
        }
        return count
    }

    SolutionB(allowed: string, words: Array<string>): number {
        const allowedSet = new Set<string>();
        [...allowed].forEach(c => allowedSet.add(c))
        let count = 0
        for (const word of words) {
            const diff = new Set(
                [...word].filter(char => !allowedSet.has(char))
            )
            if (diff.size == 0) {
                count += 1
            }
        }
        return count
    }

    Optimized(allowed: string, words: Array<string>): number {
        let allowedCharacters = 0
        for (const char of allowed){
            const val = char.charCodeAt(0) - 'a'.charCodeAt(0)
            allowedCharacters |= 1 << val
        }
        let count = 0
        for (const word of words){
            let add = 1
            for (const char of word){
                const val = char.charCodeAt(0) - 'a'.charCodeAt(0)
                if ((allowedCharacters & (1 << val)) == 0){
                    add = 0
                    break
                }
            }
            count += add
        }
        return count
    }
}
