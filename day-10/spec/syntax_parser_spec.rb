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
  
  it "calculates the correct score" do
    expect(subject.total_syntax_error_score).to eq(26397)
  end
end