#pragma once

#include <vector>
#include <string>

namespace day_01 {
	struct elf {
		std::vector<uint32_t> food;
		void add_food(uint32_t calories);
		uint32_t total_food() const;
	};

	std::string answer_a(const std::vector<std::string>& input_data);
	std::string answer_b(const std::vector<std::string>& input_data);
}

