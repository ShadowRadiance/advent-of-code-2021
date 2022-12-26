#include <string>
#include <vector>
#include <functional>

class CPU_10
{
public:
    CPU_10(const std::vector<std::string>& instructions);
    int x() const;
    int current_cycle() const;
    int signal_strength() const;
    void begin_cycle();
    void end_cycle();
    int cycles_required() const;
private:
    using Operation = std::function<void(void)>;
    using CostedOperation = std::tuple<Operation, int>;
    CostedOperation operation(const std::string& instruction);
    const CostedOperation& currentOperation() const;

    std::vector<std::string> instructions_;
    int x_;
    int current_cycle_;
    int instruction_index_;
    int last_execution_cycle_;
    std::vector<CostedOperation> operations_;
};