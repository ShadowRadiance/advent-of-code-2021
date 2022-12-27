#include <days.h>

#include <functional>
#include <numeric>
#include <sstream>
#include <regex>
#include <cassert>

namespace day_11
{
    using std::string;
    using std::vector;
    using std::function;
    using std::regex;
    using std::smatch;
    using std::sregex_token_iterator;
    using std::to_string;

    using strings = vector<string>;

    int64_t add(int64_t old, int64_t amount) {
        assert(amount < INT64_MAX - old);
        return old + amount;
    }
    int64_t multiply(int64_t old, int64_t amount) {
        assert(amount != 0);
        assert(amount < INT64_MAX / old);
        return old * amount;
    }
    int64_t square(int64_t old, int64_t _) {
        return multiply(old, old);
    }

    strings split(const std::string& str, const regex& regex) {
        return strings{
            sregex_token_iterator(str.begin(), str.end(), regex, -1),
            sregex_token_iterator()
        };
    }

    template<typename T>
    T pop_front(vector<T>& vec) {
        T item = vec.front();
        vec.erase(vec.begin());
        return item;
    }

    const int64_t UNMATCHED = -1;

    class Monkey;
    using Monkeys = vector<Monkey>;

    class Monkey
    {
    public:
        Monkey() = default;
        Monkey(const strings& lines, int64_t calmness)
            : calmnessDivisor(calmness)
        {
            // Example:
            // Monkey 1:
            //   Starting items : 54, 65, 75, 74
            //       Operation : new = old + 6
            //       Test : divisible by 19
            //       If true : throw to monkey 2
            //       If false : throw to monkey 0

            vector<regex> regexes = {
                regex{ "Monkey (\\d+) ?:" },
                regex{ "\\s*Starting items ?: (.*)" },
                regex{ "\\s*Operation ?: new = old ([+*]) (old|\\d+)" },
                regex{ "\\s*Test ?: divisible by (\\d+)" },
                regex{ "\\s*If true ?: throw to monkey (\\d+)" },
                regex{ "\\s*If false ?: throw to monkey (\\d+)" },
            };
            vector<smatch> matches{ 6 };

            for (int64_t i{ 0 }; i < 6; i++) {
                regex_match(lines[i], matches[i], regexes[i]);
            }

            id = stoi(matches[0][1].str());
            string startingItems = matches[1][1].str();
            // split(startingItems, regex{ ", " });
            regex split{ ", ?" };
            transform(
                std::sregex_token_iterator(
                    startingItems.begin(),
                    startingItems.end(),
                    split, UNMATCHED
                ),
                std::sregex_token_iterator(),
                back_inserter(itemWorryLevels),
                [](const string& s) { return stoi(s); }
            );

            string operation = matches[2][1].str();
            string next = matches[2][2].str();
            if (operation == "+") {
                worryOperation = add;
                worryOperationAmount = stoi(next);
            }
            else if (operation == "*") {
                if (next == "old") {
                    worryOperation = square;
                    worryOperationAmount = 0;
                }
                else
                {
                    worryOperation = multiply;
                    worryOperationAmount = stoi(next);
                }
            }
            divisibleByTest = stoi(matches[3][1].str());
            targetMonkeyIfTrue = stoi(matches[4][1].str());
            targetMonkeyIfFalse = stoi(matches[5][1].str());

            commonModulus *= divisibleByTest;
        }

        void catchItem(int64_t item) {
            if (item % commonModulus != 0) {
                item = item % commonModulus;
            }
            itemWorryLevels.push_back(item);
        }

        void playAround(Monkeys& allTheMonkeys) {
            while (itemWorryLevels.size() > 0) {
                int64_t item = pop_front(itemWorryLevels);
                item = worryOperation(item, worryOperationAmount);

                if (calmnessDivisor > 1) item = item / calmnessDivisor;
                int64_t targetMonkey = (item % divisibleByTest == 0)
                    ? targetMonkeyIfTrue
                    : targetMonkeyIfFalse;
                allTheMonkeys[targetMonkey].catchItem(item);
                inspectionsMade += 1;
            }
        }

        int64_t inspections() const { return inspectionsMade; }

        static void resetCommonModulus() { commonModulus = 1; }
    private:
        int64_t id{ 0 };
        vector<int64_t> itemWorryLevels;
        function<int64_t(int64_t, int64_t)> worryOperation;
        int64_t worryOperationAmount{ 0 };
        int64_t divisibleByTest{ 0 };
        int64_t targetMonkeyIfTrue{ 0 };
        int64_t targetMonkeyIfFalse{ 0 };
        int64_t calmnessDivisor;

        int64_t inspectionsMade{ 0 };

        static int64_t commonModulus;
    };

    int64_t Monkey::commonModulus = 1;

    Monkeys initializeMonkeys(const strings& input_data, int64_t calmness = 3) {
        // remember to reset static data on each run!
        Monkey::resetCommonModulus();
        Monkeys monkeys;

        strings::const_iterator itMonkeyData = input_data.begin();
        while (itMonkeyData < input_data.end()) {
            monkeys.push_back(Monkey{ strings{itMonkeyData, itMonkeyData + 6 }, calmness });

            itMonkeyData += 6; // 6 lines per monkey data
            if (itMonkeyData != input_data.end()) itMonkeyData += 1; // 1 empty line
        }

        return monkeys;
    }

    void playAround(int64_t rounds, Monkeys& monkeys) {
        for (int64_t round = 0; round < rounds; round++) {

            for (Monkey& monkey : monkeys) {
                monkey.playAround(monkeys);
            }
        }
    }

    int64_t monkeyBusiness(Monkeys& monkeys) {
        sort(monkeys.begin(), monkeys.end(),
            [](Monkey& a, Monkey& b) {
                return a.inspections() > b.inspections();
            }
        );
        return monkeys[0].inspections() * monkeys[1].inspections();
    }

    string answer_a(const strings& input_data)
    {
        Monkeys monkeys = initializeMonkeys(input_data);
        playAround(20, monkeys);
        return to_string(monkeyBusiness(monkeys));
    }

    string answer_b(const strings& input_data)
    {
        const int64_t calmness = 1;
        Monkeys monkeys = initializeMonkeys(input_data, calmness);
        playAround(10000, monkeys);
        return to_string(monkeyBusiness(monkeys));
    }
}
