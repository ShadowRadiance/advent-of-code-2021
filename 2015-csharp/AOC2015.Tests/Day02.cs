using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day02Test
    {
        // A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
        // A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
        [Theory]
        [InlineData(new string[] { "2x3x4" }, "58")]
        [InlineData(new string[] { "1x1x10" }, "43")]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day02(data).PartA());
        }


        // A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
        // A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.
        [Theory]
        [InlineData(new string[] { "2x3x4" }, "34")]
        [InlineData(new string[] { "1x1x10" }, "14")]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day02(data).PartB());
        }
    }
}
