using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day01Test
    {
        [Theory]
        [InlineData(new string[] { "(())" }, "0")]
        [InlineData(new string[] { "()()" }, "0")]
        [InlineData(new string[] { "(((" }, "3")]
        [InlineData(new string[] { "(()(()(" }, "3")]
        [InlineData(new string[] { "))(((((" }, "3")]
        [InlineData(new string[] { "())" }, "-1")]
        [InlineData(new string[] { "))(" }, "-1")]
        [InlineData(new string[] { ")))" }, "-3")]
        [InlineData(new string[] { ")())())" }, "-3")]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day01(data).PartA());
        }

        [Theory]
        [InlineData(new string[] { ")" }, "1")]
        [InlineData(new string[] { "()())" }, "5")]
        [InlineData(new string[] { "(())" }, "X")]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day01(data).PartB());
        }
    }
}
