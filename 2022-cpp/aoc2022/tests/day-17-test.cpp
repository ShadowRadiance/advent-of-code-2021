#include <gtest/gtest.h>

#include <days.h>

std::vector<std::string> data{
	">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
};

TEST(Day17, Part1Example)
{
	EXPECT_EQ("3068", day_17::answer_a(data));
}

TEST(Day17, Part2Example)
{
	EXPECT_EQ("1514285714288", day_17::answer_b(data));
}