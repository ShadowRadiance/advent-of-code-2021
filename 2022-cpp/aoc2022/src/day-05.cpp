#include <days.h>

#include <ranges>
#include <algorithm>
#include <cassert>
#include <sstream>

namespace day_05
{
  using std::string;
  using std::vector;
  using std::stringstream;

  // REFERENCE EXAMPLE
  // 0: "    [D]    ",
  // 1: "[N] [C]    ",
  // 2: "[Z] [M] [P]",
  // 3: " 1   2   3 ",
  // 4: " ",
  // 5: "move 1 from 2 to 1",
  // 6: "move 3 from 1 to 3",
  // 7: "move 2 from 2 to 1",
  // 8: "move 1 from 1 to 2",

  struct Command
  {
    int count;
    int fromStack;
    int toStack;
  };

  class Stack
  {
  public:
    bool empty() const { return data_.empty(); }
    void push(char c) { data_.push_back(c); }
    char pop() { char c = top(); data_.pop_back(); return c; }
    char top() const { assert(!empty()); return data_.back(); }
  private:
    vector<char> data_;
  };

  using Stacks = vector<Stack>;
  using strings = vector<string>;
  using Commands = vector<Command>;

  int find_blank_line_index(const strings& input_data)
  {
    // find first non-space character - if found, line is not blank
    auto isBlank = [](const string& s) -> bool {
      auto firstNonSpace = find_if_not(s.begin(), s.end(), isspace);
      return firstNonSpace == s.end();
    };
    
    auto firstBlank = find_if(input_data.begin(), input_data.end(), isBlank);
    return (firstBlank == input_data.end()) ? -1 : firstBlank - input_data.begin();
  }

  Stacks initialize_stacks(const string& count_line)
  {
    //    read STACKCOUNT from last entry of COUNTLINE        // Example => STACKCOUNT = 3
    //    build an array of STACKCOUNT stacks (STACKS)        // Example => STACKS is a vector<stack> with 3 empty stacks
    size_t number{ 0 };
    stringstream ss(count_line);
    while (ss.good()) {
      ss >> number;
    }
    return Stacks{ number };
  }

  void populate_stacks(Stacks& stacks, const strings& stack_image)
  {
    //    for LINE in STACKIMAGE.REVERSED                     // Example => Read LINE 2 then LINE 1 then line 0
    //      for IDX from 0 to STACKCOUNT-1                    // Example => 0,1,2
    //        read CHAR from LINE[4 * IDX + 1)                // Example => Z,M,P then N,C,_ then _,D,_
    //        if CHAR a letter, push CHAR onto STACKS[i]      // Example => push "Z" onto STACKS[0], push "M" onto STACKS[1], etc
    //      next IDX                                          //
    //    next LINE                                           //
    std::for_each(stack_image.rbegin(), stack_image.rend(), 
      [& stacks](const string& line) mutable {
        for (size_t stackIdx = 0; stackIdx < stacks.size(); stackIdx++) {
          int lineIdx = 4 * stackIdx + 1;
          char c = line[lineIdx];
          if (c != ' ') {
            stacks[stackIdx].push(c);
          }
        }
      }
    );
  }

  Commands initialize_commands(const strings& instructions)
  {
    //  create INSTLIST                                       // Example => INSTLIST is a vector<Instruction> initially empty
    //  for LINE in INSTRUCTIONS                              // Example => "move 1 from 2 to 1"
    //    extract COUNT, FROM_STACK, TO_STACK into new INST   // Example => COUNT=1, FROM_STACK=2, TO_STACK=1
    //    push INST into INSTLIST                             // Example => INSTLIST.push_back(Instruction{1,2,1})
    //  next LINE                                             // 
    // -------------------------------------------------------//
    Commands commands;

    for (const string& instruction : instructions) {
      stringstream ss(instruction);
      string dummy;
      int count, fromStack, toStack;
      ss >> dummy >> count >> dummy >> fromStack >> dummy >> toStack;
      commands.push_back({ count, fromStack, toStack });
    }

    return commands;
  }

  class StackMover
  {
  public:
    StackMover(Stacks& stacks) : stacks_(stacks) {}
    virtual void moveCrates(const Command& command) = 0;
  protected:
    Stacks& stacks_;
  };
  class StackMover9000 : public StackMover
  {
  public:
    StackMover9000(Stacks& stacks) : StackMover(stacks) {}

    void moveCrates(const Command& command) override {
      //  for INST in INSTLIST                                  // Example => INST = Instruction{times:1,fromStack:2,toStack:1}
      //    for I from 0 to INST.COUNT-1                        // Example => 0
      //      pop from the INST.FROM_STACK stack into CHAR      // Example => CHAR = STACKS[INST.fromStack-1].pop() = STACKS[1].pop() = "D"
      //      push CHAR onto the INST.TO_STACK stack            // Example => STACKS[INST.toStack-1].push(CHAR)     = STACKS[0].push("D") => STACKS[0] = {Z,N,D}
      //    next I                                              // 
      //  next INST                                             // 
      // -------------------------------------------------------//
      for (size_t counter = 0; counter < command.count; counter++) {
        char c = stacks_[command.fromStack - 1].pop();
        stacks_[command.toStack - 1].push(c);
      }
    }
  };
  class StackMover9001 : public StackMover
  {
  public:
    StackMover9001(Stacks& stacks) : StackMover(stacks) {}
    void moveCrates(const Command& command) override {
      Stack tempStack;
      for (size_t counter = 0; counter < command.count; counter++)
      {
        tempStack.push(stacks_[command.fromStack - 1].pop());
      }
      for (size_t counter = 0; counter < command.count; counter++)
      {
        stacks_[command.toStack - 1].push(tempStack.pop());
      }
    }
  };

  void process_commands(const Commands& commands, StackMover& stackMover)
  {
    for (const Command& command : commands) {
      stackMover.moveCrates(command);
    }
  }

  string answer_a(const strings& input_data)
  {
    // parse initial state                                    //
    //  find first empty line (BLANKLINE)                     // Example => BLANKLINE = 4
    //    COUNTLINE = BLANKLINE-1                             // Example => COUNTLINE = 3
    //    STACKIMAGE = INPUTDATA[0..BLANKLINE-2]              // Example => STACKIMAGE = INPUTDATA[0,1,2]
    //    INSTRUCTIONS = INPUTDATA[BLANKLINE+1..END]          // Example => INSTRUCTIONS = INPUTDATA[5,6,7,8]
    int blank_line = find_blank_line_index(input_data);
    int count_line = blank_line - 1;
    strings stack_image; copy_n(input_data.begin(), count_line, back_inserter(stack_image));
    strings instructions; copy(input_data.begin() + blank_line + 1, input_data.end(), back_inserter(instructions));

    Stacks stacks = initialize_stacks(input_data[count_line]);
    populate_stacks(stacks, stack_image);

    Commands commands = initialize_commands(instructions);
    StackMover9000 stackMover{ stacks };
    process_commands(commands, stackMover);

    // read final state                                       //
    //  pull each top of stacks                               //
    //    TOPS = STACKS.MAP(&:TOP)                            // Example => {C,M,Z}
    //    result = TOPS.join()                                // Example => "CMZ"
    // -------------------------------------------------------//
    string tops;
    transform(stacks.begin(), stacks.end(), back_inserter(tops), [](Stack& stack) { return stack.top(); });

    return tops;
  }

  string answer_b(const vector<string>& input_data)
  {
    // parse initial state                                    //
    //  find first empty line (BLANKLINE)                     // Example => BLANKLINE = 4
    //    COUNTLINE = BLANKLINE-1                             // Example => COUNTLINE = 3
    //    STACKIMAGE = INPUTDATA[0..BLANKLINE-2]              // Example => STACKIMAGE = INPUTDATA[0,1,2]
    //    INSTRUCTIONS = INPUTDATA[BLANKLINE+1..END]          // Example => INSTRUCTIONS = INPUTDATA[5,6,7,8]
    int blank_line = find_blank_line_index(input_data);
    int count_line = blank_line - 1;
    strings stack_image; copy_n(input_data.begin(), count_line, back_inserter(stack_image));
    strings instructions; copy(input_data.begin() + blank_line + 1, input_data.end(), back_inserter(instructions));

    Stacks stacks = initialize_stacks(input_data[count_line]);
    populate_stacks(stacks, stack_image);

    Commands commands = initialize_commands(instructions);
    StackMover9001 stackMover{ stacks };
    process_commands(commands, stackMover);

    // read final state                                       //
    //  pull each top of stacks                               //
    //    TOPS = STACKS.MAP(&:TOP)                            // Example => {C,M,Z}
    //    result = TOPS.join()                                // Example => "CMZ"
    // -------------------------------------------------------//
    string tops;
    transform(stacks.begin(), stacks.end(), back_inserter(tops), [](Stack& stack) { return stack.top(); });

    return tops;
  }
}
