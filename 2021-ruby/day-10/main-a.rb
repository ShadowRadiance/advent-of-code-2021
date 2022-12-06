require "./syntax_parser"

class App
  def initialize
    @parser = SyntaxParser.new(load_data_from_input)
  end

  def run
    puts @parser.total_syntax_error_score
  end

  def load_data_from_input
    File.readlines("./data/input.txt", chomp: true)
  end
end

App.new.run
