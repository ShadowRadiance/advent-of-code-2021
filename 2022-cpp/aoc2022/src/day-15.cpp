#include <days.h>
#include <regex>
#include <cassert>
#include <unordered_map>
#include <variant>
#include <optional>
#include <numeric>
#include <iostream>
#include <functional>

namespace day_15
{
	using std::string;
	using std::vector;

	struct Vector2D
	{
		double x;
		double y;

		static const Vector2D up() { return { 0.0,-1.0 }; }
		static const Vector2D dn() { return { 0.0,+1.0 }; }
		static const Vector2D lt() { return { -1.0,0.0 }; }
		static const Vector2D rt() { return { +1.0,0.0 }; }

		bool operator<(const Vector2D& rhs) const { return x < rhs.x || (x == rhs.x && y < rhs.y); }
		bool operator==(const Vector2D& rhs) const { return x == rhs.x && y == rhs.y; }
		Vector2D operator+(const Vector2D& rhs) const { return { x + rhs.x, y + rhs.y }; }
		Vector2D operator-(const Vector2D& rhs) const { return { x - rhs.x, y - rhs.y }; }
		Vector2D unit() const { return { x / magnitude(), y / magnitude() }; }
		double manhattanDistanceFrom(const Vector2D& rhs) const {
			Vector2D pos{ *this - rhs };
			return abs(pos.x) + abs(pos.y);
		}
		double magnitude() const {
			return sqrt(x * x + y * y);
		}
	};

	struct Vector2D_Hash {
		size_t operator()(const Vector2D& vec2d) const noexcept {
			size_t x = std::hash<int>()(vec2d.x);
			size_t y = std::hash<int>()(vec2d.y);
			return x ^ (y << 1);
		}
	};

	struct Sensor {
		Vector2D location;
		Vector2D closestBeacon;
		int distanceToClosestBeacon() const { return location.manhattanDistanceFrom(closestBeacon); }
		int reach() const { return distanceToClosestBeacon(); }

		bool operator==(const Sensor& rhs) const { return location == rhs.location; }
	};

	class Parser
	{
	public:
		vector<Sensor> parseSensors(const vector<string>& data) const
		{
			vector<Sensor> sensors;
			transform(data.begin(), data.end(), back_inserter(sensors), &parseSensor);
			return sensors;
		}
	private:
		static const std::regex re;

		static Sensor parseSensor(const string& str)
		{
			// "Sensor at x=2, y=18: closest beacon is at x=-2, y=15"
			std::smatch matches;
			bool matched = std::regex_match(str, matches, re);
			assert(matched);
			assert(matches.size() == 5);
			assert(matches[0].str() == str);
			Vector2D sensorLocation{ std::stod(matches[1].str()), std::stod(matches[2].str()) };
			Vector2D beaconLocation{ std::stod(matches[3].str()), std::stod(matches[4].str()) };

			return { sensorLocation, beaconLocation };
		}
	};
	const std::regex Parser::re{ "Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)" };

	string answer_a_impl_1(vector<Sensor> sensors, int row)
	{
		// 1. Sensors and beacons always exist at integer coordinates. 
		// 2. Each sensor knows its own position.
		// 3. Each sensor can determine the position of the closest beacon precisely(by Manhattan distance).
		// 4. None of the detected beacons seem to be producing the distress signal
		// 5. Return the number of positions in row 10 (test) or row 2_000_000 (live) that cannot contain a beacon.

		auto leftReach = [](auto& lhs, auto& rhs) { return (lhs.location.x - lhs.reach()) < (rhs.location.x - lhs.reach()); };
		auto rightReach = [](auto& lhs, auto& rhs) { return (lhs.location.x + lhs.reach()) < (rhs.location.x + lhs.reach()); };
		auto furthestLeftReachingSensorIter = std::min_element(sensors.begin(), sensors.end(), leftReach);
		auto furthestRightReachingSensorIter = std::max_element(sensors.begin(), sensors.end(), rightReach);

		int lowX = furthestLeftReachingSensorIter->location.x - furthestLeftReachingSensorIter->reach();
		int highX = furthestRightReachingSensorIter->location.x + furthestRightReachingSensorIter->reach();

		std::cout << "Reading from " << lowX << " to " << highX << std::endl;

		int reachable = 0;
		for (int x = lowX; x <= highX; ++x)
		{
			Vector2D location{ double(x), double(row) };
			auto withinRange = [&location](const Sensor& sensor) -> bool { return sensor.location.manhattanDistanceFrom(location) <= sensor.reach(); };
			if (std::any_of(sensors.begin(), sensors.end(), withinRange)) ++reachable;
		}

		vector<Vector2D> beaconLocations;
		std::transform(sensors.begin(), sensors.end(), back_inserter(beaconLocations), [](auto& sensor) { return sensor.closestBeacon; });
		std::sort(beaconLocations.begin(), beaconLocations.end());
		beaconLocations.erase(
			std::unique(beaconLocations.begin(), beaconLocations.end()),
			beaconLocations.end()
		);
		int beaconsOnRow = std::count_if(beaconLocations.begin(), beaconLocations.end(), [row](auto& location) { return location.y == row; });

		int reachableAndNotABeacon = reachable - beaconsOnRow;

		return std::to_string(reachableAndNotABeacon);
	}

	class Range
	{
		int start, end;
	public:
		Range(int start, int end) : start(start), end(end) {}

		int length() const { return end - start; }

		bool covers(const Range& other) const
		{
			return other.startsWithin(*this) && other.endsWithin(*this);
		}

		bool overlaps(const Range& other) const
		{
			return other.startsWithin(*this) || other.endsWithin(*this);
		}

		bool startsWithin(const Range& other) const
		{
			return start >= other.start && start < other.end;
		}

		bool endsWithin(const Range& other) const
		{
			return end > other.start && end <= other.end;
		}

		friend bool operator<(const Range& lhs, const Range& rhs) {
			if (lhs.start == rhs.start) return lhs.end < rhs.end;
			return lhs.start < rhs.start;
		}

		friend Range range_covering(const Range& lhs, const Range& rhs) {
			return
			{
				std::min(lhs.start, rhs.start),
				std::max(lhs.end, rhs.end)
			};
		}
	};

	struct DisjointRange
	{
		vector<Range> subRanges;

		void reduce()
		{
			std::sort(subRanges.begin(), subRanges.end());
			vector<Range> stack;
			for (auto& r : subRanges)
			{
				// put the first one on the stack
				if (stack.empty()) { stack.push_back(r); continue; }

				Range top = stack.back();

				// if the one on the stack completely covers this one, skip it
				if (top.covers(r)) {
					continue;
				}

				// if this one completely covers the one on the stack, swap em
				if (r.covers(top)) {
					stack.pop_back();
					stack.push_back(r);
					continue;
				}

				// if they overlap, merge into one range and replace
				if (r.overlaps(top)) {
					stack.pop_back();
					stack.push_back(range_covering(top, r));
					continue;
				}

				// no overlap
				stack.push_back(r);
			}
			subRanges = stack;
		}
	};

	auto rangeOnRow(int row, const Sensor& sensor) -> Range
	{
		int distanceFromRow = abs(row - int(sensor.location.y));
		if (distanceFromRow > sensor.reach())
		{
			return {
				int(sensor.location.x),
				int(sensor.location.x), // one past the end
			};
		}
		else
		{
			return {
				int(sensor.location.x - (sensor.reach() - distanceFromRow)),
				int(sensor.location.x + (sensor.reach() - distanceFromRow) + 1), // one past the end
			};
		}
	};

	string answer_a_impl_2(vector<Sensor> sensors, int row)
	{
		auto rangeOnRowN = std::bind(rangeOnRow, row, std::placeholders::_1);
		vector<Range> ranges;
		std::transform(sensors.begin(), sensors.end(), back_inserter(ranges), rangeOnRowN);
		DisjointRange disjointRange{ ranges };
		disjointRange.reduce();
		int covered = std::transform_reduce(
			disjointRange.subRanges.begin(), disjointRange.subRanges.end(),
			0,
			std::plus<>{},
			[](auto& range) { return range.length(); });

		vector<Vector2D> beaconLocations;
		std::transform(sensors.begin(), sensors.end(), back_inserter(beaconLocations), [](auto& sensor) { return sensor.closestBeacon; });
		std::sort(beaconLocations.begin(), beaconLocations.end());
		beaconLocations.erase(
			std::unique(beaconLocations.begin(), beaconLocations.end()),
			beaconLocations.end()
		);
		int beaconsOnRow = std::count_if(beaconLocations.begin(), beaconLocations.end(), [row](auto& location) { return location.y == row; });

		return std::to_string(covered - beaconsOnRow);
	}

	string answer_a_impl(vector<Sensor> sensors, int row) {
		return answer_a_impl_2(sensors, row);
	}

	string answer_a_test(const vector<string>& input_data)
	{
		Parser parser;
		vector<Sensor> sensors = parser.parseSensors(input_data);
		return answer_a_impl(sensors, 10);
	};

	string answer_a(const vector<string>& input_data)
	{
		Parser parser;
		vector<Sensor> sensors = parser.parseSensors(input_data);
		return answer_a_impl(sensors, 2000000);
	}

	class BeaconLocator
	{
	public:
		BeaconLocator(const vector<Sensor>& sensors, int min, int max) 
			: sensors(sensors), min(min), max(max)
		{}

		std::optional<Vector2D> locateBeacon()
		{
			// GIVEN: there is only one location within min,min to max,max that is possible
			// GIVEN: the beacon must be one unit away from the covered reach of the sensors
			// WHEN: we check each sensors "border" (1 px further than its reach)
			//  AND: if we find a location that none of the remaining sensors reach
			// THEN: we have found our beacon

			// 14 sensors in test, 25 in live

			auto itSensorCurrent = sensors.begin();
			auto itEnd = sensors.end();

			while (itSensorCurrent != itEnd)
			{
				std::optional<Vector2D> result = checkSensorBoundary(itSensorCurrent);
				if (result.has_value()) 
					return result.value();
				++itSensorCurrent;
			}
			return {};
		}
	private:
		std::optional<Vector2D>checkSensorBoundary(auto current)
		{
			vector<Vector2D> borderLocations = generateBorderLocations(*current);

			for (auto& location : borderLocations) {
				if (location.x < min || location.x > max) continue;
				if (location.y < min || location.y > max) continue;

				if (!sensorsCanSeeLocation(*current, location)) return location;
			}

			return {};
		}

		bool sensorsCanSeeLocation(const Sensor& activeSensor, Vector2D location)
		{
			for(auto& sensor : sensors)
			{
				if (sensor == activeSensor) continue;
				if (sensor.location.manhattanDistanceFrom(location) <= sensor.reach()) return true;
			}
			return false;
		}

		vector<Vector2D> generateBorderLocations(const Sensor& sensor) {
			vector<Vector2D> locations;

			Vector2D top{ sensor.location.x, sensor.location.y - sensor.reach() - 1 };
			Vector2D left{ sensor.location.x - sensor.reach() - 1, sensor.location.y };
			Vector2D right{ sensor.location.x + sensor.reach() + 1, sensor.location.y };
			Vector2D bottom{ sensor.location.x, sensor.location.y + sensor.reach() + 1 };

			Vector2D fromTopToLeft{ -1, +1 };
			Vector2D fromLeftToBottom{ +1, +1 };
			Vector2D fromBottomToRight{ +1, -1 };
			Vector2D fromRightToTop{ -1, -1 };

			Vector2D current = top;
			while (current != left) { locations.push_back(current); current = current + fromTopToLeft; }
			while (current != bottom) { locations.push_back(current); current = current + fromLeftToBottom; }
			while (current != right) { locations.push_back(current); current = current + fromBottomToRight; }
			while (current != top) { locations.push_back(current); current = current + fromRightToTop; }

			return locations;
		}

		const vector<Sensor>& sensors;
		int min;
		int max;
	};

	string answer_b_impl(const vector<Sensor>& sensors, int max)
	{
		BeaconLocator beaconLocator(sensors, 0, max);

		std::optional<Vector2D> distressBeaconLocation = beaconLocator.locateBeacon();

		if (!distressBeaconLocation.has_value())
			return "NOT FOUND :-(";

		int64_t frequency =
			int64_t(distressBeaconLocation.value().x) * 4000000ll
			+ int64_t(distressBeaconLocation.value().y);
		return std::to_string(frequency);
	}

	string answer_b_test(const vector<string>& input_data)
	{
		Parser parser;
		vector<Sensor> sensors = parser.parseSensors(input_data);
		return answer_b_impl(sensors, 20);
	}

	string answer_b(const vector<string>& input_data)
	{
		Parser parser;
		vector<Sensor> sensors = parser.parseSensors(input_data);
		return answer_b_impl(sensors, 4000000);
	}
}
