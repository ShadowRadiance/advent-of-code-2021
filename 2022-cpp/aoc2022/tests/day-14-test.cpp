#include <gtest/gtest.h>

#include <days.h>
#include <algorithm>

const std::vector<std::string> data{
    "498,4 -> 498,6 -> 496,6",
    "503,4 -> 502,4 -> 502,9 -> 494,9"
};

TEST(Day14, Part1Example)
{
    EXPECT_EQ("24", day_14::answer_a(data));
}

TEST(Day14, Part2Example)
{
    EXPECT_EQ("93", day_14::answer_b(data));
}