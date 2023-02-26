#include <gtest/gtest.h>

#include <days.h>

const std::vector<std::string> data{
    "1",
    "2",
    "-3",
    "3",
    "-2",
    "0",
    "4",
};

TEST(Day20, Part1Example)
{
    EXPECT_EQ("3", day_20::answer_a(data));
}

TEST(Day20, Part2Example)
{
    EXPECT_EQ("1623178306", day_20::answer_b(data));
}