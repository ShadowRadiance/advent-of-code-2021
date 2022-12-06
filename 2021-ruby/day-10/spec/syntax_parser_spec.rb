require "./syntax_parser"

RSpec.describe "SyntaxParser" do
  subject { SyntaxParser.new(data) }

  let(:data) {
    [
      "[({(<(())[]>[[{[]{<()<>>",
      "[(()[<>])]({[<{<<[]>>(",
      "{([(<{}[<>[]}>{[]{[(<()>",
      "(((({<>}<{<{<>}{[]{[]{}",
      "[[<[([]))<([[{}[[()]]]",
      "[{[{({}]{}}([{[{{{}}([]",
      "{<[[]]>}<{[{[{[]{()[[[]",
      "[<(<(<(<{}))><([]([]()",
      "<{([([[(<>()){}]>(<<{{",
      "<{([{{}}[<[[[<>{}]]]>[]]",
    ]
  }

  it "catches the correct errors" do
    expect(subject.errors.count).to eq(5)
    expect(subject.errors.map(&:to_s)).to eq([
      "{([(<{}[<>[]}>{[]{[(<()> - Expected ], but found } instead.",
      "[[<[([]))<([[{}[[()]]] - Expected ], but found ) instead.",
      "[{[{({}]{}}([{[{{{}}([] - Expected ), but found ] instead.",
      "[<(<(<(<{}))><([]([]() - Expected >, but found ) instead.",
      "<{([([[(<>()){}]>(<<{{ - Expected ], but found > instead.",
    ])
  end
  
  it "calculates the correct error score" do
    expect(subject.total_syntax_error_score).to eq(26397)
  end

  it "calculates the correct autocompletes" do
    expect(subject.autocompletions).to eq(["}}]])})]", ")}>]})", "}}>}>))))", "]]}}]}]}>", "])}>"])
  end

  it "calculates the correct autocomplete scores" do
    expect(subject.autocomplete_scores).to eq([288957, 5566, 1480781, 995444, 294])
  end

  it "calculates the correct autocomplete score" do
    expect(subject.autocomplete_score).to eq(288957)
  end
end