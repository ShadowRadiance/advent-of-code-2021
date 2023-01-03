#include <days.h>
#include <variant>
#include <optional>
#include <utility>
#include <memory>
#include <sstream>
#include <cassert>
#include <numeric>
#include <algorithm>
#include <array>
#include <iostream>

namespace Packet13
{
    using std::string;
    using std::vector;
    using std::variant;

    template<class... Ts> struct overloaded : Ts... { using Ts::operator()...; };
    //template<class... Ts> overloaded(Ts...) -> overloaded<Ts...>;

    struct Node
    {
        variant<int, vector<Node>> data;

        explicit Node(int i) : data(i) {}
        explicit Node(vector<Node> v) : data(std::move(v)) {}

        friend bool operator==(Node const& lhs, Node const& rhs)
        {
            auto overloads = overloaded{
                [](int l, int r) { return l == r; },
                [](vector<Node> const& l, vector<Node> const& r) { return l == r; },
                [](int l, vector<Node> const& r) { return vector<Node>{Node{l}} == r; },
                [](vector<Node> const& l, int r) { return l == vector<Node>{Node{r}}; },
            };
            return std::visit(overloads, lhs.data, rhs.data);
        }

        friend bool operator<(Node const& lhs, Node const& rhs)
        {
            auto overloads = overloaded{
                [](int l, int r) { return l < r; },
                [](vector<Node> const& l, vector<Node> const& r) { return l < r; },
                [](int l, vector<Node> const& r) { return vector<Node>{Node{l}} < r; },
                [](vector<Node> const& l, int r) { return l < vector<Node>{Node{r}}; },
            };
            return std::visit(overloads, lhs.data, rhs.data);
        }

        friend auto operator<<(std::ostream& os, Node& node) -> std::ostream& {
            if (std::holds_alternative<vector<Node>>(node.data)) {
                os << "[";
                bool firstTime = true;
                for (Node& node : std::get<vector<Node>>(node.data)) {
                    if (!firstTime) os << ",";
                    os << node;
                    firstTime = false;                    
                }
                os << "]";
            }
            else {
                os << std::get<int>(node.data);
            }
            return os;
        }
    };

    int parseInt(auto& it, auto end)
    {
        size_t size{0};
        int parsed = std::stoi(string{ it, end }, &size);
        it += size;
        return parsed;
    }

    vector<Node> parseList(auto& it, auto end)
    {
        assert(*it == '[');
        ++it;

        vector<Node> nodes;

        while (it != end && *it != ']') {
            if (std::isdigit(*it)) {
                nodes.emplace_back(parseInt(it, end)); // parseInt will advance it
            }
            else if (*it == '[') {
                nodes.emplace_back(parseList(it, end)); // recurse,  advancing it
            }
            if (*it == ',') ++it;
        }


        assert(*it == ']');
        ++it;

        return nodes;
    }


    Node parseNode(string s)
    {
        auto it = s.begin();
        auto end = s.end();
        vector<Node> children = parseList(it, end);
        return Node{ children };
    }
}

namespace day_13
{
    using std::string;
    using std::vector;
    using std::pair;
    
    using Packet13::Node;
    using Packet13::parseNode;

    struct NodePair
    {
        size_t number = 0;
        pair<Node, Node> pair;
    };
    using NodePairs = vector<NodePair>;
    using Nodes = vector<Node>;

    NodePairs parseNodePairs(vector<string> const& input)
    {
        vector<string> filtered;
        copy_if(input.begin(), input.end(), back_inserter(filtered), [](auto& s) { return !s.empty();  });

        vector<Node> allNodes;
        std::transform(filtered.begin(), filtered.end(), back_inserter(allNodes), parseNode);

        NodePairs result;
        auto it = allNodes.begin();
        auto end = allNodes.end();
        while (it != end) {
            result.emplace_back(result.size()+1, std::make_pair(*it, *(it + 1)));
            it += 2;
        }

        return result;
    }

    Nodes parseNodes(vector<string> const& input) {
        vector<string> filtered;
        copy_if(input.begin(), input.end(), back_inserter(filtered), [](auto& s) { return !s.empty();  });

        vector<Node> allNodes;
        std::transform(filtered.begin(), filtered.end(), back_inserter(allNodes), parseNode);

        return allNodes;
    }

    bool ordered(NodePair& nodePair)
    {
        return nodePair.pair.first < nodePair.pair.second;
    }

    string answer_a(const vector<string>& input_data)
    {
        NodePairs nodePairs = parseNodePairs(input_data);
        //std::cout << "Pairs 1, 2, 4, and 6 should be considered ordered correctly." << std::endl;

        int sum = 0;
        for (NodePair& nodePair : nodePairs)
        {
            //std::cout << "PACKET " << nodePair.number << "A:" << nodePair.pair.first << std::endl;
            //std::cout << "PACKET " << nodePair.number << "B:" << nodePair.pair.second << std::endl;

            if (ordered(nodePair)) {
                //std::cout << "Pair " << nodePair.number << " is considered ordered." << std::endl;
                sum += nodePair.number;
            }
            //std::cout << std::endl;
        }
        return std::to_string(sum);
    }

    string answer_b(const vector<string>& input_data)
    {
        vector<string> modified_data{ input_data.begin(), input_data.end() };
        modified_data.push_back("[[2]]");
        modified_data.push_back("[[6]]");

        Nodes nodes = parseNodes(modified_data);

        Node twoSeparator = parseNode("[[2]]");
        Node sixSeparator = parseNode("[[6]]");

        std::sort(nodes.begin(), nodes.end());

        auto itSeparator2 = std::find(nodes.begin(), nodes.end(), twoSeparator);
        auto itSeparator6 = std::find(nodes.begin(), nodes.end(), sixSeparator);

        int index2 = std::distance(nodes.begin(), itSeparator2) + 1;
        int index6 = std::distance(nodes.begin(), itSeparator6) + 1;

        return std::to_string(index2 * index6);
    }
}
