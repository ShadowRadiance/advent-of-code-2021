#include <format>
#include <numeric>

#include <days.h>

namespace day_02
{
  using std::string;
  using std::vector;
  using std::to_string;

  enum class Throwable {
    rock,                                                    // r-s 0-2   r -2  1                       r-p 0-1  p -1     2
    paper,                                                   // p-r 1-0   p +1  4 (1)                   p-s 1-2  s -1     2
    scissors                                                 // s-p 2-1   s +1  4 (1)                   s-r 2-0  r + 2     5 (mod3 =2)
  };

  Throwable make_throwable(char c) {
    switch (c) {
    case 'A': return Throwable::rock;
    case 'B': return Throwable::paper;
    case 'C': return Throwable::scissors;
    case 'X': return Throwable::rock;
    case 'Y': return Throwable::paper;
    case 'Z': return Throwable::scissors;
    default: throw "Do not be so damned daft";
    }
  }

  Throwable choose_throwable_for_outcome(Throwable theirs, char outcome) {
    switch (outcome) {
    case 'X': // lose
      return Throwable((uint32_t(theirs) + 2) % 3);
    case 'Y': // draw
      return theirs;
    case 'Z': // win
      return Throwable((uint32_t(theirs) + 1) % 3);
    default: throw "Do not be so damned daft";
    }
  }

  struct Round {
    Throwable them;
    Throwable us;

    uint32_t score_us() const {
      return score_thrown(us) + winlosedraw(us, them);
    }

    uint32_t score_them() const {
      return score_thrown(them) + winlosedraw(them, us);
    }

    uint32_t score_thrown(Throwable t) const {
      return uint32_t(t) + 1;
    }

    static constexpr uint32_t LOSS = 0;
    static constexpr uint32_t TIE = 3;
    static constexpr uint32_t WIN = 6;

    uint32_t winlosedraw(Throwable first, Throwable second) const {
      uint32_t difference = uint32_t(first) - uint32_t(second);

      if ((3 + difference) % 3 == 1) return WIN;
      if ((3 + difference) % 3 == 2) return LOSS;

      return TIE;
    }
  };

  using Game = vector<Round>;

  Game initialize(const vector<string>& input_data)
  {
    Game game;

    for (auto& line : input_data) {
      char theirs{ line[0] };
      char ours{ line[2] };
      game.push_back(Round{ make_throwable(theirs), make_throwable(ours) });
    }

    return game;
  }

  Game initialize_correctly(const vector<string>& input_data)
  {
    Game game;

    for (auto& line : input_data) {
      Throwable theirs = make_throwable(line[0]);
      char winlosedraw{ line[2] };
      game.push_back(Round{ theirs, choose_throwable_for_outcome(theirs, winlosedraw) });
    }

    return game;
  }

  string answer_a(const vector<string>& input_data)
  {
    Game game = initialize(input_data);

    // what's my score

    auto sum = accumulate(game.begin(), game.end(), uint32_t{ 0 },
      [](uint32_t current, const Round& round) {
        return current + round.score_us();
      }
    );

    return to_string(sum);
  }

  string answer_b(const vector<string>& input_data)
  {
    Game game = initialize_correctly(input_data);

    // what's my score

    auto sum = accumulate(game.begin(), game.end(), uint32_t{ 0 },
      [](uint32_t current, const Round& round) {
        return current + round.score_us();
      }
    );

    return to_string(sum);
  }
}
