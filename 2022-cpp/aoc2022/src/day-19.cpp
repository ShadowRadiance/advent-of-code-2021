#include <days.h>

#include <algorithm>
#include <array>
#include <functional>
#include <iomanip>
#include <iostream>
#include <iterator>
#include <numeric>
#include <optional>
#include <queue>
#include <regex>

namespace day_19
{
    using std::array;
    using std::optional;
    using std::ostream;
    using std::queue;
    using std::regex;
    using std::smatch;
    using std::string;
    using std::vector;

    enum class Resource { ore, clay, obsidian, geode };
    enum class Bot { ore, clay, obsidian, geode };
    const array<Resource, 4> ResourceTypes = { Resource::ore, Resource::clay, Resource::obsidian, Resource::geode };
    const array<Bot, 4> ResourceBotTypes = { Bot::ore, Bot::clay, Bot::obsidian, Bot::geode };
    using Resources = array<int, ResourceTypes.size()>;
    Resources operator+(Resources const& lhs, Resources const& rhs);
    Resources operator-(Resources const& lhs, Resources const& rhs);
    using Bots = array<int, ResourceBotTypes.size()>;
    Resources operator*(Bots const& bots, int minutes);
    using BotCosts = array<Resources, ResourceBotTypes.size()>;
    string name(Resource res);
    string botName(Bot bot);

    class Blueprint
    {
    public:
        Blueprint(int id, array<Resources, ResourceTypes.size()> botCosts);
        int id() const;
        Resources botCost(Bot bot) const;
        friend ostream& operator<<(ostream& os, Blueprint const& blueprint);
    private:
        int id_ = 0;
        BotCosts botCosts_;
    };

    Blueprint parseBlueprint(string const& str);

    class Blueprints
    {
    public:
        Blueprints(vector<string> const& data);
        auto size();
        auto begin();
        auto end();
        auto begin() const;
        auto end() const;
        void applyHungryElephants();
    private:
        vector<Blueprint> blueprints_;
    };

    class Factory
    {
        class State
        {
        private:
            int minutesElapsed_;
            Resources resources_;
            Bots bots_;
        public:
            State();
            State(int, Resources const&, Bots const&);

            int minutesElapsed() const;

            int ore() const;
            int clay() const;
            int obsidian() const;
            int geodes() const;
            Resources resources() const;
            int resources(Resource type) const;

            int oreBots() const;
            int clayBots() const;
            int obsidianBots() const;
            int geodeBots() const;
            Bots bots() const;
            int bots(Bot type) const;

            State wait(int minutes) const;
            State buildBot(Bot type, Resources const& botCost) const;
        };
    public:
        Factory(Blueprint const& blueprint);
        int maxGeodes() const;
        void determineMaxGeodes();
        int qualityLevel() const;
        friend ostream& operator<<(ostream& os, Factory const& factory);
        static int setMaxExecutionTime(int minutes);
        static int maxTime();
    private:
        const Blueprint& blueprint_;
        Bots botLimits_;
        mutable int maxGeodes_;
        static int maxTime_;
    };

    using Factories = vector<Factory>;

    ostream& operator<<(ostream& os, Factories const& factories);
    ostream& operator<<(ostream& os, Blueprints const& blueprints);
    ostream& operator<<(ostream& os, Resources const& cost);

#pragma region HELPER FUNCTIONS
    string name(Resource res)
    {
        switch (res) {
        case Resource::ore: return "Ore";
        case Resource::clay: return "Clay";
        case Resource::obsidian: return "Obsidian";
        case Resource::geode: return "Geode";
        default: return "";
        };
    }

    string botName(Bot bot)
    {
        return name(static_cast<Resource>(bot)) + "Bot";
    }

    ostream& operator<<(ostream& os, Resources const& cost)
    {
        for (int i : cost) { os << "." << i; }
        return os;
    }

    Resources operator*(Bots const& bots, int minutes)
    {
        Resources resources{};
        std::transform(bots.begin(), bots.end(), resources.begin(), [minutes](int b) { return b * minutes; });
        return resources;
    }

    Resources operator+(Resources const& lhs, Resources const& rhs)
    {
        Resources resources{};
        std::transform(lhs.begin(), lhs.end(), rhs.begin(), resources.begin(), std::plus<int>{});
        return resources;
    }

    Resources operator-(Resources const& lhs, Resources const& rhs)
    {
        Resources resources{};
        std::transform(lhs.begin(), lhs.end(), rhs.begin(), resources.begin(), std::minus<int>{});
        return resources;
    }
#pragma endregion

#pragma region BLUEPRINT IMPLEMENTATION
    Blueprint::Blueprint(int id, array<Resources, 4> botCosts) : id_(id), botCosts_(botCosts) {}

    int Blueprint::id() const
    {
        return id_;
    }

    Resources Blueprint::botCost(Bot bot) const
    {
        return botCosts_[static_cast<int>(bot)];
    }

    ostream& operator<<(ostream& os, Blueprint const& bp)
    {
        os << bp.id() << ": /";
        for (Bot const& bot : ResourceBotTypes) { os << botName(bot) << bp.botCost(bot) << "/"; }
        return os;
    }

    Blueprint parseBlueprint(string const& str)
    {
        const regex re(
            "Blueprint (\\d+): Each ore robot costs (\\d+) ore. Each clay robot costs (\\d+) ore. " \
            "Each obsidian robot costs (\\d+) ore and (\\d+) clay. Each geode robot costs (\\d+) ore and (\\d+) obsidian."
        );
        smatch match;
        std::regex_match(str, match, re);
        int id = std::stoi(match[1].str());
        array<array<int, 4>, 4> costs{};
        costs[static_cast<int>(Bot::ore)][static_cast<int>(Resource::ore)] = std::stoi(match[2].str());
        costs[static_cast<int>(Bot::clay)][static_cast<int>(Resource::ore)] = std::stoi(match[3].str());
        costs[static_cast<int>(Bot::obsidian)][static_cast<int>(Resource::ore)] = std::stoi(match[4].str());
        costs[static_cast<int>(Bot::obsidian)][static_cast<int>(Resource::clay)] = std::stoi(match[5].str());
        costs[static_cast<int>(Bot::geode)][static_cast<int>(Resource::ore)] = std::stoi(match[6].str());
        costs[static_cast<int>(Bot::geode)][static_cast<int>(Resource::obsidian)] = std::stoi(match[7].str());

        return { id, costs };
    }
#pragma endregion

#pragma region BLUEPRINTS IMPLEMENTATION
    Blueprints::Blueprints(vector<string> const& data)
    {
        std::transform(data.begin(), data.end(),
                       std::back_inserter(blueprints_),
                       parseBlueprint);
    }

    auto Blueprints::size() { return blueprints_.size(); }

    auto Blueprints::begin() { return blueprints_.begin(); }

    auto Blueprints::end() { return blueprints_.end(); }

    auto Blueprints::begin() const { return blueprints_.begin(); }

    auto Blueprints::end() const { return blueprints_.end(); }

    void Blueprints::applyHungryElephants()
    {
        if (blueprints_.size() > 3) {
            blueprints_.erase(blueprints_.begin()+3, blueprints_.end());
        }
    }

    ostream& operator<<(ostream& os, Blueprints const& blueprints)
    {
        os << "BLUEPRINTS:" << "\n";
        for (Blueprint const& bp : blueprints) os << bp << "\n";
        return os;
    }
#pragma endregion

#pragma region FACTORY IMPLEMENTATION
    int Factory::maxTime_ = 0;

    Factory::Factory(Blueprint const& blueprint)
        : blueprint_(blueprint)
        , botLimits_()
        , maxGeodes_(0)
    {
        botLimits_ = {
            std::max({
                blueprint_.botCost(Bot::ore)[static_cast<int>(Resource::ore)],
                blueprint_.botCost(Bot::clay)[static_cast<int>(Resource::ore)],
                blueprint_.botCost(Bot::obsidian)[static_cast<int>(Resource::ore)],
                blueprint_.botCost(Bot::geode)[static_cast<int>(Resource::ore)],
            }),
            blueprint_.botCost(Bot::obsidian)[static_cast<int>(Resource::clay)],
            blueprint_.botCost(Bot::geode)[static_cast<int>(Resource::obsidian)],
            0
        };
    }

    int Factory::maxGeodes() const { return maxGeodes_; }

    void Factory::determineMaxGeodes()
    {
        using MaybeState = optional<State>;
        auto waitTime = [&](Resources botCost, State current) -> int {
            int wait = 0;
            for (Resource resource : ResourceTypes) {
                int resourceCost = botCost[static_cast<int>(resource)];
                if (resourceCost == 0) continue;
                
                int resourceCurr = current.resources(resource);
                if (resourceCost <= resourceCurr) continue;

                // resourceCost > resourceCurr so we have to wait
                int botsGeneratingResource = current.bots(static_cast<Bot>(resource));
                if (botsGeneratingResource == 0) return Factory::maxTime() + 1;
                
                int thisWait = (resourceCost - resourceCurr) / botsGeneratingResource;
                int thisRemainder = (resourceCost - resourceCurr) % botsGeneratingResource;
                if (thisRemainder > 0) thisWait += 1;

                wait = std::max(wait, thisWait);
            }
            return wait;
        };
        auto enoughBots = [&](Bot bot, State current) -> bool {
            int limit = botLimits_[static_cast<int>(bot)];
            if (limit == 0) return false;
            return current.bots(bot) >= limit;
        };
        auto noTime = [&](int wait, State current) -> bool {
            return wait + 1 + current.minutesElapsed() > maxTime_;
        };
        auto tryCreateBot = [&](Bot bot, State current) -> MaybeState {
            if (enoughBots(bot, current)) return {};
            
            int wait = waitTime(blueprint_.botCost(bot), current);
            if (noTime(wait, current)) return {};

            State newState = current.wait(wait + 1);
            newState = newState.buildBot(bot, blueprint_.botCost(bot));
            return newState;
        };

        queue<State> q;
        q.push(State(0, { 0,0,0,0 }, { 1,0,0,0 }));

        while (!q.empty()) {
            State current = q.front(); q.pop();

            // figure out max for just waiting for current to finish
            maxGeodes_ = std::max(
                maxGeodes_,
                current.geodes() + ((maxTime_ - current.minutesElapsed()) * current.geodeBots())
            );

            for (Bot bot : ResourceBotTypes) {
                MaybeState maybeState = tryCreateBot(bot, current);
                if (!maybeState.has_value()) continue;

                // If we theoretically only built geode bots every turn, 
                // and we still wouldn’t beat the current maximum, 
                // don’t push the state the the queue, but skip to the next item.
                // THIS MAKES AN ELEPHANTINE DIFFERENCE
                State newState = maybeState.value();
                int minutesRemaining = maxTime_ - newState.minutesElapsed();
                // 1+2+3+..n == n(n-1)/2;
                int maxPotentialNewGeodes =
                    (minutesRemaining * (minutesRemaining - 1)) / 2
                    + minutesRemaining * newState.bots(Bot::geode);
                if (newState.geodes() + maxPotentialNewGeodes < maxGeodes_) continue;

                q.push(maybeState.value());
            }
        }
    }

    int Factory::qualityLevel() const
    {
        return maxGeodes_ * blueprint_.id();
    }

    int Factory::maxTime()
    {
        return Factory::maxTime_;
    }

    int Factory::setMaxExecutionTime(int minutes)
    {
        int original = Factory::maxTime_;
        Factory::maxTime_ = minutes;
        return original;
    }

    ostream& operator<<(ostream& os, Factory const& factory)
    {
        os
            << "\tBlueprint: " << factory.blueprint_ << "\n"
            << "\tBotLimits: " << factory.botLimits_ << "\n"
            << "\tMaxGeodes: " << factory.maxGeodes_ << "\n"
            << "\n";
        return os;
    }
#pragma endregion

#pragma region FACTORY::STATE IMPLEMENTATION
    Factory::State::State()
        : resources_(), bots_(), minutesElapsed_(0)
    {}
    Factory::State::State(int minutes, Resources const& resources, Bots const& bots)
        : resources_(resources), bots_(bots), minutesElapsed_(minutes)
    {}

    int Factory::State::minutesElapsed() const { return minutesElapsed_; }
    int Factory::State::ore() const { return resources(Resource::ore); }
    int Factory::State::clay() const { return resources(Resource::clay); }
    int Factory::State::obsidian() const { return resources(Resource::obsidian); }
    int Factory::State::geodes() const { return resources(Resource::geode); }
    int Factory::State::resources(Resource type) const { return resources_[static_cast<int>(type)]; }
    Resources Factory::State::resources() const { return resources_; }
    int Factory::State::oreBots() const { return bots(Bot::ore); }
    int Factory::State::clayBots() const { return bots(Bot::clay); }
    int Factory::State::obsidianBots() const { return bots(Bot::obsidian); }
    int Factory::State::geodeBots() const { return bots(Bot::geode); }
    int Factory::State::bots(Bot type) const { return bots_[static_cast<int>(type)]; }
    Bots Factory::State::bots() const { return bots_; }

    Factory::State Factory::State::wait(int minutes) const
    {
        return Factory::State(
            minutesElapsed_ + minutes,
            resources_ + bots_ * minutes,
            bots_
        );
    }

    Factory::State Factory::State::buildBot(Bot type, Resources const& botCost) const
    {
        Bots newBots = bots_;
        newBots[static_cast<int>(type)] += 1;
        Resources newResources = resources_ - botCost;

        return State(minutesElapsed_, newResources, newBots);
    }
#pragma endregion

#pragma region FACTORIES IMPLEMENTATION
    ostream& operator<<(ostream& os, Factories const& factories)
    {
        os << "FACTORIES:\n";
        for (auto& factory : factories) os << factory << "\n";
        return os;
    }
#pragma endregion

    string answer_a(vector<string> const& input_data)
    {
        Factory::setMaxExecutionTime(24);
        Blueprints blueprints(input_data);

        Factories factories(blueprints.begin(), blueprints.end());
        //std::cout << factories << "\n";

        std::for_each(factories.begin(), factories.end(), std::mem_fn(&Factory::determineMaxGeodes));
        //std::cout << factories << "\n";

        return std::to_string(std::transform_reduce(
            factories.begin(), factories.end(),
            0, std::plus<int>{}, std::mem_fn(&Factory::qualityLevel)));
    }

    string answer_b(vector<string> const& input_data)
    {
        Factory::setMaxExecutionTime(32);
        Blueprints blueprints(input_data);
        blueprints.applyHungryElephants();

        Factories factories(blueprints.begin(), blueprints.end());
        std::cout << factories << "\n";

        std::for_each(factories.begin(), factories.end(), std::mem_fn(&Factory::determineMaxGeodes));
        std::cout << factories << "\n";

        return std::to_string(std::transform_reduce(
            factories.begin(), factories.end(),
            1, std::multiplies<int>{}, std::mem_fn(&Factory::maxGeodes)));
    }
}
