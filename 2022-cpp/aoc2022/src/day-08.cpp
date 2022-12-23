#include <days.h>

#include <optional>

namespace day_08
{
  using std::string;
  using std::vector;
  using std::optional;

  using NilableBool = optional<bool>;
  struct Tree
  {
      int height = 0;
      NilableBool is_visible;
      int max_north = 0;
      int max_east = 0;
      int max_south = 0;
      int max_west = 0;
  };

  using strings = vector<string>;
  using Forest = vector<Tree>;

  Forest generateForest(const strings& input_data)
  {
    Forest forest;
    for (size_t y{0}; y < input_data.size(); y++)
    {
      const string& row = input_data[y];
      for (size_t x{0}; x < row.length(); x++)
      {
        char cHeight = row[x];
        int nHeight = cHeight - '0';
        Tree tree {
          nHeight,
          y==0 || x==0 || y==input_data.size()-1 || x==row.size()-1
        };
        forest.push_back(tree);
      }
    }
    return forest;
  }

  string answer_a(const strings& input_data)
  {
    // read input data into grid (set height, set is visible if on edge)
    Forest forest = generateForest(input_data);

    // each grid location stores: north_max, east_max, etc. and is_visible
    // pass 1 (forward - or "moving W->E and N->S"
    //   look north record max in that direction (if current > max, is visible)
    //   look west record max in that direction (if current > max, is visible)
    // pass 2 (backward - or "moving W<-E and N<-S"
    //   look south record max in that direction (if current > max, is visible)
    //   look east record max in that direction (if current > max, is visible)
    // summarize: count is visible

    return "PENDING";
  }

  string answer_b(const vector<string>& input_data)
  {
    return "PENDING";
  }
}
