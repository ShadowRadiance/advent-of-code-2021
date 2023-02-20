#include <cpu_10.h>

#include <map>
#include <sstream>
#include <algorithm>
#include <tuple>
#include <numeric>

using std::string;
using std::vector;
using std::map;
using std::stringstream;
using std::tuple;
using std::make_tuple;

using strings = vector<string>;

const map<string, int> operations_costs{
    {"noop", 1},
    {"addx", 2},
};

CPU_10::CPU_10(const strings& instructions)
    : instructions_(instructions)
    , x_(1)
    , current_cycle_(0)
    , instruction_index_(0)
    , last_execution_cycle_(0)
{
    transform(
        instructions_.begin(), instructions_.end(),
        back_inserter(operations_),
        [&](const string& instruction) {
            return operation(instruction);
        }
    );
}

int CPU_10::cycles_required() const
{
    return accumulate(
        operations_.begin(), operations_.end(),
        0,
        [](int accum, const CostedOperation& op) {
            auto& [_, cost] = op;
    return accum + cost;
        }
    );
}

int CPU_10::x() const { return x_; }

int CPU_10::current_cycle() const { return current_cycle_; }

int CPU_10::signal_strength() const
{
    return x_ * current_cycle_;
}

void CPU_10::begin_cycle()
{
    current_cycle_ += 1;
}

void CPU_10::end_cycle()
{
    auto [op, cost] = currentOperation();
    bool completed = (cost == current_cycle_ - last_execution_cycle_);
    if (completed) {
        op();
        last_execution_cycle_ = current_cycle_;
        instruction_index_ += 1;
    }
}

const CPU_10::CostedOperation& CPU_10::currentOperation() const
{
    return operations_[instruction_index_];
}

CPU_10::CostedOperation CPU_10::operation(const string& instruction)
{
    stringstream ss(instruction);

    string command;
    int parameter;
    ss >> command;
    if (command == "noop") {
        return make_tuple(
            [=, this]() -> void { return; },
            operations_costs.at(command)
        );
    }
    if (command == "addx") {
        ss >> parameter;
        return make_tuple(
            [=, this]() -> void { x_ += parameter; },
            operations_costs.at(command)
        );
    }

    return make_tuple(
        [=, this]() -> void { return; },
        0
    );
}
