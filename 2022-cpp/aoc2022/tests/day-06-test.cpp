#include <gtest/gtest.h>

#include <days.h>

TEST(Day06, Part1Example1)
{
    std::vector<std::string> data{
      "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
    };

    EXPECT_EQ("7", day_06::answer_a(data));
}

TEST(Day06, Part1Example2)
{
    std::vector<std::string> data{
      "bvwbjplbgvbhsrlpgdmjqwftvncz",
    };

    EXPECT_EQ("5", day_06::answer_a(data));
}

TEST(Day06, Part1Example3)
{
    std::vector<std::string> data{
      "nppdvjthqldpwncqszvftbrmjlhg",
    };

    EXPECT_EQ("6", day_06::answer_a(data));
}

TEST(Day06, Part1Example4)
{
    std::vector<std::string> data{
      "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
    };

    EXPECT_EQ("10", day_06::answer_a(data));
}

TEST(Day06, Part1Example5)
{
    std::vector<std::string> data{
      "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
    };

    EXPECT_EQ("11", day_06::answer_a(data));
}

TEST(Day06, Part2Example1)
{
    std::vector<std::string> data{
      "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
    };

    EXPECT_EQ("19", day_06::answer_b(data));
}

TEST(Day06, Part2Example2)
{
    std::vector<std::string> data{
      "bvwbjplbgvbhsrlpgdmjqwftvncz",
    };

    EXPECT_EQ("23", day_06::answer_b(data));
}

TEST(Day06, Part2Example3)
{
    std::vector<std::string> data{
      "nppdvjthqldpwncqszvftbrmjlhg",
    };

    EXPECT_EQ("23", day_06::answer_b(data));
}

TEST(Day06, Part2Example4)
{
    std::vector<std::string> data{
      "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
    };

    EXPECT_EQ("29", day_06::answer_b(data));
}

TEST(Day06, Part2Example5)
{
    std::vector<std::string> data{
      "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
    };

    EXPECT_EQ("26", day_06::answer_b(data));
}

