class SyntaxParser
  attr_reader :lines
  def initialize(lines)
    @lines = lines
  end

  OPENING = [ "[", "(", "{", "<" ]

  CLOSING_MAP = {
    "[" => "]",
    "{" => "}",
    "(" => ")",
    "<" => ">",
  }.freeze

  SCORING_MAP = {
    "]" => 57,
    "}" => 1197,
    ")" => 3,
    ">" => 25137,
  }.freeze


  SymbolNestingError = Struct.new(:line, :expected, :actual) do
    def to_s
      "#{line} - Expected #{expected}, but found #{actual} instead."
    end
    def score
      SCORING_MAP[actual]
    end
  end

  def errors
    @errors ||= lines.map { |line| parse(line) }.compact
  end

  def error_scores
    errors.map(&:score)
  end

  def total_syntax_error_score
    error_scores.sum
  end

  # return nil if everything okay or SymbolNestingError otherwise
  def parse(line)
    working = ""
    line.chars.each do |char|
      if open?(char) 
        working += char
        next
      end
      
      # char is a closing-char ), ], or }
      expected = closing_of(working.chars.last)
      if expected == char
        working.chop!
      else
        return SymbolNestingError.new(line, expected, char) 
      end
    end
    
    nil
  end

  def open?(ch)
    OPENING.include?(ch)
  end

  def closing_of(ch)
    CLOSING_MAP[ch]
  end
end