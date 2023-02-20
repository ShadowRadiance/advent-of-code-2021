#include <gtest/gtest.h>

#include <days.h>

std::vector<std::string> const data{
	"2,2,2",
	"1,2,2",
	"3,2,2",
	"2,1,2",
	"2,3,2",
	"2,2,1",
	"2,2,3",
	"2,2,4",
	"2,2,6",
	"1,2,5",
	"3,2,5",
	"2,1,5",
	"2,3,5",
};

TEST(Day18, Part1Example)
{

  EXPECT_EQ("64", day_18::answer_a(data));
}

TEST(Day18, Part2Example)
{
  EXPECT_EQ("PENDING", day_18::answer_b(data));
}