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
    using std::stringstream;
    using std::to_string;
    using std::includes;
    using std::set_intersection;

    class Range
    {
        int start_, end_;
    public:
        Range(int start, int end) : start_(start), end_(end) {}

        bool overlaps(const Range& other)
        {
            return
                other.start_ >= this->start_ && other.start_ <= this->end_ // start within
                ||
                other.end_ >= this->start_ && other.end_ <= this->end_; // end within
        }

        bool covers(const Range& other)
        {
            return
                other.start_ >= this->start_ // start later
                &&
                other.end_ <= this->end_;   // end earlier
        }
    };

    using strings = vector<string>;
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
          Range(start1, end1),
          Range(start2, end2),
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
        return rangePair[0].covers(rangePair[1]) || rangePair[1].covers(rangePair[0]);
    }

    bool overlaps(RangePair rangePair)
    {
        return rangePair[0].overlaps(rangePair[1]);
    }

    RangePairs findCompletelyOverlapping(RangePairs rangePairs)
    {
        RangePairs filteredPairs;
        copy_if(rangePairs.begin(), rangePairs.end(), back_inserter(filteredPairs), completelyOverlaps);
        return { filteredPairs.begin(), filteredPairs.end() };
    }

    RangePairs findOverlapping(RangePairs rangePairs)
    {
        RangePairs filteredPairs;
        copy_if(rangePairs.begin(), rangePairs.end(), back_inserter(filteredPairs), overlaps);
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
