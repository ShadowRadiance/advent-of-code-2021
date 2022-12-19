#include <vector>
#include <string>
#include <array>
#include <algorithm>
#include <iterator>
#include <numeric>
#include <set>

#include <day-03.h>

namespace day_03
{
  using Backpack = std::array<std::string, 2>;
  using Backpacks = std::vector<Backpack>;

  Backpack backpack_from_line(const std::string& line)
  {
    size_t halfLength = line.length() / 2;
    return { line.substr(0, halfLength), line.substr(halfLength) };
  }

  std::string determineIntersections(const Backpack& backpack) {
    std::string splitUpItemTypes;
    std::string first = backpack[0];
    std::string second = backpack[1];
    std::sort(first.begin(), first.end());
    std::sort(second.begin(), second.end());

    std::set_intersection(first.begin(), first.end(), second.begin(), second.end(), std::back_inserter(splitUpItemTypes));
    splitUpItemTypes.erase(
      std::unique(splitUpItemTypes.begin(), splitUpItemTypes.end())
    );
    return splitUpItemTypes;
  };

  int scoreChar(char c) {
    // - Lowercase item types a through z have priorities 1 through 26.
    // - Uppercase item types A through Z have priorities 27 through 52.

    // ASCII a ==> INT 97, so subtract 96
    // ASCII A ==> INT 65, so subtract 38

    if (c >= 'a' && c <= 'z') return c - 96;
    if (c >= 'A' && c <= 'Z') return c - 38;

    return 0;
  }

  std::string answer_a(const std::vector<std::string>& input_data)
  {
    Backpacks backpacks;
    std::transform(input_data.begin(), input_data.end(), std::back_inserter(backpacks), backpack_from_line);

    std::vector<std::string> splitUpItemTypesPerBackpack;
    std::transform(backpacks.begin(), backpacks.end(), std::back_inserter(splitUpItemTypesPerBackpack), determineIntersections);

    // the text seems to presume there will only be one overlapping item type per pair, but lets be more careful

    int score = std::transform_reduce(
      splitUpItemTypesPerBackpack.begin(), splitUpItemTypesPerBackpack.end(),
      0,
      std::plus<>{},
      [](const std::string& s) {
        return std::transform_reduce(
          s.begin(), s.end(),
          0,
          std::plus<>{},
          [](const char& c) {
            return scoreChar(c);
          }
        );
      }
    );

    return std::to_string(score);
  }

  std::string answer_b(const std::vector<std::string>& input_data)
  {
    return "PENDING";
  }
}
