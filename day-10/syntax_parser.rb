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

  COMPLETION_SCORES = {
    ")" => 1,
    "]" => 2,
    "}" => 3,
    ">" => 4,
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
    @errors ||= lines.map { |line| parse(line).first }.compact
  end

  def error_scores
    errors.map(&:score)
  end

  def total_syntax_error_score
    error_scores.sum
  end

  def autocompletions
    incomplete_lines.map { |line| completion_of(line) }
  end

  def autocomplete_scores
    autocompletions.map { |completion| completion_score(completion) }
  end

  def autocomplete_score
    median(autocomplete_scores)
  end

  def median(array)
    return nil if array.size == 0
    sorted = array.sort
    mid = (sorted.size / 2).to_i
    sorted[mid]
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
        return [SymbolNestingError.new(line, expected, char), working]
      end
    end
    
    [nil, working]
  end

  def open?(ch)
    OPENING.include?(ch)
  end

  def closing_of(ch)
    CLOSING_MAP[ch]
  end

  def incomplete_lines
    @incomplete_lines ||= (@lines - errors.map(&:line))
  end

  def completion_of(line)
    _, incomplete = parse(line)
    incomplete.chars.reverse.map { |ch| closing_of(ch) }.join
  end

  def completion_score(completion)
    completion.chars.reduce(0) do |memo, ch|
      memo * 5 + COMPLETION_SCORES[ch]
    end
  end
end