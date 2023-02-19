#include <days.h>

#include <cassert>
#include <functional>
#include <iterator>
#include <memory>
#include <optional>
#include <sstream>
#include <unordered_map>
#include <stdexcept>
#include <unordered_set>
#include <ranges>

#include <iostream>
#include <algorithm>
#include <pathfinder.h>
#include <deque>
#include <numeric>
#include <map>

namespace day_16
{
	using std::string;
	using std::vector;
	using std::pair;
	using std::invalid_argument;
	using std::unordered_set;
	using std::unordered_map;
	using std::map;
	using std::deque;
	namespace vw = std::views;
	using std::none_of;
	using std::all_of;
	using std::stoi;
	using std::accumulate;
	using std::prev_permutation;
	using std::max_element;
	using std::to_string;
	using std::transform;
	using std::sort;

	using advent_of_code::Pathfinder::floyd_warshall;
	using advent_of_code::Pathfinder::reconstruct_path;
	using advent_of_code::Pathfinder::FloydWarshallDistances;
	using advent_of_code::Pathfinder::FloydWarshallNexts;
	using advent_of_code::Pathfinder::IWeightedGraph;

	using NodeHandle = advent_of_code::Pathfinder::IGraph::NodeHandle;
	using NEIGHBOURS = unordered_set<string>;

	struct Valve
	{
		string name{};
		int flow{ 0 };
		NEIGHBOURS tunnels{};

		bool hasEdge(string name) const
		{
			return tunnels.contains(name);
		}
	};
	using VALVES = unordered_map<string, Valve>;

	class ValveGraph : public IWeightedGraph
	{
	public:
		const Valve& valve(NodeHandle handle) const
		{
			return nodes.at(name(handle));
		}
		double cost(NodeHandle from, NodeHandle to) const override
		{
			const Valve& fromNode = nodes.at(name(from));
			const Valve& toNode = nodes.at(name(to));

			if (!fromNode.hasEdge(toNode.name)) throw invalid_argument("from and to must be adjacent");

			return 1;
		}
		vector<NodeHandle> allNodes() const override
		{
			vector<NodeHandle> result(nodes.size());
			transform(nodes.begin(), nodes.end(), result.begin(),
					  [&](auto& pair) { return handle(pair.second.name); });
			return result;
		}
		size_t numberOfNodes() const override
		{
			return nodes.size();
		}
		vector<NodeHandle> neighbours(NodeHandle handle) const override
		{
			const Valve& node = nodes.at(name(handle));

			vector<NodeHandle> result(node.tunnels.size());
			transform(node.tunnels.begin(),
					  node.tunnels.end(),
					  result.begin(),
					  [&](string name) { return findNode(name); });
			return result;
		}

		NodeHandle findNode(const string& name) const
		{
			if (!nodes.contains(name)) return nullptr;

			return handle(nodes.at(name).name);
		}

		NodeHandle addNode(const Valve& node)
		{
			if (nodes.contains(node.name)) throw invalid_argument("id already exists");

			nodes[node.name] = node;
			return handle(nodes.at(node.name).name);
		}

	private:
		VALVES nodes;

		NodeHandle handle(const string& name) const
		{
			return reinterpret_cast<NodeHandle>(&name);
		}

		const string& name(NodeHandle handle) const
		{
			return *reinterpret_cast<const string*>(handle);
		}

	};

	int parseInt(auto& it, auto end)
	{
		size_t size{ 0 };
		int parsed = stoi(string{ it, end }, &size);
		it += size;
		return parsed;
	}

	Valve parseValve(const string& str)
	{
		auto it = str.begin();
		auto end = str.end();

		assert(*it == 'V');
		it += 6;                 // Valve_
		string name{ it, it + 2 }; // AA
		it += 2;
		assert(*it == ' ');
		it += 15; // _has_flow_rate=
		int flowRate = parseInt(it, end);
		assert(*it == ';');
		it += 23;             // ;_tunnels_lead_to_valve
		if (*it == 's') ++it; // s
		vector<string> tunnels;
		while (it != end) { // DD, II, BB
			if (*it >= 'A' && *it <= 'Z') {
				tunnels.push_back({ it, it + 2 });
				it += 2;
			} else {
				++it;
			}
		}

		return Valve{ name, flowRate, NEIGHBOURS(tunnels.begin(), tunnels.end()) };
	}

	VALVES parseValves(const vector<string>& data)
	{
		VALVES valves;
		for (auto& s : data) {
			Valve v = parseValve(s);
			valves[v.name] = v;
		}
		return valves;
	}

	void printDistances(const ValveGraph& graph, const FloydWarshallDistances& distances)
	{
		for (auto& handleV : graph.allNodes()) {
			std::cout << "Distances from " << graph.valve(handleV).name << "\n";
			for (auto& handleU : graph.allNodes()) {
				if (handleV == handleU) continue;
				std::cout
					<< "\t to "
					<< graph.valve(handleU).name << ": "
					<< distances.at(handleV).at(handleU) << "\n";
			}
		}
	}

	void printPaths(const ValveGraph& graph, const FloydWarshallNexts& nexts)
	{
		for (auto& handleV : graph.allNodes()) {
			std::cout << "Paths from " << graph.valve(handleV).name << "\n";
			for (auto& handleU : graph.allNodes()) {
				if (handleV == handleU) continue;
				std::cout << "\t to " << graph.valve(handleU).name << ": ";
				auto path = reconstruct_path(handleV, handleU, nexts);
				for (auto& handle : path) {
					std::cout << graph.valve(handle).name << " ";
				}
				std::cout << "(" << path.size() << ")\n";
			}
		}
	}

	struct State
	{
		NodeHandle current;
		unordered_set<NodeHandle> opened;
		int elapsed_seconds;
		int relieved_pressure;
	};

	auto relieved_per_min(const unordered_set<NodeHandle>& opened, const ValveGraph& graph)
	{
		auto relieved_per_min_by_valve = opened
			| vw::transform([&graph](auto handle) { return graph.valve(handle).flow; })
			| vw::common;
		return accumulate(
			relieved_per_min_by_valve.begin(),
			relieved_per_min_by_valve.end(),
			0);
	}

	int wait_until_end(int end_time, int elapsed_seconds, int relieved_pressure, const unordered_set<NodeHandle>& opened, const ValveGraph& graph)
	{
		auto time_left = end_time - elapsed_seconds;
		return relieved_pressure + time_left * relieved_per_min(opened, graph);
	}

	// requires a < for unordered_set<NodeHandle> to order them in the map
	// using RELIEVED_STATES_MAP = map<unordered_set<NodeHandle>, int>;

	// requires a hasher for unordered_set<NodeHandle> to lookup by hash
	struct hash_us_nh
	{
		size_t operator()(const unordered_set<NodeHandle>& us_nh) const
		{
			// hash each of the node handles, sort the hashes, concatenate the sorted hashes, hash the concatenation

			// use the fact that we know the node handles are actually pointers to strings
			vector<string> names(us_nh.size());
			// convert the node_handles to copies of their strings
			transform(us_nh.begin(), us_nh.end(), names.begin(), [](const NodeHandle& nh) { return *reinterpret_cast<const string*>(nh); });
			// sort the strings
			sort(names.begin(), names.end());
			string concat = accumulate(names.begin(), names.end(), string{});

			std::hash<string> hasher;
			return hasher(concat);
		}
	};
	using RELIEVED_STATES_MAP = unordered_map<unordered_set<NodeHandle>, int, hash_us_nh>;
	using NODE_HANDLE_LIST = unordered_set<NodeHandle>;
	using RELIEVED_STATES_PAIR = pair<const NODE_HANDLE_LIST, int>;
	using COMBO = vector<RELIEVED_STATES_PAIR>;

	// implements nCr for container of size n
	auto best_combination(const RELIEVED_STATES_MAP& container, size_t r) -> int
	{
		auto n = container.size();
		std::cout << "Building combinations for " << n << "C" << r << "\n";

		auto vecDisjoint = [](const NODE_HANDLE_LIST& lhs, const NODE_HANDLE_LIST& rhs) -> bool {
			for (const NodeHandle& lhsNH : lhs) {
				for (const NodeHandle& rhsNH : rhs) {
					if (lhsNH == rhsNH) return false;
				}
			}
			return true;
		};
		auto comboDisjoint = [vecDisjoint](const COMBO& vec) {
			for (auto itLHS = vec.begin(); itLHS != vec.end() - 1; ++itLHS) {
				for (auto itRHS = itLHS + 1; itRHS != vec.end(); ++itRHS) {
					if (!vecDisjoint(itLHS->first, itRHS->first)) return false;
				}
			}
			return true;
		};
		auto comboValue = [](const COMBO& combo) {
			return accumulate(combo.begin(), combo.end(), 0, [](int acc, auto& pair) { return acc + pair.second; });
		};

		// r leading 1's, n-r trailing 0's (eg. 11000)
		string bitmask(r, 1);
		bitmask.resize(n, 0);

		int bestCombinationValue = 0;

		// permute bitmask
		int checked = 0;
		do {
			COMBO thisCombination;
			auto it = container.begin();
			auto end = container.end();

			for (int i = 0; i < n && it != end; ++i) // [0..N-1] integers
			{
				if (bitmask[i]) {
					thisCombination.push_back(*it);
				}
				++it;
			}
			++checked;
			//if (checked % 1000000 == 0) std::cout << "M";
			//else if (checked % 100000 == 0) std::cout << "H";
			//else if (checked % 1000 == 0) std::cout << ".";

			if (!comboDisjoint(thisCombination)) continue;

			int thisCombinationValue = comboValue(thisCombination);
			if (bestCombinationValue < thisCombinationValue) {
				bestCombinationValue = thisCombinationValue;
				//std::cout << "(" << bestCombinationValue << ")";
			}

		} while (prev_permutation(bitmask.begin(), bitmask.end()));

		std::cout << "\n Done checking " << checked << "combinations.\n";

		return bestCombinationValue;
	}

	int solve(const VALVES& valves, int end_time = 30, int helpers = 0)
	{
		ValveGraph graph; for (auto& valve : valves) { graph.addNode(valve.second); }
		FloydWarshallDistances distances;	// distances[u][v] = (double) distance from u to v
		FloydWarshallNexts nexts;			// nexts[u][target] = (handle) next step from u to get to target
		floyd_warshall(graph, distances, nexts);
		//printDistances(graph, distances);	// checks out
		//printPaths(graph, nexts);			// checks out

		auto hasFlow = [&graph](auto& handle) { return graph.valve(handle).flow > 0; };
		auto flowing_view = graph.allNodes() | vw::filter(hasFlow);
		auto flowing = vector<NodeHandle>(flowing_view.begin(), flowing_view.end());

		RELIEVED_STATES_MAP max_relieved_states;

		auto maxRelieved = 0;
		auto queue = deque<State>();

		queue.push_back(State{ graph.findNode("AA"), {}, 0, 0 });

		while (!queue.empty()) {
			State state = queue.front(); queue.pop_front();

			auto relieved_at_end = wait_until_end(end_time, state.elapsed_seconds, state.relieved_pressure, state.opened, graph);
			// record state. only update state if it beats the `relieved_at_end` number
			if (max_relieved_states.contains(state.opened) && max_relieved_states.at(state.opened) < relieved_at_end) {
				max_relieved_states[state.opened] = relieved_at_end;
			} else {
				max_relieved_states.insert({ state.opened, relieved_at_end });
			}

			// if all flowing valves are opened, wait until the end
			if (state.opened.size() == flowing.size() || state.elapsed_seconds >= end_time) {
				continue;
			}

			// for every unopened valve, run simulation
			auto handleNotInOpened = [&state](auto& handle) {
				return none_of(
					state.opened.begin(),
					state.opened.end(),
					[&handle](auto& openHandle) {
						return openHandle == handle;
					}
				);
			};
			auto unopened = flowing | vw::filter(handleNotInOpened);

			for (auto& destination : unopened) {
				auto cost = int(distances[state.current][destination] + 1);
				auto newElapsed = state.elapsed_seconds + cost;
				// if opening the dest valve would exceed the time limit, wait until the end
				if (newElapsed >= end_time) {
					continue;
				}
				// relieve pressure of existing opened valves while we move to dest and open it
				auto newRelieved = state.relieved_pressure + relieved_per_min(state.opened, graph) * cost;

				unordered_set<NodeHandle> newOpened(state.opened);
				newOpened.insert(destination);

				queue.push_back(State{ destination, newOpened, newElapsed, newRelieved });
			}
		}

		if (helpers > 0) {
			return best_combination(
				max_relieved_states,
				static_cast<size_t>(helpers) + 1
			);
		} else {
			return max_element(
				max_relieved_states.begin(),
				max_relieved_states.end(),
				[](auto& lhsPair, auto& rhsPair) {
					return lhsPair.second < rhsPair.second;
				}
			)->second;
		}
	}

	string answer_a(const vector<string>& input_data)
	{
		return to_string(solve(parseValves(input_data), 30));
	}

	string answer_b(const vector<string>& input_data)
	{
		return to_string(solve(parseValves(input_data), 26, 1));
	}
}
