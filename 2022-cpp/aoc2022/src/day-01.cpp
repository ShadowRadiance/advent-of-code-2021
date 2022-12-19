#include <vector>
#include <string>
#include <format>
#include <algorithm>
#include <numeric>

#include <day-01.h>

namespace day_01
{
	bool lt_by_total_food(const elf& first, const elf& second) {
		return first.total_food() < second.total_food();
	};

	bool gt_by_total_food(const elf& first, const elf& second) {
		return first.total_food() > second.total_food();
	};

	std::vector<elf> initialize_elves(const std::vector<std::string>& input_data) {
		std::vector<elf> elves;

		elf anElf{};
		for (auto& s : input_data) {
			if (s.empty()) {
				elves.push_back(anElf);
				anElf = elf{};
			}
			else {
				anElf.add_food(std::stoul(s));
			}
		}
		elves.push_back(anElf);
		return elves;
	}


	std::string answer_a(const std::vector<std::string>& input_data)
	{
		// total Calories carried by the top Elf carrying the most Calories

		auto elves = initialize_elves(input_data);

		auto it = std::max_element(elves.begin(), elves.end(), lt_by_total_food);

		return std::format("{}", it->total_food());
	}

	std::string answer_b(const std::vector<std::string>& input_data)
	{
		// total Calories carried by the highest three Elves carrying the most Calories

		auto elves = initialize_elves(input_data);

		auto begin = elves.begin();
		auto third = elves.begin() + 3;
		auto end = elves.end();

		std::nth_element(begin, third, end, gt_by_total_food);

		std::vector<uint32_t> totals;
		std::transform(begin, third, std::back_inserter(totals),
			[](auto& anElf) { return anElf.total_food(); }
		);
		return std::format("{}", 
			std::accumulate(totals.begin(), totals.end(), 0)
		);
	}

	void elf::add_food(uint32_t calories)
	{
		food.push_back(calories);
	}
	
	uint32_t elf::total_food() const
	{
		return std::accumulate(food.begin(), food.end(), 0);
	}
}
