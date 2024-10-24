using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day12Test
    {
        [Theory]
        [InlineData(new string[] { $$"""[1,2,3]""" }, "6")]
        [InlineData(new string[] { $$"""{"a":2,"b":4}""" }, "6")]
        [InlineData(new string[] { $$"""[[[3]]]""" }, "3")]
        [InlineData(new string[] { $$"""{"a":{"b":4},"c":-1}""" }, "3")]
        [InlineData(new string[] { $$"""{"a":[-1,1]}""" }, "0")]
        [InlineData(new string[] { $$"""[-1,{"a":1}]""" }, "0")]
        [InlineData(new string[] { $$"""[]""" }, "0")]
        [InlineData(new string[] { $$"""{}""" }, "0")]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day12(data).PartA());
        }


        [Theory]
        [InlineData(new string[] { $$"""[1,2,3]""" }, "6")]
        [InlineData(new string[] { $$"""{"a":2,"b":"red"}""" }, "0")]
        [InlineData(new string[] { $$"""[[[3,"red",4]]]""" }, "7")]
        [InlineData(new string[] { $$"""{"a":{"q":"red","b":4},"c":-1}""" }, "-1")]
        [InlineData(new string[] { $$"""{"a":{"b":4},"c":"red"}""" }, "0")]
        [InlineData(new string[] { $$"""{"a":[-1,1]}""" }, "0")]
        [InlineData(new string[] { $$"""[-1,{"a":1}]""" }, "0")]
        [InlineData(new string[] { $$"""[]""" }, "0")]
        [InlineData(new string[] { $$"""{}""" }, "0")]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day12(data).PartB());
        }
    }
}
