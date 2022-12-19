#include <vector>
#include <string>
#include <format>
#include <numeric>

#include <day-02.h>

namespace day_02
{
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

  using Game = std::vector<Round>;

  Game initialize(const std::vector<std::string>& input_data)
  {
    Game game;

    for (auto& line : input_data) {
      char theirs{ line[0] };
      char ours{ line[2] };
      game.push_back(Round{ make_throwable(theirs), make_throwable(ours) });
    }

    return game;
  }

  std::string answer_a(const std::vector<std::string>& input_data)
  {
    Game game = initialize(input_data);

    // what's my score

    auto sum = std::accumulate(game.begin(), game.end(), uint32_t{ 0 },
      [](uint32_t current, const Round& round) {
        return current + round.score_us();
      }
    );

    return std::format("{}", sum);
  }

  std::string answer_b(const std::vector<std::string>& input_data)
  {
    return "PENDING";
  }
}
