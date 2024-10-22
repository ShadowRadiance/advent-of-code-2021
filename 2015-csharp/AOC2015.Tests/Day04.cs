using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day04Test
    {
        [Theory]
        [InlineData(new string[] { "abcdef" }, "609043")]
        [InlineData(new string[] { "pqrstuv" }, "1048970")]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day04(data).PartA());
        }


        [Theory]
        [InlineData(new string[] { "abcdef" }, "6742839")]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day04(data).PartB());
        }
    }
}
