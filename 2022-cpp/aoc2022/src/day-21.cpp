#include <days.h>
#include <memory>
#include <algorithm>
#include <iterator>
#include <unordered_map>

namespace day_21
{
    using std::string;
    using std::vector;
    using std::unordered_map;
    using std::shared_ptr;

    class Monkey;
    using MonkeyMap = unordered_map<string, shared_ptr<Monkey>>;

    class Monkey
    {
    public:
        Monkey(string name)
            : name(name)
        {}
        virtual int64_t evaluate() const = 0;
        virtual ~Monkey() = default;
    private:
        string name;
    };

    class NumberMonkey : public Monkey
    {
    public:
        NumberMonkey(string name, int64_t number)
            : Monkey(name), number(number)
        {}
        int64_t evaluate() const override
        {
            return number;
        }
    private:
        int64_t number;
    };

    class OperationMonkey : public Monkey
    {
    public:
        OperationMonkey(string name, string first, char op, string second)
            : Monkey(name), first(first), op(op), second(second)
        {}
        int64_t evaluate() const override
        {
            switch (op) {
            case '+': return firstMonkey->evaluate() + secondMonkey->evaluate();;
            case '-': return firstMonkey->evaluate() - secondMonkey->evaluate();
            case '*': return firstMonkey->evaluate() * secondMonkey->evaluate();
            case '/': return firstMonkey->evaluate() / secondMonkey->evaluate();
            default: return 0;
            }
        }
        void connectSources(MonkeyMap const& monkeyMap)
        {
            firstMonkey = monkeyMap.at(first);
            secondMonkey = monkeyMap.at(second);
        }
    private:
        string first;
        char op;
        string second;
        shared_ptr<Monkey> firstMonkey;
        shared_ptr<Monkey> secondMonkey;
    };

    int64_t parseInt(auto& it, auto end)
    {
        size_t size{ 0 };
        int parsed = std::stoll(string{ it, end }, &size);
        it += size;
        return parsed;
    }

    MonkeyMap::value_type
        parseMonkey(string const& str)
    {
        // name: first op last
        // name: number

        auto it = str.begin();
        auto end = str.end();

        string name(it, it + 4);
        it += 6;

        if (std::isdigit(*it)) {
            int64_t number = parseInt(it, end);
            return { name, std::make_shared<NumberMonkey>(name, number) };
        } else {
            string first(it, it + 4);
            it += 5;
            char op = *it;
            it += 2;
            string last(it, it + 4);
            return { name, std::make_shared<OperationMonkey>(name, first, op, last) };
        }
    }

    string answer_a(vector<string> const& inputData)
    {
        MonkeyMap monkeys;
        std::transform(
            inputData.begin(), inputData.end(),
            std::inserter(monkeys, monkeys.end()),
            parseMonkey
        );
        std::for_each(
            monkeys.begin(), monkeys.end(),
            [&monkeys](auto& pair) {
                auto& [name, monkey] = pair;
                auto opMonkey = dynamic_pointer_cast<OperationMonkey>(monkey);
                if (opMonkey) opMonkey->connectSources(monkeys);
            }
        );

        return std::to_string(monkeys["root"]->evaluate());
    }

    string answer_b(vector<string> const& inputData)
    {
        return "PENDING";
    }
}
