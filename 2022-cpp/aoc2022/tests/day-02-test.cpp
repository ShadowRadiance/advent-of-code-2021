#include <gtest/gtest.h>

#include <vector>
#include <string>

#include <day-02.h>

TEST(Day02, Part1_TDD_AY)
{
  std::vector<std::string> data{
	"A Y",
  };

  EXPECT_EQ("8", day_02::answer_a(data));
}

TEST(Day02, Part1_TDD_BX)
{
  std::vector<std::string> data{
	"B X",
  };

  EXPECT_EQ("1", day_02::answer_a(data));
}

TEST(Day02, Part1_TDD_CZ)
{
  std::vector<std::string> data{
	"C Z",
  };

  EXPECT_EQ("6", day_02::answer_a(data));
}

TEST(Day02, Part1Example)
{
	std::vector<std::string> data{
	  "A Y",
	  "B X",
	  "C Z",
	};

	EXPECT_EQ("15", day_02::answer_a(data));
}

TEST(Day02, Part2Example)
{
	std::vector<std::string> data{
	};

	EXPECT_EQ("PENDING", day_02::answer_b(data));
}
