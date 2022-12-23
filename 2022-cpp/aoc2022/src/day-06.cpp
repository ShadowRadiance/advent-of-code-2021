#include <days.h>
#include <iostream>
#include <cassert>
#include <array>

namespace day_06
{
  using std::string;
  using std::vector;
  using std::array;
  using strings = vector<string>;

  using std::to_string;

  template<typename TCollection>
  class Window
  {
  public:
    typedef typename TCollection::const_iterator const_iterator;

    Window(const TCollection& collection, int windowSize)
      : collection_(collection)
      , windowSize_(windowSize)
      , start_(collection.begin())
    {
      assert(collection.begin() + windowSize <= collection.end());
    }

    TCollection view()
    {
      TCollection coll;
      std::copy(start_, start_ + windowSize_, back_inserter(coll));
      return coll;
    }

    void slide(int count)
    {
      assert(start_ + count + windowSize_ <= collection_.end());

      start_ += count;
    }

    bool hasMore()
    {
      return !atEnd();
    }

    bool atEnd()
    {
      return start_ + windowSize_ == collection_.end();
    }

    int offset() {
      return start_ - collection_.begin();
    }

  private:
    const TCollection& collection_;
    const int windowSize_;
    const_iterator start_;
  };

  template<int WindowSize>
  bool allDifferent(string s) {
    array<char, WindowSize> chars;
    copy(s.begin(), s.end(), chars.begin());

    // are "bvwb" (or "vncz") all different characters
    for (size_t idx{ 0 }; idx < WindowSize; idx++) {
      char checkChar = s[idx];
      for (size_t subIdx{ idx + 1 }; subIdx < WindowSize; subIdx++) {
        if (checkChar == s[subIdx]) {
          return false;
        }
      }
    }

    return true;
  }

  string answer_a(const strings& input_data)
  {
    const int WINDOW_SIZE = 4;
    Window window(input_data[0], WINDOW_SIZE);

    string s;
    while (!window.atEnd()) {
      if (allDifferent<WINDOW_SIZE>(window.view())) {
        return to_string(window.offset() + WINDOW_SIZE);
      }
      window.slide(1);
    }

    return "NO FOUR-CHARACTER SUBSTRING FOUND WITH FOUR UNIQUE CHARACTERS "
      "AFTER " + to_string(window.offset()) + " SLIDES";
  }

  string answer_b(const strings& input_data)
  {
    const int WINDOW_SIZE = 14;
    Window window(input_data[0], WINDOW_SIZE);

    string s;
    while (!window.atEnd()) {
      if (allDifferent<WINDOW_SIZE>(window.view())) {
        return to_string(window.offset() + WINDOW_SIZE);
      }
      window.slide(1);
    }

    return "NO FOUR-CHARACTER SUBSTRING FOUND WITH FOUR UNIQUE CHARACTERS "
      "AFTER " + to_string(window.offset()) + " SLIDES";
  }
}
