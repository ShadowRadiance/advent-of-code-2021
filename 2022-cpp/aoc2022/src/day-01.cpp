#include <vector>
#include <string>
#include <format>
#include <algorithm>
#include <numeric>

#include <day-01.h>

namespace day_01
{
	class Elf {
	public:
		uint32_t totalFood() const { return totalFood_; }
		void addFood(uint32_t calories) {
			food_.push_back(calories);
			totalFood_ += calories;
		}
	private:
		std::vector<uint32_t> food_;
		uint32_t totalFood_{ 0 };
	};

	bool lesserTotalFood(const Elf& first, const Elf& second) {
		return first.totalFood() < second.totalFood();
	};

	bool greaterTotalFood(const Elf& first, const Elf& second) {
		return first.totalFood() > second.totalFood();
	};

	std::vector<Elf> initializeElves(const std::vector<std::string>& input_data) {
		std::vector<Elf> elves;

		Elf anElf{};
		for (auto& s : input_data) {
			if (s.empty()) {
				elves.push_back(anElf);
				anElf = Elf{};
			}
			else {
				anElf.addFood(std::stoul(s));
			}
		}
		elves.push_back(anElf);
		return elves;
	}


	std::string answer_a(const std::vector<std::string>& input_data)
	{
		// total Calories carried by the top Elf carrying the most Calories

		auto elves = initializeElves(input_data);

		auto it = std::max_element(elves.begin(), elves.end(), lesserTotalFood);

		return std::format("{}", it->totalFood());
	}

	std::string answer_b(const std::vector<std::string>& input_data)
	{
		// total Calories carried by the highest three Elves carrying the most Calories

		auto elves = initializeElves(input_data);

		auto begin = elves.begin();
		auto third = elves.begin() + 3;
		auto end = elves.end();

		std::nth_element(begin, third, end, greaterTotalFood);

		std::vector<uint32_t> totals;
		std::transform(begin, third, std::back_inserter(totals),
			[](auto& anElf) { return anElf.totalFood(); }
		);
		return std::format("{}", 
			std::accumulate(totals.begin(), totals.end(), 0)
		);
	}
}
