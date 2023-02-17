#pragma once

#include <queue>
#include <unordered_map>
#include <unordered_set>
#include <vector>
#include <algorithm>
#include <functional>
#include <utility>

namespace advent_of_code
{
    namespace Pathfinder
    {
        using std::greater;
        using std::pair;
        using std::priority_queue;
        using std::queue;
        using std::reverse;
        using std::unordered_map;
        using std::unordered_set;
        using std::vector;

        template <typename Graph, typename Location>
        auto
        breadth_first_search(const Graph& graph, Location start, Location goal
        )
        {
            queue<Location> frontier;
            frontier.push(start);

            unordered_map<Location, Location> came_from;
            came_from[start] = start;

            while (!frontier.empty()) {
                Location current = frontier.front();
                frontier.pop();
                if (current == goal) break;
                // yield current if yield-fn provided
                for (Location next : graph.neighbours(current)) {
                    if (came_from.find(next) == came_from.end()) {
                        frontier.push(next);
                        came_from[next] = current;
                    }
                }
            }
            return came_from;
        }

        template <typename T, typename priority_t>
        struct PriorityQueue
        {
            typedef pair<priority_t, T> Element;
            priority_queue<Element, vector<Element>, greater<Element>> elements;

            bool empty() const { return elements.empty(); }

            void put(T item, priority_t priority)
            {
                elements.emplace(priority, item);
            }

            T get()
            {
                T best_item = elements.top().second;
                elements.pop();
                return best_item;
            }
        };

        template <typename WeightedGraph, typename Location>
        auto dijkstra_search(WeightedGraph graph,
                             Location start,
                             Location goal,
                             unordered_map<Location, Location>& came_from,
                             unordered_map<Location, double>& cost_so_far)
        {
            PriorityQueue<Location, double> frontier;
            frontier.put(start, 0);
            came_from[start] = start;
            cost_so_far[start] = 0;

            while (!frontier.empty()) {
                Location current = frontier.get();
                if (current == goal) break;
                for (Location next : graph.neighbours(current)) {
                    double new_cost = cost_so_far[current] +
                                      graph.cost(current, next);
                    if (cost_so_far.find(next) == cost_so_far.end() ||
                        new_cost < cost_so_far[next]) {
                        cost_so_far[next] = new_cost;
                        came_from[next] = current;
                        frontier.put(next, new_cost);
                    }
                }
            }
        }

        template <typename Location>
        vector<Location>
        reconstruct_path(Location start,
                         Location goal,
                         unordered_map<Location, Location> came_from)
        {
            vector<Location> path;
            Location current = goal;
            if (came_from.find(goal) == came_from.end()) return path;
            while (current != start) {
                path.push_back(current);
                current = came_from[current];
            }
            path.push_back(start);
            reverse(path.begin(), path.end());
            return path;
        }

        template <typename Graph, typename Location>
        void a_star_search(Graph graph,
                           Location start,
                           Location goal,
                           auto heuristic,
                           unordered_map<Location, Location>& came_from,
                           unordered_map<Location, double>& cost_so_far)
        {
            PriorityQueue<Location, double> frontier;
            frontier.put(start, 0);
            came_from[start] = start;
            cost_so_far[start] = 0;

            while (!frontier.empty()) {
                Location current = frontier.get();
                if (current == goal) break;

                for (Location next : graph.neighbours(current)) {
                    double new_cost = cost_so_far[current] +
                                      graph.cost(current.next);
                    if (cost_so_far.find(next) == cost_so_far.end() ||
                        new_cost < cost_so_far[next]) {
                        cost_so_far[next] = new_cost;
                        double priority = new_cost + heuristic(next, goal);
                        frontier.put(next, priority);
                        came_from[next] = current;
                    }
                }
            }
        }
    }; // namespace Pathfinder
} // namespace advent_of_code
