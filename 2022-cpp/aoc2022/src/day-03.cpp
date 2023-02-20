#include <array>
#include <algorithm>
#include <iterator>
#include <numeric>
#include <cassert>
#include <span>

#include <days.h>

namespace day_03
{
    using std::string;
    using std::vector;
    using std::span;
    using std::plus;
    using std::to_string;

    using strings = vector<string>;

    class Backpack
    {
    public:
        Backpack(const string& s = "") : contents_(s)
        {
            size_t halfLength = contents_.length() / 2;
            compartments_ = { contents_.substr(0, halfLength), contents_.substr(halfLength) };
            sort(compartments_[0].begin(), compartments_[0].end());
            sort(compartments_[1].begin(), compartments_[1].end());

            sort(contents_.begin(), contents_.end());
        }

        const strings& compartments() const { return compartments_; }
        const string& contents() const { return contents_; }
    private:
        string contents_;
        strings compartments_;
    };

    using Backpacks = vector<Backpack>;
    using Group = span<Backpack>;
    using Groups = vector<Group>;

    int scoreChar(char c);
    Backpack buildBackpack(const string& s);
    Backpacks buildBackpacks(const strings& input_data);
    Groups groupBackpacks(Group backpacks, size_t count);
    vector<char> determineCommonItemTypes(const Groups& groups);
    vector<char> determineCommonItemTypes(const Backpacks& backpacks);
    char determineCommonItemTypeAcrossBackpacks(Group group);
    char determineCommonItemTypeAcrossCompartments(const Backpack& backpack);
    std::string determineCommonCharacters(const strings& list);

    int scoreChar(char c)
    {
        // - Lowercase item types a through z have priorities 1 through 26.
        // - Uppercase item types A through Z have priorities 27 through 52.

        // ASCII a ==> INT 97, so subtract 96
        // ASCII A ==> INT 65, so subtract 38

        if (c >= 'a' && c <= 'z') return c - 96;
        if (c >= 'A' && c <= 'Z') return c - 38;

        return 0;
    }

    Backpack buildBackpack(const string& s) { return Backpack{ s }; }

    Backpacks buildBackpacks(const strings& input_data)
    {
        Backpacks backpacks;
        transform(input_data.begin(), input_data.end(), back_inserter(backpacks), buildBackpack);
        return backpacks;
    }

    Groups groupBackpacks(Group backpacks, size_t count)
    {
        Groups groups;
        size_t max = backpacks.size();
        for (size_t i = 0; i < max; i += count) {
            if (i + count > max)
                groups.push_back(backpacks.subspan(i));
            else
                groups.push_back(backpacks.subspan(i, count));
        }
        return groups;
    }

    vector<char> determineCommonItemTypes(const Backpacks& backpacks)
    {
        vector<char> commonItemTypes;
        transform(backpacks.begin(), backpacks.end(), back_inserter(commonItemTypes), determineCommonItemTypeAcrossCompartments);
        return commonItemTypes;
    }

    vector<char> determineCommonItemTypes(const Groups& groups)
    {
        vector<char> commonItemTypes;
        transform(groups.begin(), groups.end(), back_inserter(commonItemTypes), determineCommonItemTypeAcrossBackpacks);
        return commonItemTypes;
    }

    char determineCommonItemTypeAcrossCompartments(const Backpack& backpack)
    {
        string list = determineCommonCharacters(backpack.compartments());
        assert(!list.empty());
        return list[0];
    };

    char determineCommonItemTypeAcrossBackpacks(Group group)
    {
        strings backpackContents;
        transform(group.begin(), group.end(), back_inserter(backpackContents),
                  [](const Backpack& backpack) { return backpack.contents(); });
        string list = determineCommonCharacters(backpackContents);
        assert(!list.empty());
        return list[0];
    };

    string determineCommonCharacters(const strings& list)
    {
        if (list.empty()) return {};

        string commonSoFar = list[0];
        for (size_t i = 1U; i < list.size(); i++) {
            string next = list[i];
            string newResult;
            set_intersection(commonSoFar.begin(), commonSoFar.end(), next.begin(), next.end(), back_inserter(newResult));
            swap(commonSoFar, newResult);
        }
        return commonSoFar;
    }

    string answer_a(const strings& input_data)
    {
        Backpacks backpacks = buildBackpacks(input_data);
        vector<char> commonItemTypes = determineCommonItemTypes(backpacks);
        int score = transform_reduce(commonItemTypes.begin(), commonItemTypes.end(), 0, plus<>{}, scoreChar);

        return to_string(score);
    }

    string answer_b(const strings& input_data)
    {
        Backpacks backpacks = buildBackpacks(input_data);
        Groups groups = groupBackpacks(backpacks, 3);
        vector<char> badges = determineCommonItemTypes(groups);
        int score = transform_reduce(badges.begin(), badges.end(), 0, plus{}, scoreChar);

        return to_string(score);
    }
}
