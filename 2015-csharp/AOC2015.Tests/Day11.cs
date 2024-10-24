using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day11Test
    {
        [Theory]
        [InlineData("hijklmmn", "hijklmmp")]
        [InlineData("abbceffz", "abbcefga")]
        [InlineData("abbcefzz", "abbcegaa")]
        public void Increment_Test(string input, string expectation)
        {
            Assert.Equal(expectation, Day11.Increment(input));
        }

        [Theory]
        [InlineData("hijklmmn", false)] // contains i, l, or o
        [InlineData("abbceffg", false)] // does not have increasing sequence
        [InlineData("abbcegjk", false)] // does not have increasing sequence
        [InlineData("ghjaabcd", false)] // does not have two overlapping pairs
        [InlineData("abcdffaa", true)]
        [InlineData("ghjaabcc", true)]
        public void PasswordValid_Test(string input, bool expectation)
        {
            Assert.Equal(expectation, Day11.PasswordValid(input));
        }


        [Theory]
        [InlineData(new string[] { "abcdefgh" }, "abcdffaa")]
        [InlineData(new string[] { "ghijklmn" }, "ghjaabcc")]
        [InlineData(new string[] { "ghjaabcc" }, "ghjbbcdd")]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day11(data).PartA());
        }

        [Theory]
        [InlineData(new string[] { "abcdefgh" }, "abcdffbb")]
        [InlineData(new string[] { "ghijklmn" }, "ghjbbcdd")]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day11(data).PartB());
        }
    }
}
