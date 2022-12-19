#include <vector>
#include <string>
#include <format>
#include <algorithm>
#include <numeric>

#include <day-01.h>

namespace day_01
{
	std::string answer_a(const std::vector<std::string>& input_data)
	{
		std::vector<elf> elves;

		elf anElf{};
		for (auto& s : input_data) {
			if (s.empty()) {
				elves.push_back(anElf);
				anElf = elf{};
			} else {
				anElf.add_food(std::stoul(s));
			}
		}
		elves.push_back(anElf);

		auto lt_by_total_food = [](const elf& first, const elf& second) {
			return first.total_food() < second.total_food();
		};
		auto it = std::max_element(elves.begin(), elves.end(), lt_by_total_food);

		return std::format("{}", it->total_food());
	}
	std::string answer_b(const std::vector<std::string>& input_data)
	{
		return "";
	}

	void elf::add_food(uint32_t calories)
	{
		food.push_back(calories);
	}
	
	uint32_t elf::total_food() const
	{
		return std::accumulate(std::begin(food), std::end(food), 0);
	}
}
