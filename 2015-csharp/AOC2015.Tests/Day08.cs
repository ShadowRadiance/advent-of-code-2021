using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day08Test
    {
        const string dquote = "\"";
        const string bslash = "\\";

        [Theory]
        [InlineData(new string[] { $"{dquote}{dquote}" }, "2")]
        [InlineData(new string[] { $"{dquote}abc{dquote}" }, "2")]
        [InlineData(new string[] { $"{dquote}aaa{bslash}{dquote}aaa{dquote}" }, "3")]
        [InlineData(new string[] { $"{dquote}{bslash}x27{dquote}" }, "5")]
        [InlineData(
            new string[] {
                $"{dquote}{dquote}",
                $"{dquote}abc{dquote}",
                $"{dquote}aaa{bslash}{dquote}aaa{dquote}",
                $"{dquote}{bslash}x27{dquote}",
            },
            "12"
        )]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day08(data).PartA());
        }


        [Theory]
        [InlineData(new string[] { $"{dquote}{dquote}" }, "4")]
        [InlineData(new string[] { $"{dquote}abc{dquote}" }, "4")]
        [InlineData(new string[] { $"{dquote}aaa{bslash}{dquote}aaa{dquote}" }, "6")]
        [InlineData(new string[] { $"{dquote}{bslash}x27{dquote}" }, "5")]
        [InlineData(
            new string[] {
                $"{dquote}{dquote}",
                $"{dquote}abc{dquote}",
                $"{dquote}aaa{bslash}{dquote}aaa{dquote}",
                $"{dquote}{bslash}x27{dquote}",
            },
            "19"
        )]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day08(data).PartB());
        }
    }
}
