// Write a function that takes an unsigned integer and return the number of '1' bits it has (also known as the Hamming weight).

// Input: 00000000000000000000000000001011
// Output: 3
// Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.

func hammingWeight(num uint32) int {
    result := 0
    for num > 0 {
        if num & 1 != 0 {
            result += 1
        }
        num >>= 1
    }
    return result
}
