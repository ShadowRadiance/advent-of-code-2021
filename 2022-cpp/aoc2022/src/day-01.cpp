#include <format>
#include <algorithm>
#include <numeric>

#include <days.h>

namespace day_01
{
  using std::string;
  using std::vector;
  using std::format;

  class Elf {
  public:
    uint32_t totalFood() const { return totalFood_; }
    void addFood(uint32_t calories) {
      food_.push_back(calories);
      totalFood_ += calories;
    }
  private:
    vector<uint32_t> food_;
    uint32_t totalFood_{ 0 };
  };

  bool lesserTotalFood(const Elf& first, const Elf& second) {
    return first.totalFood() < second.totalFood();
  };

  bool greaterTotalFood(const Elf& first, const Elf& second) {
    return first.totalFood() > second.totalFood();
  };

  vector<Elf> initializeElves(const vector<string>& input_data) {
    vector<Elf> elves;

    Elf anElf{};
    for (auto& s : input_data) {
      if (s.empty()) {
        elves.push_back(anElf);
        anElf = Elf{};
      }
      else {
        anElf.addFood(stoul(s));
      }
    }
    elves.push_back(anElf);
    return elves;
  }


  string answer_a(const vector<string>& input_data)
  {
    // total Calories carried by the top Elf carrying the most Calories

    auto elves = initializeElves(input_data);

    auto it = max_element(elves.begin(), elves.end(), lesserTotalFood);

    return format("{}", it->totalFood());
  }

  string answer_b(const vector<string>& input_data)
  {
    // total Calories carried by the highest three Elves carrying the most Calories

    auto elves = initializeElves(input_data);

    auto begin = elves.begin();
    auto third = elves.begin() + 3;
    auto end = elves.end();

    nth_element(begin, third, end, greaterTotalFood);

    vector<uint32_t> totals;
    transform(begin, third, back_inserter(totals),
      [](auto& anElf) { return anElf.totalFood(); }
    );
    return format("{}",
      accumulate(totals.begin(), totals.end(), 0)
    );
  }
}
