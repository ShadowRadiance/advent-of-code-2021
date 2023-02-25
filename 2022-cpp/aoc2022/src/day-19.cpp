#include <days.h>
#include <cassert>
#include <algorithm>
#include <stack>
#include <optional>
#include <unordered_map>
#include <numeric>
#include <functional>
#include <array>
#include <iterator>
#include <queue>
#include <stdexcept>
#include <iostream>
#include <unordered_set>
#include <sstream>
#include <iomanip>
#include <array>
#include <regex>

namespace day_19
{
    using std::string;
    using std::vector;
    using std::array;

    enum class Resource { ore, clay, obsidian, geode };
    enum class Bot { ore, clay, obsidian, geode };
    const array<Resource, 4> ResourceTypes = { Resource::ore, Resource::clay, Resource::obsidian, Resource::geode };
    const array<Bot, 4> ResourceBotTypes = { Bot::ore, Bot::clay, Bot::obsidian, Bot::geode };
    using Resources = array<int, ResourceTypes.size()>;
    using Bots = array<int, ResourceBotTypes.size()>;
    using BotCosts = array<Resources, ResourceBotTypes.size()>;
    string name(Resource res);
    string botName(Bot bot);

    class Blueprint
    {
    public:
        Blueprint(int id, array<Resources, ResourceTypes.size()> botCosts);
        int id() const;
        Resources botCost(Bot bot) const;
        friend std::ostream& operator<<(std::ostream& os, Blueprint const& blueprint);
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
    private:
        vector<Blueprint> blueprints_;
    };

    class Factory
    {
    public:
        Factory(Blueprint const& blueprint);
        int maxGeodes() const;
        void determineMaxGeodes();
        int qualityLevel() const;
        friend std::ostream& operator<<(std::ostream& os, Factory const& factory);
    private:
        const Blueprint& blueprint_;
        Resources resources_;
        Bots bots_;
        Bots botLimits_;
        mutable int maxGeodes_;
    };

    using Factories = vector<Factory>;

    std::ostream& operator<<(std::ostream& os, Factories const& factories);
    std::ostream& operator<<(std::ostream& os, Blueprints const& blueprints);
    std::ostream& operator<<(std::ostream& os, Resources const& cost);

    // ----------------------------------------------------------------------------------------------------------------

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

    std::ostream& operator<<(std::ostream& os, Resources const& cost)
    {
        for (int i : cost) { os << "." << i; }
        return os;
    }

    // ----------------------------------------------------------------------------------------------------------------

    Blueprint::Blueprint(int id, array<Resources, 4> botCosts) : id_(id), botCosts_(botCosts) {}

    int Blueprint::id() const
    {
        return id_;
    }

    Resources Blueprint::botCost(Bot bot) const
    {
        return botCosts_[static_cast<int>(bot)];
    }

    std::ostream& operator<<(std::ostream& os, Blueprint const& bp)
    {
        os << bp.id() << ": /";
        for (Bot const& bot : ResourceBotTypes) { os << botName(bot) << bp.botCost(bot) << "/"; }
        return os;
    }

    Blueprint parseBlueprint(string const& str)
    {
        const std::regex re(
            "Blueprint (\\d+): Each ore robot costs (\\d+) ore. Each clay robot costs (\\d+) ore. " \
            "Each obsidian robot costs (\\d+) ore and (\\d+) clay. Each geode robot costs (\\d+) ore and (\\d+) obsidian."
        );
        std::smatch match;
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

    // ----------------------------------------------------------------------------------------------------------------

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

    std::ostream& operator<<(std::ostream& os, Blueprints const& blueprints)
    {
        os << "BLUEPRINTS:" << std::endl;
        for (Blueprint const& bp : blueprints) os << bp << std::endl;
        return os;
    }

    // ----------------------------------------------------------------------------------------------------------------

    Factory::Factory(Blueprint const& blueprint)
        : blueprint_(blueprint)
        , resources_()
        , bots_()
        , botLimits_()
        , maxGeodes_(0)
    {
        bots_[static_cast<int>(Bot::ore)] = 1;

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

    int Factory::maxGeodes() const
    {
        return maxGeodes_;
    }

    void Factory::determineMaxGeodes()
    {


    }

    int Factory::qualityLevel() const
    {
        return maxGeodes_ * blueprint_.id();
    }

    std::ostream& operator<<(std::ostream& os, Factory const& factory)
    {
        os
            << "\tBlueprint: " << factory.blueprint_ << "\n"
            << "\tResources: " << factory.resources_ << "\n"
            << "\tBots: " << factory.bots_ << "\n"
            << "\tBotLimits: " << factory.botLimits_ << "\n"
            << "\tMaxGeodes: " << factory.maxGeodes_ << "\n"
            << "\n";
        return os;
    }

    // ----------------------------------------------------------------------------------------------------------------

    std::ostream& operator<<(std::ostream& os, Factories const& factories)
    {
        os << "FACTORIES:\n";
        for (auto& factory : factories) os << factory << std::endl;
        return os;
    }

    // ----------------------------------------------------------------------------------------------------------------

    string answer_a(vector<string> const& input_data)
    {
        Blueprints blueprints(input_data);

        Factories factories(blueprints.begin(), blueprints.end());
        std::cout << factories << std::endl;

        std::for_each(factories.begin(), factories.end(), std::mem_fn(&Factory::determineMaxGeodes));

        return std::to_string(std::transform_reduce(
            factories.begin(), factories.end(),
            0, std::plus<int>{}, std::mem_fn(&Factory::qualityLevel)));
    }

    string answer_b(vector<string> const& input_data)
    {
        return "PENDING";

        //Blueprints blueprints(input_data);
        //std::cout << blueprints << std::endl;

        //Factories factories(blueprints.begin(), blueprints.end());
        //std::for_each(factories.begin(), factories.end(), std::mem_fn(Factory::determineMaxGeodes));

        //return std::to_string(std::transform_reduce(
        //    factories.begin(), factories.end(),
        //    0, std::plus<int>{}, std::mem_fn(Factory::qualityLevel)));
    }
}
