#include <gtest/gtest.h>

#include <days.h>

std::vector<std::string> data = {
    "[1,1,3,1,1]",
    "[1,1,5,1,1]",
    "",
    "[[1],[2,3,4]]",
    "[[1],4]",
    "",
    "[9]",
    "[[8,7,6]]",
    "",
    "[[4,4],4,4]",
    "[[4,4],4,4,4]",
    "",
    "[7,7,7,7]",
    "[7,7,7]",
    "",
    "[]",
    "[3]",
    "",
    "[[[]]]",
    "[[]]",
    "",
    "[1,[2,[3,[4,[5,6,7]]]],8,9]",
    "[1,[2,[3,[4,[5,6,0]]]],8,9]",
};

TEST(Day13, Part1Example)
{
    EXPECT_EQ("13", day_13::answer_a(data));
}

TEST(Day13, Part2Example)
{    
    EXPECT_EQ("140", day_13::answer_b(data));
}