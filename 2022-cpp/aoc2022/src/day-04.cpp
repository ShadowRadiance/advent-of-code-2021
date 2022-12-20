#include <days.h>

#include <ranges>
#include <algorithm>
#include <tuple>
#include <array>
#include <sstream>

namespace day_04
{
  using std::string;
  using std::vector;
  using std::array;
  using std::ranges::iota_view;
  using std::stringstream;
  using std::views::iota;
  using std::to_string;
  using std::ranges::includes;
  using std::views::filter;
  using std::ranges::set_intersection;

  using strings = vector<string>;
  using Range = iota_view<int, int>;
  using RangePair = array<Range, 2>;
  using RangePairs = vector<RangePair>;


  RangePair buildRangePair(const string& s)
  {
    stringstream ss{ s };
    // ss => 22-33,44-55
    int start1, end1, start2, end2;
    char dummy;
    ss >> start1 >> dummy >> end1 >> dummy >> start2 >> dummy >> end2;
    
    return {
      // +1 because iota requires "one-past-the-end" like iterators
      iota(start1, end1+1),
      iota(start2, end2+1),
    };
  }

  RangePairs buildRangePairs(const strings& list)
  {
    RangePairs rangePairs;
    transform(list.begin(), list.end(), back_inserter(rangePairs), buildRangePair);
    return rangePairs;
  }

  bool completelyOverlaps(RangePair rangePair)
  {
    return includes(rangePair[0], rangePair[1]) || includes(rangePair[1], rangePair[0]);
  }

  bool overlaps(RangePair rangePair)
  {
    // is the intersection empty? - then overlaps == false - else true
    vector<int> intersection;
    set_intersection(
      rangePair[0],
      rangePair[1],
      back_inserter(intersection)
    );
    return !intersection.empty();
  }

  RangePairs findCompletelyOverlapping(RangePairs rangePairs)
  {
    auto filteredPairs = rangePairs | filter(completelyOverlaps);
    return { filteredPairs.begin(), filteredPairs.end() };
  }

  RangePairs findOverlapping(RangePairs rangePairs)
  {
    auto filteredPairs = rangePairs | filter(overlaps);
    return { filteredPairs.begin(), filteredPairs.end() };
  }

  string answer_a(const strings& input_data)
  {
    RangePairs rangePairs = buildRangePairs(input_data);
    RangePairs rangePairsWithOverlap = findCompletelyOverlapping(rangePairs);
    // In how many assignment pairs does one range fully contain the other?
    int count = rangePairsWithOverlap.size();
    return to_string(count);
  }

  string answer_b(const vector<string>& input_data)
  {
    RangePairs rangePairs = buildRangePairs(input_data);
    RangePairs rangePairsWithOverlap = findOverlapping(rangePairs);
    // In how many assignment pairs does one range overlap the other at all?
    int count = rangePairsWithOverlap.size();
    return to_string(count);
  }
}
