#include <days.h>
#include <memory>
#include <algorithm>
#include <iterator>
#include <unordered_map>
#include <stack>
#include <cassert>
#include <iostream>

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
        string getName() const
        {
            return name;
        }
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
        int64_t getNumber() const
        {
            return number;
        }
        void setNumber(int64_t number)
        {
            this->number = number;
        }
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

        string getFirst() const
        {
            return first;
        }

        string getSecond() const
        {
            return second;
        }

        void setOperation(char op)
        {
            this->op = op;
        }

        char getOperation() const
        {
            return op;
        }

        int64_t evaluate() const override
        {
            switch (op) {
            case '+': return firstMonkey->evaluate() + secondMonkey->evaluate();;
            case '-': return firstMonkey->evaluate() - secondMonkey->evaluate();
            case '*': return firstMonkey->evaluate() * secondMonkey->evaluate();
            case '/': return firstMonkey->evaluate() / secondMonkey->evaluate();
                // for the == case we determine the difference... 
                // 0 means equal
                // -ve means secondMonkey bigger
                // +ve means firstMonkey bigger
            case '=': return firstMonkey->evaluate() - secondMonkey->evaluate();
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

    shared_ptr<Monkey> monkeyReferringTo(string name, MonkeyMap const& monkeys)
    {
        auto itReferrer = std::find_if(
            monkeys.begin(), monkeys.end(),
            [name](auto& pair) {
                auto& [_, monkey] = pair;
                auto opMonkey = dynamic_pointer_cast<OperationMonkey>(monkey);
                if (!opMonkey) return false;
                return
                    opMonkey->getFirst() == name ||
                    opMonkey->getSecond() == name;
            }
        );
        if (itReferrer == monkeys.end()) return {};
        return itReferrer->second;
    }

    string answer_b(vector<string> const& inputData)
    {
        // oh mannnn - communication is a bitch amirite?!

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

        auto root = std::static_pointer_cast<OperationMonkey>(monkeys["root"]);
        auto humn = std::static_pointer_cast<NumberMonkey>(monkeys["humn"]);
        root->setOperation('=');

        // what number (int64_t) do you have to yell such that 
        // monkeys["root"]->evaluate() will equal zero ("equal")

        humn->setNumber(0);

        std::stack<shared_ptr<Monkey>> stack;
        std::shared_ptr<Monkey> nullMonkey;
        shared_ptr<Monkey> current = humn;
        while (current != nullMonkey) {
            std::cout << current->getName() << std::endl;
            stack.push(current);
            current = monkeyReferringTo(current->getName(), monkeys);
        }
        assert(stack.top() == root);

        int64_t target = 0;
        while (!stack.empty()) {
            shared_ptr<Monkey> current = stack.top(); stack.pop();

            auto opCurrent = std::dynamic_pointer_cast<OperationMonkey>(current);
            if (opCurrent) {
                if (opCurrent->getFirst() == stack.top()->getName()) {
                    auto secondTotal = monkeys[opCurrent->getSecond()]->evaluate();
                    switch (opCurrent->getOperation()) {
                    case '+': target -= secondTotal; break;         // target^: target+second  ==> target = target^-second
                    case '-': target += secondTotal; break;         // target^: target-second  ==> target = target^+second
                    case '*': target /= secondTotal; break;         // target^: target*second  ==> target = target^/second
                    case '/': target *= secondTotal; break;         // target^: target/second  ==> target = target^*second
                    case '=': target = secondTotal; break;
                    }
                } else {
                    auto firstTotal = monkeys[opCurrent->getFirst()]->evaluate();
                    switch (opCurrent->getOperation()) {
                    case '+': target -= firstTotal; break;          // target^: first+target  ==> target = target^-first
                    case '-': target = firstTotal - target; break;  // target^: first-target  ==> target = first-target^
                    case '*': target /= firstTotal; break;          // target^: first*target  ==> target = target^/first
                    case '/': target = firstTotal / target; break;  // target^: first/target  ==> target = first/target^
                    case '=': target = firstTotal; break;
                    }
                }
            }
        }

        return std::to_string(target);
    }
}
