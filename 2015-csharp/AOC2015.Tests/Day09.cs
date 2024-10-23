using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day09Test
    {
        [Theory]
        [InlineData(
            new string[] {
                "London to Dublin = 464",
                "London to Belfast = 518",
                "Dublin to Belfast = 141"
            },
            "605"
        )]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day09(data).PartA());
        }


        [Theory]
        [InlineData(
            new string[] {
                "London to Dublin = 464",
                "London to Belfast = 518",
                "Dublin to Belfast = 141"
            },
            "982"
        )]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day09(data).PartB());
        }
    }
}
